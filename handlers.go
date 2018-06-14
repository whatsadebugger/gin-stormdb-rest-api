package main

import (
	"errors"
	"fmt"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"strconv"
)

func createAddress(c *gin.Context) {
	ad := address{}

	if err := c.ShouldBindJSON(&ad); err != nil {
		publicError(c, 400, err)
	} else if err = database.Save(&ad); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, ad)
	}
}

func getAddress(c *gin.Context) {
	id := c.Param("id")
	var ad address
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if err := database.One("ID", aid, &ad); err == storm.ErrNotFound {
		publicError(c, 404, storm.ErrNotFound)
	} else if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, ad)
	}
}

func getEveryAddress(c *gin.Context) {
	var book []address
	if err := database.All(&book); err == storm.ErrNotFound {
		publicError(c, 404, storm.ErrNotFound)
	} else if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, book)
	}
}

func updateAddress(c *gin.Context) {
	var ad address

	if err := c.ShouldBindJSON(&ad); err != nil {
		publicError(c, 400, err)
	} else if err = database.Update(&ad); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, ad)
	}
	fmt.Println(ad)
}

func deleteAddress(c *gin.Context) {
	id := c.Param("id")

	if ad, err := findAddress(id); err == storm.ErrNotFound {
		c.AbortWithError(404, errors.New("address entry not found"))
	} else if err != nil {
		c.AbortWithError(500, err)
	} else if err = database.DeleteStruct(&ad); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.String(200, "Success")
	}
}

// func importAddressBook(c *gin.Context) {}

// func exportAddressBook(c *gin.Context) {}
