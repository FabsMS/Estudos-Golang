package main

import (
	"fmt"
	"math"
	"math/rand"
)

func MathPackage() {
	fmt.Println("Número aleatório entre 0 e 100:", rand.Intn(100))
}

func SqrtFuncion() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func PiValue() {
	fmt.Println("O valor de Pi é:", math.Pi)
}

func Soma(x int, y int) { // Já que as duas variáveis tem o mesmo tipo, poderia tipar elas dessa maneira -> (x, y int)
	fmt.Println(x + y)
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
