package controllers

import (
	"RESTAPI/initializers"
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Get All Products operation
func GetAllProducts(c *gin.Context) {
	// 1. write query
	query := `SELECT id, name, description, price, created_at, updated_at FROM productsGO`

	// 2. Execute the query
	rows, err := initializers.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	// 3. close the result
	defer rows.Close()

	// 4. Array for products to store the results
	var products []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// 5. Loop through the results
	for rows.Next() {
		var product struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			Price       float64   `json:"price"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan product"})
			return
		}

		// Add product to products array
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching results"})
		return
	}

	c.JSON(http.StatusOK, products)

	// ----------------------------------------------------------------------------
	// -----------------------------------------------------------------------------

	// Additional code, if use model

	/*
		var products []models.Product

		initializers.DB.Find(&products)

		c.JSON(http.StatusOK, products)



	*/

}

// Get a Product by ID
func GetProductByID(c *gin.Context) {

	/// 1. Take the id from the url
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// 2. Write query
	query := `SELECT id, name, description, price, created_at, updated_at FROM productsGO WHERE id = $1`

	// 3. Execute the query
	row := initializers.DB.QueryRow(context.Background(), query, id)

	//4. var product to store the result
	var product struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, product)

	// ----------------------------------------------------------------------------
	// -----------------------------------------------------------------------------

	// Additional code, if use model

	/*
		id := c.Param("id")
		var product models.Product
		initializers.DB.First(&product, id)
		c.JSON(http.StatusOK, product)


	*/

}
