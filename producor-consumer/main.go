package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzaMade, pizzaFailed, totalPizzas int

type PizzaOrder struct {
	pizzaNumber int
	success     bool
	message     string
}

type Producor struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producor) close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		color.Green("Received order #%d!", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		totalPizzas++
		color.Green("Making pizza #%d. It will take %d seconds...", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd <= 2 {
			msg = color.RedString("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		}
		if rnd <= 4 && rnd > 2 {
			msg = color.RedString("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = color.GreenString("Pizza order #%d is ready!", pizzaNumber)
		}
		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			success:     success,
			message:     msg,
		}
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMake *Producor) {
	// keep track of pizzas we are making
	var pizzaCount int
	//run forever until we receive a quit notification
	// try to make a pizza
	var i = 0
	fmt.Printf("Total pizzas made: %d\n", i)
	for {
		// try to make a pizza
		currentPizza := makePizza(pizzaCount)
		i = currentPizza.pizzaNumber
		if currentPizza != nil {
			select {
			case pizzaMake.data <- *currentPizza:

			case quitChan := <-pizzaMake.quit:
				close(pizzaMake.data)
				color.Red("Pizzeria is closing down!")
				close(quitChan)
			}
		}
		pizzaCount++
	}
}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print out message
	color.New(color.FgCyan).Println("pizzeria is open for business")

	//create a producor
	pizzaJob := &Producor{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	//run the producor in background
	go pizzeria(pizzaJob)

	//create and run cosumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Pizza #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("customer is not happy with pizza #%d!", i.pizzaNumber)
			}
		} else {
			color.Cyan("All pizzas have been made!")
			err := pizzaJob.close()
			if err != nil {
				color.Red("Error closing pizzeria: %v", err)
			}
		}
	}
}
