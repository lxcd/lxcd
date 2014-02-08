package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/lxcd/lxcd/lxc"
)

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	server := rpc.NewServer()
	server.Register(new(lxc.LXC))
	log.Print("waiting for connections ...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Listen failed:", err)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
