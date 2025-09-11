package main

import "fmt"

var c, python, java bool

func Variables() {
	var i int
	fmt.Println(i, c, python, java)
}

var x, y int = 1, 2

func Initializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(x, y, c, python, java)
}
