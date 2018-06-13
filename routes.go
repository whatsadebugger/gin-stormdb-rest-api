package main

func initializeRoutes() {

	// Handle the index route
	router.POST("/address", createAddress)
	router.GET("/address", getEveryAddress)
	router.GET("/address/:phonenumber", getAddress)
	router.PUT("/address/:phonenumber", updateAddress)
	router.DELETE("/address/:phonenumber", deleteAddress)
}
