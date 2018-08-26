package singleton

import (
	"sync"
	"time"
)

// singletonType ...
type singletonType struct {
	iField   int32
	strField string
}

var _instance *singletonType

//var lock = sync.Mutex{}

var lock sync.Mutex

var once = sync.Once{} // equals to var once sync.Once
//var once sync.Once

var onceaddr1 = &once

var onceaddr2 = &onceaddr1

// GetInstance ...
func GetInstance() *singletonType {

	return getInstance2()

}

// GetInt ...
func (s *singletonType) GetInt() int32 {
	return s.iField
}

// GetString ....
func (s *singletonType) GetString() string {
	return s.strField
}

func getInstance() *singletonType {
	return &singletonType{-1, "IamSingleton"}
}

func getInstance0() *singletonType {
	if _instance == nil {
		doInit()
	}

	return _instance
}

func getInstance1() *singletonType {

	if _instance == nil {

		lock.Lock()

		if _instance == nil {
			doInit()
		}

		lock.Unlock()
	}

	return _instance
}

func getInstance2() *singletonType {

	once.Do(func() {
		doInit()
	})

	return _instance
}

func doInit() {
	time.Sleep(1 * time.Second)
	_instance = &singletonType{-1, "IamSingleton"}
}
