package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/healthcheck", app.healthCheck)

		v1.POST("/users/register", app.register)
		v1.POST("/users/login", app.login)
	}

	return r
}
