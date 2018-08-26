package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

// Server ...
type Server struct {
	protocal string
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
	l, err := net.Listen(s.protocal, s.host+":"+strconv.Itoa(s.port))
	if err != nil {
		fmt.Println("Start server error", err.Error());
		return;
	}
	for {
		conn, err := l.Accept();
		if err != nil {
			fmt.Println("Server accept error, shutdown...", err.Error());
			return;
		}
		go handlConn(conn);
	}
}

func handlConn(conn net.Conn) {


	for {
		var b []byte;
		b = make([]byte, 128);
		n, err := conn.Read(b);
		if(err != nil) {
			// ...
			fmt.Println("Read error", err.Error());
			return;
		}
		var s = string(b[:n]);
		fmt.Printf("Rcv %d bytes: %s\n", n, s);
		//conn.Close();
		//return;
	}
	
}

