package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			content, err := ioutil.ReadFile(args[i])
			if err != nil {
				for _, j := range "ERROR:" {
					z01.PrintRune(rune(j))
				}
				io.WriteString(os.Stdout, err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				io.WriteString(os.Stdout, string(content))
			}
		}
	}
}
