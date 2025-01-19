package main

import (
	_ "net/http/pprof"

	"github.com/gokul656/obscure-fs/cmd"
)

func main() {
	cmd.Execute()
}
