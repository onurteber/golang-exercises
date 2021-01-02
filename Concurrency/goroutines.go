package main

import (
	"runtime"
	"time"
)

func main() {

	// basic uses goroutine
	//go xFunc()
	//time.Sleep(100 * time.Millisecond)

	runtime.GOMAXPROCS(1)
	go xFunc()
	time.Sleep(100 * time.Millisecond)

}

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
