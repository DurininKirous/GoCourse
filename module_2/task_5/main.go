package main

import "fmt"

func main() {
	var strs []string = []string{"123", "321", "111"}
	var res bool = Contains(strs, "111")
	var max int = GetMax(1, 5, 2, 11, 100, 2, 3, 1)
	fmt.Println(res, max)
}

func Contains(a []string, x string) bool {
	var res = false
	for _, str := range a {
		if x == str {
			res = true
			break
		}
	}
	return res
}
func GetMax(elements ...int) (max int) {
	for _, element := range elements {
		if element > max {
			max = element
		}
	}
	return
}
