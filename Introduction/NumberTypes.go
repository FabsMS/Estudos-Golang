package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false     // tipo booleano
	MaxInt uint64     = 1<<64 - 1 // valor máximo de um int 64 bits = 2⁶⁴ - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func NumberTypes() {
	fmt.Printf("Type: %T | Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T | Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T | Value: %v\n", z, z)
}
