package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {

	// fmt.Println("runtime.GOMAXPROCS:", runtime.GOMAXPROCS(8));

	fmt.Println("runtime.NumCPU:", runtime.NumCPU())
	fmt.Println("runtime.NumCgoCall:", runtime.NumCgoCall())
	fmt.Println("runtime.NumGoroutine:", runtime.NumGoroutine())

}
