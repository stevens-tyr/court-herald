package grader

import (
	tyrk8s "court-herald/util"

	"github.com/gin-gonic/gin"
)

// New spins up a new kubernetes job for a submission
func New(c *gin.Context) {
	jobName, err := tyrk8s.CreateJob(c.Param("subid"))
	if err != nil {
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to create Kubernetes Job",
			"error":       err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Grader Job Created",
		"jobName":     jobName,
	})
}
