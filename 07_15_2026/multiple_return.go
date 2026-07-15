package main

import "fmt"

func MaxMin(number []int) (int,int){
	max := number[0]
	min := number[0]
	for _, v := range number {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}