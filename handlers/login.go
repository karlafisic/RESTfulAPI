package handlers

import (
	"karla/restfulapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler prima korisničko ime i lozinku, generira JWT token
func LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if credentials.Username == "admin" && credentials.Password == "password" {
		token, err := utils.GenerateToken(credentials.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
