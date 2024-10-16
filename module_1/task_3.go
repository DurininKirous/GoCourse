package main

import (
	"fmt"
)

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64
	var var1 EuropeanVelocity = 120.4 * 3.6
	var var2 AmericanVelocity = 130 * (3600.0 / 1609.0)
	fmt.Printf("%.2f %.2f\n", var1, var2)
}
