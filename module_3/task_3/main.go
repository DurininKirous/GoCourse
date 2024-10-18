package main

import (
	"fmt"

	"task_3/durininkirous"
	"task_3/wordz"

	"github.com/DurininKirous/utils"
	utilsV2 "github.com/DurininKirous/utils/v2"
	"github.com/huandu/xstrings"
)

func main() {
	isExistV2 := utilsV2.InSlice("two", wordz.Words)
	if isExistV2 {
		fmt.Println("Using utilsV2.InSlice and find value")
	}
	isExist := utils.Contains(wordz.Random(), wordz.Words)
	if isExist {
		fmt.Println("Slice Words contain finding value")
	} else {
		fmt.Println("Slice Words doesn't contain finding value")
	}
	isExistInt := utils.ContainsInt(5, []int{1, 2, 3, 4, 5})
	if isExistInt {
		fmt.Println("Slice Int containn finding value")
	} else {
		fmt.Println("Slice Int doesn't contain finding value")
	}
	fmt.Println(xstrings.Shuffle(durininkirous.City()), durininkirous.Digit())
}
