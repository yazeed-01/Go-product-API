package controllers

import (
	"RESTAPI/initializers"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete a Product operation
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM productsGO WHERE id = $1`
	commandTag, err := initializers.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	if commandTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.Status(http.StatusNoContent)

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------

	// Additional code, if use model

	/*
		id := c.Param("id")
		initializers.DB.Delete(&models.Product{}, id)
		c.Status(http.StatusNoContent)
	*/

}
