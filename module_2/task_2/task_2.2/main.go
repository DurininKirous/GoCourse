package main

import (
	"fmt"
	"math"
)

func main() {
	var Length float64 = 35
	var Radius float64 = Length / (float64)(2*math.Pi)
	RadiusPtr := &Radius
	var Square float64 = math.Pi * math.Pow(*RadiusPtr, 2)
	fmt.Printf("%.2f\n", Square)
}
