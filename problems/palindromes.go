package problems

import (
	"fmt"
	"strconv"
)

func Palindromes() {

	c := make([]int, 1000000)
	sum := 0
	for number := range c {
		if checkIfPalindrome(strconv.Itoa(number), strconv.FormatInt(int64(number), 2)) {
			sum += int(number)
		}
	}

	fmt.Println(sum)
}

func checkIfPalindrome(n string, nBinary string) bool {
	var counter int = 0
	var binaryCounter int = 0
	strCopy, strBinaryCopy := make([]byte, len(n)), make([]byte, len(nBinary))
	for i, j := 0, len(n)-1; i <= len(n)-1; i, j = i+1, j-1 {
		strCopy[i] = n[j]
		if strCopy[i] == n[i] {
			counter++
		} else {
			counter = 0
		}
	}

	for i, j := 0, len(nBinary)-1; i <= len(nBinary)-1; i, j = i+1, j-1 {
		strBinaryCopy[i] = nBinary[j]
		if strBinaryCopy[i] == nBinary[i] {
			binaryCounter++
		} else {
			binaryCounter = 0
		}
	}

	if counter == len(n) && binaryCounter == len(nBinary) {

		return true
	} else {

		return false
	}
}
