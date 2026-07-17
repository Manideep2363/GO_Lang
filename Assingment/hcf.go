package main

import "fmt"

func main() {
	fmt.Print("Enter N and k separated by comma: ")
	var n, k int
	fmt.Scanf("%d,%d", &n, &k)

	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	fmt.Printf("%d th factor of %d is: %d\n", k, n, factors[len(factors)-k])
}
