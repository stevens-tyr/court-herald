package grader

import (
	models "court-herald/models"
	tyrk8s "court-herald/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

// New spins up a new kubernetes job for a submission
func New(c *gin.Context) {
	var requestData models.RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{
			"status_code": 400,
			"message":     "Invalid Form Data",
			"error":       err,
		})
		return
	}

	newJob, err := tyrk8s.CreateJob(requestData)
	if err != nil {
		fmt.Println("ERROR", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to attach to Kubernetes Client",
			"error":       err,
		})
		return
	}

	fmt.Println(newJob.Name)

	c.JSON(201, gin.H{
		"status_code": 201,
		"message":     "Grader Job Created",
		"job":         newJob.Name,
	})
}
