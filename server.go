package simpletcpserver

import (
	"log"
	"net"
	"sync"
)

type TCPServer struct {
	Addr  string
	ln    net.Listener
	conns map[string]net.Conn
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
	s.conns = make(map[string]net.Conn)
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

// Add new connection to its connection pool
func (s *TCPServer) newConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	log.Printf("New connection! %v", remoteAddr)
	s.mtx.Lock()
	if _, ok := s.conns[remoteAddr]; ok {
		log.Fatalf("%v already exists in connection pool!\n", remoteAddr)
	}
	s.conns[remoteAddr] = conn
	s.mtx.Unlock()
}

// Send "ping" string to all connections, and remove dead connections
func (s *TCPServer) PingAll() {
	ch := make(chan string, 10)
	s.mtx.Lock()
	connNum := len(s.conns)
	for addr, c := range s.conns {
		go s.Ping(addr, c, ch)
	}
	for i := 0; i < connNum; i++ {
		addrToRemove := <-ch
		if addrToRemove != "" {
			delete(s.conns, addrToRemove)
		}
	}
	s.mtx.Unlock()
}

// Send "ping" to given connection, send addr if fail
func (s *TCPServer) Ping(addr string, c net.Conn, ch chan string) {
	_, err := c.Write([]byte("ping\n"))
	if err != nil {
		log.Printf("%v disconnected, %v", addr, err)
		ch <- addr
	}
	ch <- ""
}
