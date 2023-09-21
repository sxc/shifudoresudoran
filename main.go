package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + name + "!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/:name", IndexHandler)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
