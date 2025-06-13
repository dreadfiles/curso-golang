package main

import (
	"dreadfiles/curso-golang/14-api-mysql/database"
	"dreadfiles/curso-golang/14-api-mysql/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting 14-api-mysql")

	// Connect to the MySQL database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("unable to connect to the database: ", err)
	}
	defer db.Close()

	// Initialize the product handler with the database connection
	productHandler := handler.NewProductHandler(db)

	// Create a new Gin router
	router := gin.Default()

	// Define the routes for product operations
	router.GET("/products", productHandler.GetProductsHandler)
	router.GET("/products/:id", productHandler.GetProductByIDHandler)
	router.POST("/products", productHandler.CreateProductHandler)
	router.PUT("/products", productHandler.UpdateProductHandler)
	router.DELETE("/products/:id", productHandler.DeleteProductHandler)

	router.Run(":8080")
}
