package grader

import (
	tyrk8s "court-herald/util"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// New spins up a new kubernetes job for a submission
func New(c *gin.Context) {
	client, err := tyrk8s.GetClient()
	if err != nil {
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to connect to Kubernetes",
		})
	}
	pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to find kubernetes pods",
		})
	}
	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Grader Job Created",
		"pods":        pods.Items,
	})
}
