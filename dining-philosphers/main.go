package main

import (
	"fmt"
	"sync"
	"time"
)

type Phillospher struct {
	Name      string
	leftFork  int
	rightFork int
}

var phillosphers = []Phillospher{
	{Name: "Plato", leftFork: 4, rightFork: 0},
	{Name: "Aristotle", leftFork: 0, rightFork: 1},
	{Name: "Socrates", leftFork: 1, rightFork: 2},
	{Name: "Descartes", leftFork: 2, rightFork: 3},
	{Name: "Kant", leftFork: 3, rightFork: 4},
}

var hunger = 2 // how many times each philosopher eats
var eatTime = 1 * time.Second
var wg sync.WaitGroup

func eat(Phillospher Phillospher, wg *sync.WaitGroup, forks map[int]*sync.Mutex) {
	defer wg.Done()
	for i := 0; i < hunger; i++ {
		name := Phillospher.Name
		leftFork := Phillospher.leftFork
		rightFork := Phillospher.rightFork
		fmt.Println(name, "is trying to pick up forks", leftFork, "and", rightFork)
		forks[leftFork].Lock()
		fmt.Println(name, "picked up left fork", leftFork)
		forks[rightFork].Lock()
		fmt.Println(name, "picked up right fork", rightFork)
		fmt.Println(name, "is eating...")
		time.Sleep(eatTime)
		forks[leftFork].Unlock()
		fmt.Println(name, "put down left fork", leftFork)
		forks[rightFork].Unlock()
		fmt.Println(name, "put down right fork", rightFork)
		fmt.Println(name, "finished eating", i+1, "times")
	}
}

func main() {
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(phillosphers); i++ {
		forks[i] = &sync.Mutex{}
	}
	var wg sync.WaitGroup
	wg.Add(len(phillosphers))

	for _, p := range phillosphers {
		go eat(p, &wg, forks)
	}

	wg.Wait()
}
