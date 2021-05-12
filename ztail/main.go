package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	x := 0
	b := false
	if len(args) > 0 {
		for i := 2; i < len(args); i++ {
			content, err := os.Open(args[i])
			if err != nil {
				x++
				fmt.Printf("%v\n", err.Error())
				if b == true {
					fmt.Printf("\n")
					b = false
				}
			} else {
				digit := 0
				for _, t := range []byte(args[1]) {
					t -= 48
					digit = digit*10 + int(t)
				}
				if len(args) > 3 {
					fmt.Printf("==> %v <==\n", args[i])
				}
				size, _ := os.Stat(args[i])
				ans := make([]byte, size.Size())
				content.Read(ans)
				fmt.Printf("%v", string(ans[len(ans)-digit:]))
				if i != len(args)-1 {
					if i+1 < len(args) {
						_, err1 := os.Open(args[i+1])
						if err1 != nil {
							b = true
						} else {
							fmt.Printf("\n")
						}
					}
				}
			}
		}
		if x > 0 {
			os.Exit(1)
		}
	}
}
