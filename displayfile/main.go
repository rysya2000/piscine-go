package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("file name missing")
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
	}
	if len(args) == 1 {
		content, err := ioutil.ReadFile(args[0])
		if len(content) > 0 {
			fmt.Println(string(content))
		} else {
			fmt.Println(err)
		}
	}
}
