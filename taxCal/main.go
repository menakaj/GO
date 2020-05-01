package main

import "fmt"

func main() {
	var totalTax float32
	itemTax := map[float32]float32{
		0.99: 7.5,
		2.75: 1.5,
		0.87: 2,
	}

	for p, t := range itemTax {
		fmt.Println("Price: ", p, " rate: ", t, " Tax:", p*(t/100))
		totalTax += p * (t / 100)
	}

	fmt.Println("total tax : ", totalTax)
}
