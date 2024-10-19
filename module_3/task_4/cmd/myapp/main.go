package main

import (
	"fmt"
	"module_3/task_4/internal/durininkirous"

	xstrings "github.com/huandu/xstrings"
)

func main() {
	fmt.Println(xstrings.Shuffle(durininkirous.City()), durininkirous.Digit())
}
