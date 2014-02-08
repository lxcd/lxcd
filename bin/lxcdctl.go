package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"

	"github.com/lxcd/lxcd/lxc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal("Dial failed", err)
	}

	var reply lxc.ContainersReply
	args := 0
	err = client.Call("LXC.Containers", &args, &reply)
	if err != nil {
		log.Fatal("Call failed", err)
	}
	data, err := json.MarshalIndent(&reply, "", "  ")
	if err != nil {
		log.Fatal("Encoding failed:", err)
	}

	fmt.Printf("%s\n", string(data))
}
