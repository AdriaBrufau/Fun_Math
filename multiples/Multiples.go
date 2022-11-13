package multiples

import "fmt"

func Multiples_go() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}

	fmt.Print(sum)
}
