package main

func initializeRoutes() {

	// Handle the index route
	router.POST("/address", createAddress)
	router.GET("/address", getEveryAddress)
	router.GET("/address/:id", getAddress)
	router.PUT("/address", updateAddress)
	router.DELETE("/address/:id", deleteAddress)
}
