package controllers

import (
	"RESTAPI/initializers"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a Product operation
func CreateProduct(c *gin.Context) {
	var input struct {
		// required data
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
	}

	// Bind data from request body to the input struct
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert product into database:
	// 1. write query
	query := `INSERT INTO productsGO (name, description, price) VALUES ($1, $2, $3) RETURNING id, name, description, price, created_at, updated_at`

	// 2. Execute the query
	row := initializers.DB.QueryRow(context.Background(), query, input.Name, input.Description, input.Price)
	var product struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product", "details": err.Error()})
		return
	}

	// Format time fields to string
	product.CreatedAt = product.CreatedAt.UTC()
	product.UpdatedAt = product.UpdatedAt.UTC()

	// Return the created product
	c.JSON(http.StatusOK, product)

	// ----------------------------------------------------------------------------
	// -----------------------------------------------------------------------------
	//additional code, if you use model:

	/*
		1. get the data
		var product struct {
			ID          int
			Name        string
			Description string
			Price       float64
			CreatedAt   time.Time
			UpdatedAt   time.Time
		}
		c.bind(&product)


		2. create the product
		post := models.Product{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}
		result := initializers.DB.Create(&post)
		if result.Error != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		3. return the created product
			c.JSON(http.StatusOK, gin.H{
				"product": product,
			})



	*/

}
