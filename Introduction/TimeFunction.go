package main

import (
	"fmt"
	"time"
)

func TimeFunction() {
	fmt.Println("Teste de função de tempo")
	fmt.Println("Data e hora atual:", time.Now())
}

func WeekdayFunction() {
	fmt.Println("Quando vai ser o próximo sábado?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Hoje é sábado.")
	case today + 1:
		fmt.Println("Amanhã.")
	case today + 2:
		fmt.Println("Em dois dias.")
	default:
		fmt.Println("Ainda falta um tempo...")
	}
}

func HourFunction() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
