package main

import (
	"os"
	//"fmt"
	"strconv"
)

func parseInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	//fmt.Println("Welcome to the website helper scripts")
	args := os.Args[1:]
	if len(args) > 1 {
		if args[0] == "--recent" {
			numToGet := parseInt(args[1])
			GetNewestArticles(numToGet)
		}
	}
}
