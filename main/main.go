package main

import (
	"time"
	// "fmt"
	// "time"
	log "github.com/AlexStocks/log4go"
)

func main() {
	log.LoadConfiguration("log.xml")
	log.Info("Hello %v", "Log")

	time.Sleep(5 * time.Second)
	// now := time.Now()

	// fmt.Printf("%v, %d, %d\n", now.Month(), now.Day(), now.Year())
}
