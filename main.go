package main

import "github.com/rodjc/gofasta/cmd"

func main() {
	cmd := cmd.NewCmdFlags()
	cmd.Execute()
}
