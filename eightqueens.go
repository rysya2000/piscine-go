package piscine

import "github.com/01-edu/z01"

var ans [8]int

func EightQueens() {
    SetQueen(0)
}

func Issafe(row int, col int) bool {
    if row == 0 {
        return true
    }
    for i := 0; i < row; j++ {
        if col == ans[i] {
            return false
        }
        x := row - i
        y := col - ans[i]
        if x < 0 {
            x *= -1
        }
        if y < 0 {
            y *= -1
        }
        if x == y {
            return false
        }
    }
    return true
}

func SetQueen(k int) {
    for i := 0; i < 8; i++ {
        if Issafe(k, i) != true {
            continue
        } else {
            ans[k] = i
            if k == 7 {
                for j := 0; j < 8; j++ {
                    z01.PrintRune(rune(ans[j] + 49))
                }
                z01.PrintRune('\n')
            } else {
                SetQueen(k + 1)
            }
        }
    }
