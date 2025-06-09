package main

import (
	"fmt"
	"sync"
)

func main() {

	// words := []string{
	// 	"a",
	// 	"b",
	// 	"c",
	// 	"d",
	// 	"e",
	// }

	// var wg sync.WaitGroup
	// wg.Add(len(words))

	// for ind, val := range words {
	// 	go printSomething(fmt.Sprintf("%d : %s", ind, val), &wg)
	// }

	// wg.Wait()

	// wg.Add(1)
	// printSomething("second", &wg)

	execute()

}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
