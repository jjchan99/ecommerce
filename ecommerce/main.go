package main

import

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
    router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addToCart", app.AddToCart())
	router.GET("removeItem", app.RemoveItem())
    router.GET("/cartCheckout", app.BuyFromCart())
	router.GET("/instantBuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}