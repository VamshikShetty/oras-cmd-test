package main

import (
	orasroot "oras.land/oras/cmd/oras/root"
	"fmt"
)

func main() {
	fmt.Println("start cmd")
	orascmd := orasroot.New()

	args := []string{"push", "localhost:5000/hello-symlink-test:v1", "hello-oras"}
	orascmd.SetArgs(args)

	err := orascmd.Execute()
	fmt.Println(err)

	fmt.Println("cmd done")
}