package main

import (
	fireGO "github.com/JEpifanio90/FireGO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	auth := r.Group("/auth")
	auth.Use(fireGO.AuthMiddleware())
	auth.GET("/", func(c *gin.Context) {
		//claims := fireGO.ExtractClaims(c)
		c.String(http.StatusOK, "success")
	})
	r.Run()
}
