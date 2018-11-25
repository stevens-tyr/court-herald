package grader

import (
	"github.com/gin-gonic/gin"
)

// Status gets the status of the Kubernetes Job
func Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Status Routes",
		"param":       c.Param("subid"),
	})
}
