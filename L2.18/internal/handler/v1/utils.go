package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondOK(c *gin.Context, response any) {
	c.JSON(http.StatusOK, gin.H{"result": response})
}

func respondBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func respondBusinessError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
}

func respondInternalError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
