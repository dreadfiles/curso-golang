package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album representa datos sobre un álbum musical.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// slice de albums para almacenar datos iniciales de álbumes.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responde con la lista de todos los álbumes como JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums agrega un álbum desde JSON recibido en el cuerpo de la solicitud.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Llama a BindJSON para vincular el JSON recibido a newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Agrega el nuevo álbum a la lista.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID encuentra el álbum cuyo valor de ID coincide con el parámetro id
// enviado por el cliente, y luego devuelve ese álbum como respuesta.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Recorre la lista de álbumes, buscando
	// un álbum cuyo valor de ID coincida con el parámetro.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

func main() {
	fmt.Println("api")

	// Crea una nueva instancia de Gin
	router := gin.Default()
	// Define la ruta GET para el endpoint "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Hola Mundo!",
		})
	})
	//GET albums
	router.GET("/albums", getAlbums)
	//GET album by ID
	router.GET("/albums/:id", getAlbumByID)
	//POST albums
	router.POST("/albums", postAlbums)
	// Inicia el servidor en el puerto 8080
	router.Run("localhost:8080")
}
