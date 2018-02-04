package controllers

import (
	"github.com/gin-gonic/gin"
)

type ProjectController struct{}

func (ctrl ProjectController) All(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
