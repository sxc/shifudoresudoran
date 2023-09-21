package main

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", IndexHandler)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHandler(c *gin.Context) {
	c.XML(http.StatusOK, Person{
		FirstName: "John",
		LastName:  "Appleseed",
	})
}
