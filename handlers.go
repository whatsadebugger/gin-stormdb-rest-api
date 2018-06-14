package main

import (
	"encoding/csv"
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

func importAddressBook(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	f, err := file.Open()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer f.Close()

	var addressbook []address

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	for _, line := range lines {
		addressbook = append(addressbook,
			address{FirstName: line[0],
				LastName: line[1],
				Email:    line[2],
				Phone:    line[3],
			})
	}

	for _, ad := range addressbook {
		err = database.Save(&ad)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
	}

}

func exportAddressBook(c *gin.Context) {
	var book []address
	if err := database.All(&book); err == storm.ErrNotFound {
		publicError(c, 404, storm.ErrNotFound)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	header := c.Writer.Header()
	header["Content-type"] = []string{"text/csv"}
	header["Content-Disposition"] = []string{"attachment; filename=backup.csv"}

	wr := csv.NewWriter(c.Writer)

	for _, v := range book {
		if err := wr.Write([]string{v.FirstName, v.LastName, v.Email, v.Phone}); err != nil {
			c.AbortWithError(500, err)
			return
		}
	}

	wr.Flush()
}
