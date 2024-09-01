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

// Update a Product operation
func UpdateProduct(c *gin.Context) {
	// 1. Take the id from the url
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// 2. Validate the input and variable to store the input
	var input struct {
		Name        *string  `json:"name"`
		Description *string  `json:"description"`
		Price       *float64 `json:"price"`
	}

	// 3. Bind the input to the struct
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. Write query
	query := `UPDATE productsGO SET updated_at = CURRENT_TIMESTAMP`
	params := []interface{}{}
	paramCount := 1

	// 5. Check if input is not nil
	if input.Name != nil {
		query += `, name = $` + strconv.Itoa(paramCount)
		params = append(params, *input.Name)
		paramCount++
	}
	if input.Description != nil {
		query += `, description = $` + strconv.Itoa(paramCount)
		params = append(params, *input.Description)
		paramCount++
	}
	if input.Price != nil {
		query += `, price = $` + strconv.Itoa(paramCount)
		params = append(params, *input.Price)
		paramCount++
	}

	// 6. Add WHERE clause
	query += ` WHERE id = $` + strconv.Itoa(paramCount) + ` RETURNING id, name, description, price, created_at, updated_at`

	params = append(params, id)

	// 7. Execute the query
	row := initializers.DB.QueryRow(context.Background(), query, params...)

	// 8. var product to store the result
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, product)

	// ----------------------------------------------------------------------------
	// ----------------------------------------------------------------------------

	// Additional code, if use model
	/*
		id := c.Param("id")

		var product struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			Price       float64   `json:"price"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		}

		c.Bind(&product)


		var post models.Product
		initializers.DB.First(&post, id)

		initializers.DB.Model(&post).Updates(models.product
		{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}
		)

		c.JSON(http.StatusOK, gin.H{
				"product": product,
			})

	*/

}
