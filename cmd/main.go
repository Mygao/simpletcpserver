package main

import (
	"flag"
	server "github.com/aiqu/simpletcpserver"
	"log"
	"time"
)

var Addr = flag.String("addr", ":8080", "Server listening address")

func main() {
	s, err := server.New(*Addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second)
		s.PingAll()
	}
}
