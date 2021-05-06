package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	}
	ans := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		ans[i] = i + min
	}
	return ans
}
