package stuff

import (
	"fmt"
)

func fib(n int) float64{
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}else if n == 2{
		return 1
	}
	return (fib(n-1)+fib(n-2))
}

func Sequence(n int){
	var array [2001]float64

	for i := 0; i<=n; i++{
		if (i == 1) || (i == 2) || (i==0){
			array[i] = fib(i)
			fmt.Printf("%d: %f\n", i, array[i])
		}else{
			array[i] = (array[i-1]+array[i-2])
			// fmt.Printf("%d: %.0f\n", i, array[i])
			fmt.Printf("NR %d: \t%f\n\n", i, array[i])
		}
	}
}
