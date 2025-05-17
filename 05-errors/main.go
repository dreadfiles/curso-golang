package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	stringNumber := "12x3"
	intNumber, err := strconv.Atoi(stringNumber)
	if err != nil {
		fmt.Println("Error:", err)
		//return
	}
	fmt.Println("Number:", intNumber)

	//Error from errors new
	divisionNew, err := divideNew(5, 0)
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	fmt.Println("divisionNew: ", divisionNew)

	//Error from frm errorf
	divisionFmt, err := divideFmt(5, 0)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println("divisionFmt: ", divisionFmt)

	//Controled and defined error
	valueMap, err := getValueItem(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("valueMap:", valueMap)

	//Defer
	deferFunction()

	//Defer with file creation
	deferFile()

	//Logs, with prefix, date and hour
	log.SetPrefix("Prefix Func: ")
	log.Println("Log Println")
	//log.Fatal("Finish the app")
	//log.Panic("Finish the app")

	//Panic division
	dividePanic(5, 0)

	//Comment for PR end of lesson 05
}

func divideNew(upNumber, downNumber int) (int, error) {
	if downNumber == 0 {
		return 0, errors.New("Error: zero division new")
	}
	return upNumber / downNumber, nil
}

func divideFmt(upNumber, downNumber int) (int, error) {
	if downNumber == 0 {
		return 0, fmt.Errorf("Error: zero division fmt")
	}
	return upNumber / downNumber, nil
}

func dividePanic(upNumber, downNumber int) {
	if downNumber == 0 {
		panic("zero division error")
	}
	fmt.Println("This line wouldn't be executed: ", upNumber, downNumber)
}

var errNotFound = errors.New("Error: item not found")

var itemsMap = map[int]string{
	1: "item1",
	2: "item2",
	3: "item3",
}

func getValueItem(key int) (string, error) {
	value, exists := itemsMap[key]
	if !exists {
		return "", errNotFound
	}
	return value, nil
}

func deferFunction() {
	// Defer: the next first line will be executed at the end
	defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)

	defer fmt.Println(4) //penultima
	defer fmt.Println(5) //antepenultima
	defer fmt.Println(6)
}

func deferFile() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello, Adrian"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
