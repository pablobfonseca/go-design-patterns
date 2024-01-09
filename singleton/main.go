package main

import "fmt"

func main() {
	fmt.Println("Singleton Pattern")

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning a a newline and
	// after the final item there must be a newline of EOF
	fmt.Scanln()
}
