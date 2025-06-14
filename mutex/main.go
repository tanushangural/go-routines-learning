package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	var mt sync.Mutex
	wg.Add(2)
	go updateMessage("a", &mt)
	go updateMessage("b", &mt)
	wg.Wait()
	fmt.Println(msg)
}
