package main

import (
	"fmt"
	"strconv"
)

func main() {
	var var1 string = "104"
	var var2 int = 35
	var3, _ := strconv.Atoi(var1)
	var4 := strconv.Itoa(var2)
	fmt.Printf("%T %T\n%d %s\n", var3, var4, var3, var4)
}
