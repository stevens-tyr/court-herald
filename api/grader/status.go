package grader

import (
	tyrk8s "court-herald/util"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Status gets the status of the Kubernetes Job
func Status(c *gin.Context) {
	client, err := tyrk8s.GetClient()

	jobsClient := client.BatchV1().Jobs("default")

	jobName := fmt.Sprintf("grader-job-%s", c.Param("subid"))
	graderJob, err := jobsClient.Get(jobName, metav1.GetOptions{})
	if err != nil {
		fmt.Println("code:", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to retrieve information about Job.",
			"err":         err,
		})
		return
	}

	podList, err := client.CoreV1().Pods("default").List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("controller-uid=%s", graderJob.Labels["controller-uid"]),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to retrieve Pod utilized by Job.",
			"err":         err,
		})
		return
	}

	if len(podList.Items) < 1 {
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "No job found for submission.",
			"err":         err,
		})
		return
	}

	jobPod := podList.Items[0]
	logOptions := apiv1.PodLogOptions{}
	req := client.CoreV1().Pods("default").GetLogs(jobPod.Name, &logOptions)
	reader, err := req.Stream()
	if err != nil {
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to retrieve log from Job Pod.",
			"err":         err,
		})
		return
	}
	defer reader.Close()
	podStdoutb64, err := ioutil.ReadAll(reader)
	if err != nil {
		c.JSON(500, gin.H{
			"status_code": 500,
			"message":     "Unable to obtain log from Job Pod.",
			"err":         err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"job":         graderJob,
		"pod":         jobPod,
		"logs":        fmt.Sprintf("%s", podStdoutb64),
	})
}
