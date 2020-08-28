package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

import _ "github.com/mkevac/debugcharts"
import _ "net/http/pprof"

var arr []string

func main() {

	go func() {
		_ = http.ListenAndServe(":6060", nil)
	}()

	//start := time.Now()

	for i := 0; i < 4; i++ {
		go func(n int) {
			loop(100_000_000, func(i int) {

				if i%100000 == 0 {
					fmt.Println(i)
					time.Sleep(500 * time.Millisecond)
				}

				do(n, i)
			})
		}(i)
	}

	select {}
	//ctx := context.Background()
	//recall.Init(ctx)
	//recall.ReadData(ctx)
	//
	//end := time.Since(start)
	//fmt.Println(end)
	//
	//select {}

}

func loop(count int, fn func(i int)) {
	for i := 0; i < count; i++ {
		fn(i)
	}
}

func do(n, number int) {
	s := strconv.Itoa(number)
	if n%2 == 0 {
		doSomething(s)
	} else {
		doSomethingElse(s)
	}

}

func doSomething(s string) {
	arr = append(arr, s)
}

func doSomethingElse(s string) {
	arr = append(arr, s)
}
