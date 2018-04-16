package simpletcpserver

import (
	"bufio"
	"net"
	"os"
	"testing"
	"time"
)

const addr = ":21637"

func TestMain(m *testing.M) {
	ret := m.Run()

	os.Exit(ret)
}

func TestNew(t *testing.T) {
	var (
		s   *TCPServer
		err error
	)
	if s, err = New(addr); err != nil {
		t.Fatal(err)
	}

	if _, err = net.Dial("tcp", addr); err != nil {
		t.Fatal(err)
	}
	s.Stop()
}

func TestPingAll(t *testing.T) {
	var (
		s    *TCPServer
		conn net.Conn
		err  error
	)
	if s, err = New(addr); err != nil {
		t.Fatal(err)
	}

	if conn, err = net.Dial("tcp", addr); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)
	go s.PingAll()
	ret, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if ret != "ping\n" {
		t.Fatalf("Expected ping but %s returned", ret)
	}
	s.Stop()
}
