package main

import "fmt"

func main() {

	countryCode := make(map[int]string)
	countryCode[62] = "Indonesia"
	countryCode[1] = "United States"
	countryCode[44] = "United Kingdom"
	fmt.Println("countryCode :", countryCode)

	countryCode[62] = "Jepang"
	fmt.Println("countryCode :", countryCode)

	delete(countryCode, 44)
	fmt.Println("countryCode :", countryCode)

	i, prs := countryCode[1]
	fmt.Println("present :", prs)
	fmt.Println("present at:", i)

	i, prs = countryCode[33]
	fmt.Println("present :", prs)
	fmt.Println("present at:", i)
}
