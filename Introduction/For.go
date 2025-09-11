package main

import "fmt"

func ForTest() {
	sum := 0
	for i := 0; 1 < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
