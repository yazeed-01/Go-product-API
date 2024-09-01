package main

import (
	"RESTAPI/controllers"
	"RESTAPI/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	// CRUD routes;
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	// Auth routes

	r.Run()

}
