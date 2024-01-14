package main

import "fmt"

func main() {
	fmt.Println("Decorator Pattern")

	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of VeggieMania with tomato and cheese topping is $%.1f\n", pizzaWithCheeseAndTomato.getPrice())
}
