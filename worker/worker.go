package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Count = 0

type Result struct {
	Line    string
	LineNum int
	Path    string
	Count   int
}

type Results struct {
	Inner []Result
}

func NewResult(line string, lineNum int, path string, count int) Result {

	return Result{line, lineNum, path, count}
}

func FindInFile(path string, find string) *Results {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}
	results := Results{make([]Result, 0)}
	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), find) {
			Count += 1
			r := NewResult(scanner.Text(), lineNum, path, Count)
			results.Inner = append(results.Inner, r)
		}
		lineNum += 1

	}
	if len(results.Inner) == 0 {
		return nil
	} else {
		return &results
	}
}
