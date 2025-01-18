package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gokul656/obscure-fs/internal/networking"
	"github.com/libp2p/go-libp2p/core/protocol"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		return
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid port: %s\n", err)
	}

	ctx := context.Background()
	network := networking.NewNetwork(ctx, port)

	protocolID := "oscure-fs/1.0.0"
	network.StartSimpleProtocol(protocol.ID(protocolID))

	log.Printf("Node ID: %s\n", network.GetHost().ID().String())
	log.Println("Waiting for connections...")

	if port == 5002 {
		err = network.ConnectToPeer("/ip4/127.0.0.1/tcp/5001/p2p/12D3KooWQ9bJZ9GjAw7JDR2LNvibL2Jgd6uqkC9yTBxovPf27oja")
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Successfully connected to peer")
		}
	}

	select {}
}
