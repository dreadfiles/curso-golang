package main

import "fmt"

func main() {
	//Matrices
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	fmt.Println(a)

	var b = [5]int{10, 20, 30, 40, 50}
	fmt.Println(b)

	c := [...]int{10, 20, 30, 40, 50}
	fmt.Println(c)

	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

	//Matriz bidimensional
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, matriz[i])
	}
	fmt.Println("\nAll at once:", matriz)

	//Slices
	//Declaraciín e incializacion de una rebanada
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	// Crear una rebanada apartir de otra
	diasRebanada := diasSemana[0:5]
	// Motrar tambien la longitud y la capacidad
	fmt.Println(diasRebanada)
	fmt.Println("Longitud: ", len(diasRebanada))
	fmt.Println("Capacidad: ", cap(diasRebanada))

	// Agregar datos
	diasRebanada = append(diasRebanada, "Viernes")
	diasRebanada = append(diasRebanada, "Sabado")
	fmt.Println(diasRebanada)

	// Quitar elementos
	diasRebanada = append(diasRebanada[:2], diasRebanada[3:]...)
	fmt.Println(diasRebanada)

	// Crear una rebanada con make
	nombres := make([]string, 5, 10)
	fmt.Println(nombres)
	fmt.Println("Longitud: ", len(nombres))
	fmt.Println("Capacidad: ", cap(nombres))

	//Copy
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	rebanadas := copy(rebanada2, rebanada1)
	fmt.Println(rebanadas)
	fmt.Println(rebanada2)

	//Mapas
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	fmt.Println(colors["verde"])
	//Agregar nuevo elemento
	colors["negro"] = "#000000"
	fmt.Println(colors)
	//Verificar si existe
	_, ok := colors["violeta"]
	fmt.Println("Existe: ", ok)
	//Borrar elemento
	delete(colors, "verde")
	fmt.Println(colors)

	var p Persona
	p.Nombre = "Juan"
	p.Edad = 30
	p.Email = "juan@example.com"

	fmt.Println(p)

	r := Persona{
		Nombre: "Adrian",
		Edad:   42,
		Email:  "hola@mail.com",
	}
	fmt.Println(r)

	s := Persona{"Juan", 30, "juan@example.com"}
	fmt.Println(s)

	// Punteros
	var x int = 10
	var w *int = &x
	var y int = *w
	fmt.Println(y)

	//Llamado a funcion apuntando a struct con puntero
	r.DiHola()

	//Comentario al final de la unidad para PR
}

// ESTRUCTURAS
type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

func (p *Persona) DiHola() {
	fmt.Println("Hola, mi nombre es", p.Nombre)
}
