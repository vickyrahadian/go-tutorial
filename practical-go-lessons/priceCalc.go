package main

import "fmt"

func main() {
	johnPrice := computePrice(145.90, 3)
	fmt.Printf("TOTAL : %0.2f $", johnPrice)
}

// compute the price with a 200% margin
func computePrice(rate float32, nights int) (price float32) {
	p := rate * float32(nights)
	price = p * 2
	return
}
