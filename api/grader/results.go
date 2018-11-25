package grader

import (
	"github.com/gin-gonic/gin"
)

// Results gets the results of the grader (stdout for now)
func Results(c *gin.Context) {
	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Results Route",
		"param":       c.Param("subid"),
	})
}
