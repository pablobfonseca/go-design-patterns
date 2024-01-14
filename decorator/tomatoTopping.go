package main

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() float64 {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7.0
}
