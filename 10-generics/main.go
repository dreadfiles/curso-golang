package main

import (
	"fmt"
	//"golang.org/x/exp/constraints"
)

func main() {
	//Variadic function any
	PrintList("Alex", "Roel", "Juan", "Pedro")
	PrintList(100, 456, 789, 456, 452)
	PrintList("Hola", 452, 4.25, true)

	//Variadic function int
	fmt.Println("Sum: ", Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	//Variadic function int, type restriction
	fmt.Println("SumTypeRestriction: ", SumTypeRestriction(1, 2, 3, 4, 5, 6, 7, 8, 9))

	//Variadic function int float64, union restriction
	fmt.Println("SumUnionRestriction: ", SumUnionRestriction(100, 456, 789, 456, 452))
	fmt.Println("SumUnionRestriction: ", SumUnionRestriction(4.5, 5.5, 7.5))

	//Variadic function ~int ~float64, proximity restriction
	var num1 integer = 100
	var num2 integer = 300
	fmt.Println("SumProxRestriction: ", SumProxRestriction(100, 456, 789, 456, 452))
	fmt.Println("SumProxRestriction: ", SumProxRestriction(14.5, 5.5, 7.5))
	fmt.Println("SumProxRestriction: ", SumProxRestriction(num1+num2))

	//Variadic function, custom restriction
	fmt.Println("SumCustomRestriction: ", SumCustomRestriction(100, 456, 789, 456, 452))
	fmt.Println("SumCustomRestriction: ", SumCustomRestriction[float32](14.5, 5.5, 7.5))
	fmt.Println("SumCustomRestriction: ", SumCustomRestriction(num1+num2))

	//Variadic function, constraints library
	//fmt.Println("SumConstraintsLibrary: ", SumConstraintsLibrary(100, 456, 789, 456, 452))
	//fmt.Println("SumConstraintsLibrary: ", SumConstraintsLibrary[float32](14.5, 5.5, 7.5))
	//fmt.Println("SumConstraintsLibrary: ", SumConstraintsLibrary(num1+num2))

	// Operator constrictions, only comparable types
	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Exists: ", Exists(strings, "a"))
	fmt.Println("Exists: ", Exists(strings, "f"))
	fmt.Println("Exists: ", Exists(numbers, 4))
	fmt.Println("Exists: ", Exists(numbers, 5))

	// Operator constrictions, only ordered, usirng library
	fmt.Println(Filter(numbers, func(value int) bool { return value > 2 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

	// Generic structures
	product1 := Product[uint]{1, "Shoes", 50}
	product2 := Product[string]{"sadad-ad5d4-asd", "Other Shoes", 100}
	fmt.Println("Products: ", product1, product2)
}

// Variadic function any
func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
	fmt.Println("******************")
}

// Variadic function int
func Sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

// Variadic function int, type restriction
func SumTypeRestriction[T int](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Variadic function int - float64, union restriction
func SumUnionRestriction[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Variadic function ~int ~float64, proximity restriction
type integer int

func SumProxRestriction[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Variadic function, custom restriction
type Integer interface {
	~int | ~float64 | ~float32
}

func SumCustomRestriction[T Integer](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Variadic function, constraints library
/*
func SumConstraintsLibrary[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
*/

// Operator constrictions, only comparable types
func Exists[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Operator constrictions, only ordered, usirng library
func Filter /*[T constraints.Ordered]*/ [T comparable](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

// Generic structures
type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}
