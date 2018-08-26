package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)




func TestMain(t *testing.T) {


	fmt.Println("runtime.GOMAXPROCS:", runtime.GOMAXPROCS(4));

	fmt.Println("runtime.NumCPU:", runtime.NumCPU());
	fmt.Println("runtime.NumCgoCall:", runtime.NumCgoCall());
	fmt.Println("runtime.NumGoroutine:", runtime.NumGoroutine());

	
}