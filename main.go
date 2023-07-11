package main

import (
	"github.com/gin-gonic/gin"
	"github.com/p97k/on-mark/controller"
	"github.com/p97k/on-mark/database"
	"github.com/p97k/on-mark/middleware"
	"github.com/p97k/on-mark/route"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controller.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	route.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/add-to-cart", app.AddToCart())
	router.GET("/remove-item", app.RemoveItem())
	router.GET("/cart-checkout", app.BuyFromCart())
	router.GET("/instant-buy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
