package palindromes

import "fmt"

func Palindrome() {
	v_2 := 999
	reversed := 0
	remainder := 0
	max_pal := 100 * 100
	for i := v_2; i >= 100; i-- {
		for j := v_2; j >= 100; j-- {
			result := j * i
			MAX := result
			for result > 0 {
				remainder = result % 10
				reversed = reversed*10 + remainder
				result = result / 10
			}
			if MAX == reversed {
				if max_pal < reversed {
					max_pal = reversed
				}
			}
			result = 0
			reversed = 0
		}
	}
	fmt.Printf("%v \n", max_pal)
}
