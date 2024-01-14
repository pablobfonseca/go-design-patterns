package main

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() float64 {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10.0
}
