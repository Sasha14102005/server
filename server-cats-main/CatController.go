package main

import (
	"github.com/gin-gonic/gin"
)

func GetAllCats(c *gin.Context) {
	c.JSON(200, FindAllCats())
}

// Добавить котика.

func AddCat(c *gin.Context) {
	var cat Cat

	if err := c.BindJSON(&cat); err != nil {
		return
	}

	c.JSON(200, CreateCat(cat))
}
