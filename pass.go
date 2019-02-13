package pass

import (
	"bytes"
	"errors"
	"os/exec"
)

type passCmd struct {
	exec.Cmd
}

func (c *passCmd) Output() (string, error) {
	if c.Stdout != nil {
		return "", errors.New("exec: Stdout already set")
	}
	if c.Stderr != nil {
		return "", errors.New("exec: Stderr already set")
	}
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	c.Stdout = &outBuf
	c.Stderr = &errBuf
	if err := c.Run(); err == nil {
		return outBuf.String(), nil
	}
	return outBuf.String(), errors.New(errBuf.String())
}

func makePassCmd() *passCmd {
	cmd := exec.Command("pass")
	return &passCmd{*cmd}
}

func (cmd *passCmd) WithArgs(args ...string) *passCmd {
	cmd.Args = append(cmd.Args, args...)
	return cmd
}

func (cmd *passCmd) WithInput(input string) *passCmd {
	cmd.Stdin = bytes.NewBuffer([]byte(input + "\n"))
	return cmd
}

func Show(name string) (string, error) {
	return makePassCmd().
		WithArgs("show", name).
		Output()
}

func Insert(name, password string) error {
	_, err := makePassCmd(). // guaranteed to not print anything to stdout on success
					WithArgs("insert", "-ef", name).
					WithInput(password).
					Output() // TODO: create a proper method to print only errors
	return err
}

func Update(name, password string) error {
	return Insert(name, password)
}

func Remove(name string) (string, error) {
	return makePassCmd().
		WithArgs("rm", name).
		Output()
}
