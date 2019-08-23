package main

import (
	"fmt"
	"go-gaia-updater/cmd"
	"go-gaia-updater/cmd/reverse"
)

func main() {
	fmt.Println(reverse.Reverse("!oG ,olleH"))
	cmd.Execute()
}
