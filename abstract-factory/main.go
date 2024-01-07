package main

import "fmt"

func main() {
	fmt.Println("Abstract Factory Pattern")

	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetail(nikeShoe)
	printShirtDetail(nikeShirt)

	printShoeDetail(adidasShoe)
	printShirtDetail(adidasShirt)
}

func printShoeDetail(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShirtDetail(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %s\n", s.getSize())
}
