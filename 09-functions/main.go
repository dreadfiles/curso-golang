package main

import "fmt"

func main() {
	//Variadic function
	fmt.Println("Sum1: ", sum("Sum1", 1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println("Sum2: ", sum("Sum2", 1, 2, 3, 4, 5))

	//Variadic function, different type of arguments
	printData("Hola", 42, true, 3.14)

	//Recursive function
	fmt.Println("Factorial: ", factorial(5))

	//Anonime function
	greeting := func(name string) {
		fmt.Println("Hello: ", name)
	}
	greeting("Adrian")

	//Value function, the previous func as variable
	sayHello := func(name string, variableFunc func(string)) {
		variableFunc(name)
	}
	sayHello("Andres", greeting)

	//Other example of value function
	value1 := multiply(twice, 5)
	value2 := multiply(threeTimes, 5)
	fmt.Println("Result values: ", value1, value2)

	//Higher-level functions, orden superior
	value3 := double(addOne, 3)
	fmt.Println("Value3: ", value3)

	//Closure funcitons
	nextInt := incrementer()
	fmt.Println("Incrementer1: ", nextInt())
	fmt.Println("Incrementer2: ", nextInt())
	fmt.Println("Incrementer3: ", nextInt())
}

// Variadic function, n arguments
func sum(title string, numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println("Title: ", title)
	return total
}

// Variadic function, n arguments, different type
func printData(data ...interface{}) {
	for _, item := range data {
		fmt.Printf("Type: %T - Value: %v\n", item, item)
	}
}

// Recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Other example of value function
func twice(value int) int {
	return value * 2
}

func threeTimes(value int) int {
	return value * 3
}

func multiply(variableFunc func(int) int, value int) int {
	return variableFunc(value)
}

// Higher-level functions, orden superior
func addOne(value int) int {
	return value + 1
}

func double(variableFunc func(int) int, value int) int {
	return variableFunc(value * 2)
}

// Closure functions
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
