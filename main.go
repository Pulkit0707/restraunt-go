package main

import (
	"os"
	"restraunt-go/database"
	"restraunt-go/middleware"
	"restraunt-go/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.MenuRoutes(router)
	routes.CartRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":"+port)
}