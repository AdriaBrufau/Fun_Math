package fibonacci

import "fmt"

func Fibonacci() {
	fib_seq := make([]uint64, 0, 10)
	fib_seq = append(fib_seq, 1)
	fib_seq = append(fib_seq, 1)
	//var last_val uint64 = 0
	var total_sum uint64 = 0
	for i := 0; fib_seq[len(fib_seq)-1] < 4000000; i++ {

		new_val := fib_seq[len(fib_seq)-1] + fib_seq[len(fib_seq)-2]
		fib_seq = append(fib_seq, new_val)
		if new_val%2 == 0 {
			total_sum += new_val
		}
		fmt.Printf("%d \n %d\n %d", total_sum, i, fib_seq)
	}

}
