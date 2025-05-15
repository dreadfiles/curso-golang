package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var name string
	name = "Adrian"
	fmt.Println(name)

	name = "Andres"
	fmt.Println(name)

	name = "Ordonez"
	fmt.Printf("Type: %T - Name: %s \n", name, name)

	object := "table"
	fmt.Println(object)

	age := 42
	fmt.Printf("Age: %d\n", age)

	count1, count2, count3 := 1, 2, 3
	fmt.Println(count1, count2, count3)

	var (
		firstName  = "Adrian"
		lastName   = "Ordonez"
		currentAge = 42
	)
	fmt.Println(firstName, lastName, currentAge)

	const pi = 3.141592
	fmt.Println(pi)

	const (
		x = 100
		y = 0b1010 // binario
		z = 0o12   // octal
		w = 0xFF   // hexadecimal
	)
	fmt.Println(x, y, z, w)

	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807

	fmt.Println(integer8, integer16, integer32, integer64)
	/* Utiliza los constantes del paquete math
	para saber cantidad de números que puede almacenar*/
	fmt.Println(math.MaxInt8, math.MaxInt8)

	// Ejemplo con Uint
	var integerUint8 uint8 = 255
	fmt.Println(integerUint8, math.MaxUint8)

	//Ejemplo de tipo flotantes
	var numberFloat32 float32 = 2.54
	fmt.Println(numberFloat32, math.MaxFloat32)

	//Ejemplo de Booleano
	var valueBool bool = true
	fmt.Println(valueBool)

	//Ejemplo de byte
	var a byte = 'a' // Asignar el valor decimal 97 (equivalente a "a" en ASCII)
	fmt.Println(a)   // Imprimir el valor de la variable b

	s := "hola"       // Crear una cadena de texto
	fmt.Println(s[0]) // Imprimir el valor del primer byte de la cadena s

	//Ejemplo de rune
	var r rune = '❤' // Asignar el valor de un caracter Unicode
	fmt.Println(r)   // Imprimir el valor de la variable r

	//Default values
	var defaultInt int
	var defaultUint uint
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultUint, defaultFloat64, defaultBool, defaultString)

	//Conversion de tipos
	var integer16_ int16 = 50
	var integer32_ int32 = 100
	fmt.Println(int32(integer16_) + integer32_)

	// Convertir una cadena a un número entero
	sta := "100"
	i, error := strconv.Atoi(sta)
	fmt.Println(i + 1)
	fmt.Println(error)

	// Convertir un número entero a una cadena
	nu := 42
	sta = strconv.Itoa(nu)
	fmt.Println("hello" + sta)

	//fmt package
	name = "ALex"
	age = 27

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)
	fmt.Println(greeting)

	//Operadores
	a2 := 10
	b2 := 3
	fmt.Println(a2 + b2) // 13
	fmt.Println(a2 - b2) // 7
	fmt.Println(a2 * b2) // 30
	fmt.Println(a2 / b2) // 3
	fmt.Println(a2 % b2) // 1

	// Ejemplo de potencias
	fmt.Println(math.Pow(2, 3)) // Imprime: 8
	fmt.Println(math.Sqrt(64))  // Imprime: 8
	fmt.Println(math.Cbrt(27))  // Imprime: 3
	fmt.Println(math.Pow10(2))  // Imprime: 100

	//Triangulo
	var lado1, lado2 float64
	const precision = 2

	lado1 = 5
	lado2 = 3

	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	area := (lado1 * lado2) / 2
	perimetro := lado1 + lado2 + hipotenusa

	fmt.Printf("\nÁrea: %.*f\n", precision, area)
	fmt.Printf("Perímetro: %.*f\n", precision, perimetro)
}
