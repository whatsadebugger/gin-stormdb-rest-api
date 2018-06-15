package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	// Handle the index route
	r.POST("/address", createAddress)
	r.GET("/address", getEveryAddress)
	r.GET("/address/:id", getAddress)
	r.PUT("/address", updateAddress)
	r.DELETE("/address/:id", deleteAddress)
	r.POST("/address/upload", importAddressBook)
	r.GET("/addressbook", exportAddressBook)
}
