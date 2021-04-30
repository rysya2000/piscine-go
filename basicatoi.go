package piscine

func BasicAtoi(s string) int {
    n := 0
    for _, ch := range []byte(s) {
        ch -= 48
        n = n*10 + int(ch)
    }
    return int(n)
}