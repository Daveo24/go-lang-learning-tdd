package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var mutex = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(2)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go sayHello()
		mutex.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	mutex.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	mutex.Unlock()
	wg.Done()
}
