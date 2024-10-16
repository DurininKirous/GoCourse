package main

import (
	"fmt"
)

func main() {
	src := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	wrk := make([]string, 5)
	copy(wrk, src[:5])
	src = src[5:]
	fmt.Println(wrk)
	fmt.Println(src)
	wrk = append(wrk, src...)
	fmt.Println(wrk)
}
