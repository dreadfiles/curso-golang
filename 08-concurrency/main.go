package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//Se crea un canal para la concurrencia, de tipo string por las url's
	channel := make(chan string)

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		//con la palabra go antes de una funcion se llama a goroutines o concurrencia
		//se adiciona el parametro para el channel de concurrencia
		go checkAPI(api, channel)
	}

	//Leer los datos del canal, esto solo muestra la primera respuesta
	//fmt.Println(<-channel)
	//Si se adiciona esta linea, entonces muestra la siguiente respuesta
	//fmt.Println(<-channel)

	//Si se ponen 6, funciona para todos las API
	//fmt.Println(<-channel)
	//fmt.Println(<-channel)
	//fmt.Println(<-channel)
	//fmt.Println(<-channel)

	//Si se pone una de mas, el programa se bloquea, queda esperando
	//fmt.Println(<-channel)

	//de esta forma se garantiza que no se pase el numero de respuestas y falle
	for i := 0; i < len(apis); i++ {
		fmt.Println(<-channel)
	}

	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
}

// La funcion debe tener un parametro tipo chan para el canal
func checkAPI(api string, channel chan string) {
	_, err := http.Get(api)
	if err != nil {
		channel <- fmt.Sprintf("ERROR: %s no access\n", api)
		return
	}
	channel <- fmt.Sprintf("Success: %s OK\n", api)
}
