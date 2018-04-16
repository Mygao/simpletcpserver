package main

import (
	"flag"
	"log"
  "net"
  "bufio"
)

var Addr = flag.String("addr", ":8080", "Server listening address")

func main() {
  flag.Parse()
  conn, err := net.Dial("tcp", *Addr)
  if err != nil {
    log.Fatal(err)
  }
  for {
    ret, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
      log.Fatal(err)
    }
    log.Print(ret)
  }
}
