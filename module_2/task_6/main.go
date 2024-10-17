package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	LogTime()
	in, _ := os.Open("./data/06_task_in.txt")
	defer in.Close()
	out, _ := os.OpenFile("./data/06_task_out.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer out.Close()
	s := bufio.NewReader(in)
	var countRows int
	var countBytes int
	var str string
	for {
		line, err := s.ReadString('\n')
		str = strconv.Itoa(countRows) + " " + line
		out.Write([]byte(str))
		if err == io.EOF {
			break
		}
		countRows++
		countBytes += len(line)
	}
	fmt.Println(countRows, countBytes)
}

func LogTime() {
	start := time.Now()
	defer func() {
		res := time.Since(start)
		fmt.Print(res, " ")
	}()
}
