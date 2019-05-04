package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

// Server ...
type Server struct {
	protocol string
	host     string
	port     int
}

var _s *Server

var once sync.Once

// NewServer return the Server
func NewServer(p, h string, port int) *Server {
	once.Do(func() {
		var s = Server{p, h, port}
		_s = &s
	})

	return _s
}

// Start ...
func (s *Server) Start() {
	l, err := net.Listen(s.protocol, s.host+":"+strconv.Itoa(s.port))
	if err != nil {
		fmt.Println("Start server error", err.Error())
		return
	}
	fmt.Printf("Server started at %s:%s:%d\n", s.protocol, s.host, s.port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Server accept error, shutdown...", err.Error())
			return
		}
		go handelConn(conn)
	}
}

func handelConn(conn net.Conn) {

	for {
		var b []byte
		b = make([]byte, 128)
		n, err := conn.Read(b)
		if err != nil {
			// ...
			if io.EOF == err {
				fmt.Println("Remote closed.")
			} else {
				fmt.Println("Read error", err.Error())
			}
			return
		}
		var s = string(b[:n])
		fmt.Printf("Rcv %d bytes: %s\n", n, s)
	}

}
