package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

func main() {
	var cats []Cat

	pgConnect := PostgresConnect()
	if pgConnect == nil {
		log.Println("Отсутсвует подключение к БД!")
	}
	pgConnect.Close()

	r := gin.Default()

	r.POST("/api/cat/add", func(c *gin.Context) {
		var cat Cat

		if err := c.BindJSON(&cat); err != nil {
			return
		}

		cat.ID = uuid.NewString()

		cats = append(cats, cat)

		c.JSON(200, cat)
	})

	r.GET("/api/cats", func(c *gin.Context) {
		c.JSON(200, cats)
	})

	r.GET("/api/cat/:id", func(c *gin.Context) {
		id := c.Param("id")

		var cat Cat

		for _, ct := range cats {
			if ct.ID == id {
				cat = ct
			}
		}

		c.JSON(200, cat)
	})

	r.DELETE("/api/cat/:id", func(c *gin.Context) {
		id := c.Param("id")

		var index int
		var cat Cat

		for i, ct := range cats {
			if ct.ID == id {
				index = i
				cat = ct
			}
		}

		cats = append(cats[:index], cats[index+1:]...)

		c.JSON(200, cat)
	})

	r.PUT("/api/cat/:id", func(c *gin.Context) {
		var cat Cat
		id := c.Param("id")

		if err := c.BindJSON(&cat); err != nil {
			return
		}

		for i, ct := range cats {
			if ct.ID == id {
				cats[i].Color = cat.Color
				cats[i].Name = cat.Name
				cats[i].IsStrip = cat.IsStrip
				cat = cats[i]
			}
		}

		c.JSON(200, cat)

	})

	r.Run("0.0.0.0:8888")

}
