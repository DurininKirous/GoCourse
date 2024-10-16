package main

import "fmt"

func main() {
	var var1 *int
	var var2 int = 7
	var1 = &var2
	fmt.Println(*var1)
}
