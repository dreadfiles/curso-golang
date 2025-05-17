package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	//Time - if else
	t := time.Now()

	if t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}

	//Switch - runtime os
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Go run -> Linux.")
	case "windows":
		fmt.Println("Go run -> Windows.")
	case "darwin":
		fmt.Println("Go run -> macOS.")
	default:
		fmt.Println("Go run -> Otro SO.")
	}

	//For - pares impares
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//Llamado funcion
	hello("Adrian")

	//Llamado a funcion multiple retorno
	a, b := calc(3, 5)
	fmt.Println(a, b)

	//Numero aleatorio
	randomNumber := rand.Intn(100)
	fmt.Println(randomNumber)
}

func hello(name string) {
	fmt.Println("Hola, ", name)
}

func calc(a, b int) (int, int) {
	x := a + b
	y := a * b
	return x, y
}
