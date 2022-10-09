package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	//lastOccurred := make(map[byte]int)
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("jiaoyixia"))
	fmt.Println(lengthOfNonRepeatingSubStr("butanle"))
	fmt.Println(lengthOfNonRepeatingSubStr("教一下求求你"))
	fmt.Println(lengthOfNonRepeatingSubStr("今天是个好日子"))
}
