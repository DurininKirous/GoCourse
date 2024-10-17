package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inPath := "./data/08_task_in.txt"
	file, err := os.Open(inPath)
	if err != nil {
		panic("error of opening file")
	}
	defer func() {
		if res := file.Close(); res != nil {
			fmt.Errorf("some error: %s", res)
		}
	}()
	s := bufio.NewReader(file)
	var countRows int
	for {
		_, err := s.ReadString('\n')
		if err == io.EOF {
			break
		}
		countRows++
	}
	fmt.Printf("Total strings: %d\n", countRows)
}
