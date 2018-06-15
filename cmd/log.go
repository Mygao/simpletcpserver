package main

import (
	zmq "github.com/pebbe/zmq4"
	server "github.com/aiqu/simpletcpserver"
	"flag"
	"log"
)

var Addr = flag.String("addr", ":8080", "Server listening address")

func main() {
  log_sub, err := zmq.NewSocket(zmq.SUB)
	defer log_sub.Close()

	if err != nil {
		panic(err)
	}

	log_sub.Connect("ipc:///tmp/log")
	log_sub.SetSubscribe("")

	flag.Parse()
	s, err := server.New(*Addr)
	if err != nil {
		log.Fatal(err)
	}

  for {
    msg, e := log_sub.Recv(0)

    if e != nil {
      log.Println(e)
      return
    }

    log.Println(msg)

    if msg == "start" {
			s.SendAll("start\n")
    } else if msg == "stop" {
			s.SendAll("stop\n")
    } else {
    }

  }
}
