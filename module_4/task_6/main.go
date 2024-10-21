package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type LimitExceededError struct {
	message     string
	stringLimit int
	lastString  string
}

func (e *LimitExceededError) Error() string {
	return fmt.Sprintf("%s, limit: %d, last string: %s", e.message, e.stringLimit, e.lastString)
}

func countLinesWithLimit(filePath string, limit int) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var countRows int

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, fmt.Errorf("error reading file: %w", err)
		}

		countRows++
		if countRows > limit {
			return countRows, &LimitExceededError{
				message:     "string count exceeded limit",
				stringLimit: limit,
				lastString:  line,
			}
		}
	}

	return countRows, nil
}

func main() {
	inPath := "./data/08_task_in.txt"
	limit := 152

	count, err := countLinesWithLimit(inPath, limit)
	if err != nil {
		var limitErr *LimitExceededError
		if errors.As(err, &limitErr) {
			fmt.Println("string count exceed limit, please read another file =) err: ", err.Error())
			return
		}
		fmt.Println("an error occurred:", err)
		return
	}

	fmt.Printf("Total strings: %d\n", count)
}
