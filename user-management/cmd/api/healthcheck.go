package main

import "github.com/gin-gonic/gin"

func (app *application) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"Port": app.config.port, "version": version})
}
