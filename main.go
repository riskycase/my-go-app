package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	
	// Return a preset JSON on receiving a get request at /hello
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message" : "Hello world!" })
	})

	//Construct and send a JSON greeting the person in the parameter
	router.GET("/greet/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message" : "Hello " + c.Param("name") + "!"})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}