package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}
	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/codebreaker/setup/:number", func(c *gin.Context) {
		number := c.Param("number")
		setCode(number)
		c.String(http.StatusOK, "Hello %s have configured", number)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/codebreaker/guest/:number", func(c *gin.Context) {
		name := c.Param("number")
		result := validateCode(name)
		c.String(http.StatusOK, "answer "+result)
	})

	router.Run(":" + port)
}
