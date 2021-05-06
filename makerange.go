package main

import "fmt"

func MakeRange(min, max int) []int {
	ans := make([]int, max, min)
	for i := min; i < max; i++ {
		ans[i] = i
	}
	return ans
}
func main() {
	fmt.Println(MakeRange(5, 10))
	//fmt.Println(MakeRange(10, 5))
}
