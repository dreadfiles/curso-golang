package main

import "fmt"

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
}
