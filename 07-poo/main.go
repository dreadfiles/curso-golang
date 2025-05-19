package main

import (
	"fmt"
)

/* Book Object */
type Book struct {
	Title  string //public
	Author string //public
	pages  int    //private
}

/* Book Object Constructor */
func NewBook(pTitle string, pAuthor string, pPages int) *Book {
	return &Book{
		Title:  pTitle,
		Author: pAuthor,
		pages:  pPages,
	}
}

/* Book Object Method */
func (book *Book) PrintInfo() {
	fmt.Println("*******************")
	fmt.Println("Title: ", book.Title)
	fmt.Println("Author: ", book.Author)
	fmt.Println("Pages: ", book.pages)
	fmt.Println("*******************")
}

/* Book Object Setter */
func (book *Book) setPages(pPages int) {
	book.pages = pPages
}

/* Book Object Getter */
func (book *Book) getPages() int {
	return book.pages
}

/* TextBook Object */
type TextBook struct {
	//Book      Book //De esta forma no hereda
	Book      //De esta forma si hereda
	Editorial string
	Level     string
}

/* TextBook Object Constructor */
func NewTextBook(pTitle string, pAuthor string, pPages int, pEditorial string, pLevel string) *TextBook {
	return &TextBook{
		Book:      Book{pTitle, pAuthor, pPages},
		Editorial: pEditorial,
		Level:     pLevel,
	}
}

/* TextBook Object Method */
func (textBook *TextBook) PrintInfo() {
	fmt.Println("*******************")
	fmt.Println("Title: ", textBook.Title)
	fmt.Println("Author: ", textBook.Author)
	fmt.Println("Pages: ", textBook.pages)
	fmt.Println("Editorial: ", textBook.Editorial)
	fmt.Println("Level: ", textBook.Level)
	fmt.Println("*******************")
}

// Interfaces: este metodo debe ser implementado por varias clases
type Printable interface {
	PrintInfo()
}

// Polimorfismo: el metodo Print va a funcionar con todos los objetos que implementaron PrintInfo
func Print(printable Printable) {
	printable.PrintInfo()
}

func main() {
	//Basic struct
	var myBook Book = Book{
		Title:  "Moby Dick",
		Author: "Herman Melville",
		pages:  300,
	}
	myBook.PrintInfo()

	//Constructor
	newBook := NewBook("Book1", "Author1", 111)
	newBook.PrintInfo()
	//Private value
	newBook.setPages(987)
	fmt.Println("Pages Private: ", newBook.getPages())

	//TextBook heritance
	textBook := NewTextBook("TBT1", "TBA1", 654, "TBE1", "TBL1")
	textBook.PrintInfo()

	//Polimorfismo
	Print(newBook)
	Print(textBook)

	//Polimorfismo Animal
	dog := Dog{Name: "Kaiser"}
	cat := Cat{Name: "Lucumí"}
	MakeSound(dog)
	MakeSound(cat)

	//Multiple polimorfism
	animals := []Animal{
		Dog{Name: "Max"},
		Cat{Name: "Blue"},
		Dog{Name: "Lulú"},
		Cat{Name: "Lestat"},
	}
	for _, animal := range animals {
		animal.Sound()
	}
}

//OTHER INTERFACE

// Class Dog
type Dog struct {
	Name string
}

func (dog Dog) Sound() {
	fmt.Println("Guau Guau Guau - ", dog.Name)
}

// Class Cat
type Cat struct {
	Name string
}

func (cat Cat) Sound() {
	fmt.Println("Miau Miau Miau - ", cat.Name)
}

// Interface
type Animal interface {
	Sound()
}

// Porlimorfism
func MakeSound(animal Animal) {
	animal.Sound()
}
