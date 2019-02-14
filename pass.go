package pass

import (
	"bytes"
	"errors"
	"os/exec"
)

type passCmd struct {
	cmd *exec.Cmd
}

func (c *passCmd) Run() error {
	if c.cmd.Stderr != nil {
		return errors.New("exec: Stderr already set")
	}
	var errBuf bytes.Buffer
	c.cmd.Stderr = &errBuf
	if err := c.cmd.Run(); err != nil {
		return errors.New(errBuf.String())
	}
	return nil
}

func (c *passCmd) Output() (string, error) {
	if c.cmd.Stdout != nil {
		return "", errors.New("exec: Stdout already set")
	}
	if c.cmd.Stderr != nil {
		return "", errors.New("exec: Stderr already set")
	}
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	c.cmd.Stdout = &outBuf
	c.cmd.Stderr = &errBuf
	if err := c.cmd.Run(); err == nil {
		return outBuf.String(), nil
	}
	return outBuf.String(), errors.New(errBuf.String())
}

func makePassCmd() *passCmd {
	return &passCmd{exec.Command("pass")}
}

func (c *passCmd) WithArgs(args ...string) *passCmd {
	c.cmd.Args = append(c.cmd.Args, args...)
	return c
}

func (c *passCmd) WithInput(input string) *passCmd {
	c.cmd.Stdin = bytes.NewBuffer([]byte(input + "\n"))
	return c
}

func Get(name string) (string, error) {
	return makePassCmd().
		WithArgs("show", name).
		Output()
}

func Add(name, password string) error {
	return makePassCmd(). // guaranteed to not print anything to stdout on success
				WithArgs("insert", "-ef", name).
				WithInput(password).
				Run()
}

// Update is exactly the same function as the Insert, just renamed to make it "more pleasing to the eyes" when used
func Update(name, password string) error {
	return Add(name, password)
}

func Remove(name string) (string, error) {
	return makePassCmd().
		WithArgs("rm", name).
		Output()
}
