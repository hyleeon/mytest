package main2

import (
	"bufio"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"

	"encoding/json"
	"strings"
	"sync"

	"gohelloworld.com/config"
	"gohelloworld.com/math"
)

func main() {

	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	math.Add(1, 2)

	fmt.Println("config: ", config.Get("abc"))

	fmt.Println("Hello World!")

	r := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := r.ReadLine()
		fmt.Println("Read line: ", string(line))

		str := string(line)

		v, e := strconv.Atoi(str)

		if e != nil {
			fmt.Println(e.Error())
			return
		}

		fmt.Println("Read value: ", v)

		ch := make(chan int, 1)

		//read
		go func(v int) {

			defer func() {
				fmt.Println("read func gg....")
			}()
			if v == 6 {
				ch <- 1
			} else {
				go func() {
					time.Sleep(time.Second * 3)
					fmt.Println("nil retrue...")
				}()
			}
		}(v)

		//write
		go func() {
			select {
			case vv := <-ch:
				fmt.Println("Recv vv", vv)
				for i := 10; i > 0; i-- {
					fmt.Println("Write will over", i)
					time.Sleep(time.Second * 1)
				}
				fmt.Println("Write over")

			}
		}()

	}

}

// PanicRecover do test panic and recover
func PanicRecover() {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println("Err: ", err)
		}

	}()

	panic("abcccccccc ")

}

var waitgroup = sync.WaitGroup{}

// Env tests evn
func Env() {

	envs := os.Environ()

	fmt.Println("len(envs): ", len(envs))

	waitgroup.Add(len(envs))

	for _, env := range envs {
		//fmt.Printf("env %d: %s\n", i, env);
		//kv := strings.Split(env, "=");
		//fmt.Printf("env %d: %s = %s\n", i, kv[0], kv[1]);
		showStringSplit(env)
	}

	waitgroup.Wait()

	vnc := os.Getenv("VSCODE_NLS_CONFIG")

	//vnc = strings.Replace(vnc, "\\\\", "\\", 10)

	fmt.Println("VSCODE_NLS_CONFIG: ", vnc)
	fmt.Println("GOROOT: ", os.Getenv("GOROOT"))

	if len(vnc) != 0 {
		var m map[string]interface{}
		fmt.Println("map: ", &m)
	
		err := json.Unmarshal(([]byte)(vnc), &m)
		if err != nil {
			fmt.Println("Parse json error!", err)
		}
		//
		_cacheRoot := m["_cacheRoot"]
	
		fmt.Println("_cacheRoot: ", _cacheRoot)

		availableLanguages := (m["availableLanguages"]).(map[string]interface{})

	x := availableLanguages["*"]

	fmt.Println("availableLanguages -> *: ", x)
	}
}

func showStringSplit2(str string) {

	defer waitgroup.Done()

	arr := strings.Split(str, "")
	for _, cs := range arr {
		fmt.Printf("%s,", cs)
	}
	fmt.Println("")
}

func showStringSplit(str string) {

	defer waitgroup.Done()

	//len := len(str)
	//var buf [len]byte
	buf := ([]byte)(str)

	var isFirst bool
	isFirst = true

	for _, ch := range buf {
		var tstr string
		if '=' == ch {
			tstr = "============="
		} else {
			tstr = string(ch)
		}
		if isFirst {
			fmt.Printf("%s", tstr)
			isFirst = false
		} else {
			fmt.Printf(",%s", tstr)
		}
	}

	fmt.Println("")

}
