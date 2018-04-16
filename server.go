package simpletcpserver

import (
	"log"
	"net"
	"sync"
)

type TCPServer struct {
	Addr  string
	ln    net.Listener
	conns []net.Conn
	mtx   sync.Mutex // Connection pool locker
	done  chan bool
}

func New(addr string) (*TCPServer, error) {
	s := &TCPServer{Addr: addr, mtx: sync.Mutex{}, done: make(chan bool, 1)}
	var err error
	if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
		return nil, err
	}
	log.Printf("Started TCP server (%v)", s.Addr)
	s.conns = make([]net.Conn, 0, 10)
	go s.listen()
	return s, err
}

func (s *TCPServer) Stop() {
	s.done <- true
	s.ln.Close()
}

func (s *TCPServer) listen() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			select {
			case <-s.done:
				log.Print("Server stopped")
				return
			default:
				log.Print(err)
			}
		} else {
			go s.newConnection(conn)
		}
	}
}

func (s *TCPServer) newConnection(conn net.Conn) {
	log.Printf("New connection! local: %v, remote: %v", conn.LocalAddr(), conn.RemoteAddr())
	s.mtx.Lock()
	s.conns = append(s.conns, conn)
	s.mtx.Unlock()
}

func (s *TCPServer) PingAll() {
	s.mtx.Lock()
	for _, c := range s.conns {
		go s.Ping(c)
	}
	s.mtx.Unlock()
}

func (s *TCPServer) Ping(c net.Conn) {
	c.Write([]byte("ping\n"))
}
