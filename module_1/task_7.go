package main

import "fmt"

func main() {
	var LenOfBooks int
	var LenOfUsers int
	Users := map[string]map[string][]string{}
	Users["Ivan"] = map[string][]string{
		"First": {"First", "Second", "Third", "Fourth"},
	}
	Users["Alexey"] = map[string][]string{
		"First": {"First", "Second", "Third", "Fourth"},
	}
	Users["Alexey"]["Second"] = append(Users["Alexey"]["Second"], "First")
	LenOfUsers = len(Users)
	for _, User := range Users {
		for _, Book := range User {
			LenOfBooks += len(Book)
		}
	}
	fmt.Println(LenOfUsers, LenOfBooks)
}
