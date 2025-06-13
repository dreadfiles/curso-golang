package handler

import (
	"database/sql"
	"dreadfiles/curso-golang/14-api-mysql/dao"
	"dreadfiles/curso-golang/14-api-mysql/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles requests related to products
type ProductHandler struct {
	db  *sql.DB
	dao *dao.ProductDAO
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	dao := dao.NewProductDAO(db)
	return &ProductHandler{db: db, dao: dao}
}

func (handler *ProductHandler) GetProductsHandler(context *gin.Context) {
	products, err := handler.dao.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error retrieving products: " + err.Error()})
		return
	}

	if len(products) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "no products found"})
		return
	}

	context.IndentedJSON(http.StatusOK, products)
}

func (handler *ProductHandler) GetProductByIDHandler(context *gin.Context) {
	id := context.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := handler.dao.GetByID(parseID)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error retrieving product: " + err.Error()})
	}

	if product == nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, product)
}

func (handler *ProductHandler) CreateProductHandler(context *gin.Context) {
	var product model.Product
	if err := context.BindJSON(&product); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid product data"})
		return
	}

	if err := handler.dao.Create(&product); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error creating product: " + err.Error()})
		return
	}

	context.IndentedJSON(http.StatusCreated, product)
}

func (handler *ProductHandler) UpdateProductHandler(context *gin.Context) {
	var product model.Product
	if err := context.BindJSON(&product); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid product data"})
		return
	}

	if err := handler.dao.Update(&product); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error updating product: " + err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, product)
}

func (handler *ProductHandler) DeleteProductHandler(context *gin.Context) {
	id := context.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	if err := handler.dao.Delete(parseID); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error deleting product: " + err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
