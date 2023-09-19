package primal

import (
	"fmt"
	"math"
)

func Largest_Primal() {

	value := 600851475143
	list_primals := make([]int, 0)
	for i := 2; i <= 100000000000; i++ {
		for value%i == 0 {
			value = int(math.Floor(float64(value) / float64(i)))
			list_primals = append(list_primals, i)
			if value == 1 || value == i {
				fmt.Printf("%d", i)
			}
		}
	}
	fmt.Printf("%d", list_primals)
}
