package handlers

import (
	"github.com/RakhimovAns/sportgrum/model"
	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "user registered", "user": user})
}
