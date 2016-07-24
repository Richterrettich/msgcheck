package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/richterrettich/changelog/domain"
)

func main() {
	file := os.Args[1]
	result := domain.Commit{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
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
