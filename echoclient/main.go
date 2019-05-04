package main

import (
	"sync"
)


func main() {
	var wg sync.WaitGroup;
	wg.Add(1);
	StartClient();
	wg.Wait();
}
