package main

import (
	"bufio"
	"flag"
	server "github.com/aiqu/simpletcpserver"
	"log"
	"os"
	"strings"
)

var Addr = flag.String("addr", ":8080", "Server listening address")

func printHelpMessage() {
	log.Print("Press key to send signal\n")
	log.Print("p - send ping to all client\n")
	log.Print("r - send start signal to all client\n")
	log.Print("s - send stop signal to all client\n")
}

func main() {
	flag.Parse()
	s, err := server.New(*Addr)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		printHelpMessage()
		text, _ := reader.ReadString('\n')
		switch {
		case strings.Compare(text, "p\n") == 0:
			s.SendAll("ping\n")
		case strings.Compare(text, "r\n") == 0:
			s.SendAll("start\n")
		case strings.Compare(text, "s\n") == 0:
			s.SendAll("stop\n")
		}
	}
}
