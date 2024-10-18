package main

import (
	"fmt"

	"task_3/durininkirous"
	"task_3/wordz"

	"github.com/DurininKirous/GoCourse/utils"
	"github.com/huandu/xstrings"
)

func main() {
	isExist := utils.Contains(wordz.Random(), wordz.Words)
	if isExist {
		fmt.Println("Slice Words contain finding value")
	} else {
		fmt.Println("Slice Words doesn't contain finding value")
	}
	fmt.Println(xstrings.Shuffle(durininkirous.City()), durininkirous.Digit())
}
