package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int16  `json:"age"`
	Location  string `json:"location"`
}

type loginInput struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func (app *application) register(c *gin.Context) {
	var input registerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (app *application) login(c *gin.Context) {
	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
