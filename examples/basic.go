package main

import (
	"fmt"
	"log"

	"github.com/fesiqp/pass"
)

func main() {
	fmt.Println("pass Go wrapper")

	if err := pass.Add("testtt", "mypasswd"); err != nil {
		log.Fatal(err)
	}

	passwd, err := pass.Get("testtt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(passwd)

	out, err := pass.Remove("testtt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)
}
