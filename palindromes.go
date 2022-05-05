package main

import "fmt"

func main() {
	var x = populate(10)

	fmt.Print(x)
	//[0 0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

}

func populate(n int) []int {
	var a = make([]int, n)

	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	return a
}
