package main

import (
	"sync"
)


func main() {
	s := NewServer("tcp", "127.0.0.1", 8081);
	var wg sync.WaitGroup;
	wg.Add(1);
	s.Start();
	wg.Wait();
}


