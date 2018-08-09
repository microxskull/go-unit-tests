package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := setupRouter()
	app.Run(":3103")
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/ping", func(c *gin.Context) {
		// log.Println("[POST][/ping] start")
		// defer log.Println("[POST][/ping] finish")
		bodyMap := map[string]interface{}{}
		c.ShouldBindJSON(&bodyMap)
		// log.Printf("[POST][/ping] request: %+v", bodyMap)
		// Get data in body
		message := ""
		if val, ok := bodyMap["message"]; ok {
			message = " " + val.(string)
		}
		c.JSON(http.StatusOK, gin.H{"message": "pong" + message})
	})
	return r
}
