package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/richterrettich/changelog/domain"
)

func main() {
	fileName := os.Args[1]
	result := domain.Commit{}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			content = content + "\n" + line
		}
	}
	content = strings.TrimSpace(content)
	if strings.ToLower(content) == "initial commit" {
		os.Exit(0)
	}
	parts := strings.Split(content, "\n\n")
	result.RawSubject = parts[0]
	if len(parts) > 1 {
		result.RawBody = parts[1]
	}
	result.ParseSubject()
	result.ParseBody()
	if len(result.Errors) > 0 {
		fmt.Println("the following errors occured:")
		for _, v := range result.Errors {
			fmt.Println(v)
		}
		os.Exit(1)
	}
}
