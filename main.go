package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// fmt.Println("Hello, world")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input := scanner.Text()
	// fmt.Printf("you are typed: %q", input)

	// var dollarPrice int64 = 180

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	// answer := input * dollarPrice
	// fmt.Printf("LKR : %d", answer)

	var dollarPrice float64 = 180.9
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Printf("Doller price: %f", dollarPrice)

	fmt.Printf("Rupees : ")
	//string to flout converter
	f, _ := strconv.ParseFloat(scanner.Text(), 64)

	total := (f / dollarPrice)
	fmt.Printf("Doller price: %f", total)

}
