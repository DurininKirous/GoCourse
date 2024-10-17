package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	outPath := "./data/07_task_out.txt"
	inPath := "./data/07_task_in.txt"
	in, _ := os.Open(inPath)
	defer in.Close()
	out, _ := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0644)
	defer out.Close()
	s := bufio.NewReader(in)
	var countRows int
	for {
		line, err := s.ReadString('\n')
		parts := StringSplit(line)
		if !CheckPartsOfString(parts) {
			PanicFunction("parse error: empty field on string "+strconv.Itoa(countRows), outPath)
		}
		rows := "Rows: " + strconv.Itoa(countRows) + "\n"
		name := "Name: " + parts[0] + "\n"
		address := "Address: " + parts[1] + "\n"
		city := "City: " + parts[2] + "\n\n\n"
		total := rows + name + address + city
		out.Write([]byte(total))
		if err == io.EOF {
			break
		}
		countRows++
	}
}

func StringSplit(str string) []string {
	return strings.Split(str, "|")
}

func CheckPartsOfString(parts []string) bool {
	var res bool = false
	if parts[0] != "" && parts[1] != "" && parts[2] != "" {
		res = true
	}
	return res
}

func PanicFunction(str string, filePath string) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "regular error":
				err = fmt.Errorf("application error")
			default:
				PrintFileData(filePath)
				panic(panicErr)
			}
		}
	}()

	panic(str)
}

func PrintFileData(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
