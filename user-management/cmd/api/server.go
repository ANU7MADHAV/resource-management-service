package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	r.GET("/healthcheck", app.healthCheck)
	return r
}
