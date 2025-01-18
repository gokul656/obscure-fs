package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gokul656/obscure-fs/cmd"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	cmd.Execute()
}
