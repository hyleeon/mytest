package test

import (
	"sync"
	"fmt"
	"time"
	"testing"
	"gohelloworld.com/main2"
	"gohelloworld.com/singleton"
	log "github.com/AlexStocks/log4go"
)


func TestEnv(t *testing.T) {
	main2.Env()
}


func TestPanicRecover(t *testing.T) {
	//fmt.Println("TestAbc2")
	main2.PanicRecover();
}


func TestLog(t *testing.T) {
	log.LoadConfiguration("log.xml")
	log.Info("Hello Log!")
	time.Sleep(time.Second * 5)
}


func TestSingleton(t *testing.T) {

	var wg = sync.WaitGroup{}

	var times = 8

	wg.Add(times)

	for i:=0;i<times;i++ {
		go func(i int) {

			s := singleton.GetInstance()
	
			fmt.Printf("s_%2d: %p\n", i, s)
	
			wg.Done()
		}(i)
	}

	wg.Wait()
}

