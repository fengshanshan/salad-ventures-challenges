package main

import "fmt"

func main()  {
	var n int64
	n = 100
	s1 := sum_to_n_a(n)
	fmt.Printf("sum_to_n_a result=%d\n", s1)
	s2 := sum_to_n_b(n)
	fmt.Printf("sum_to_n_b result=%d\n", s2)
	s3 := sum_to_n_c(n)
	fmt.Printf("sum_to_n_c result=%d\n", s3)
}

func sum_to_n_a(n int64) int64 {
	return n*(n+1)/2
}

func sum_to_n_b(n int64) int64 {
	var i, sum int64
	for i=1; i<=n; i++ {
		sum += i
	}
	return sum
}

func sum_to_n_c(n int64) int64 {
	return f(n);
}

func f(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n + f(n-1)
}
