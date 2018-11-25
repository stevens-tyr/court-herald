package api

import (
	"court-herald/api/grader"

	"github.com/gin-gonic/gin"
	tyrgin "github.com/stevens-tyr/tyr-gin"
)

// SetUp is a function to set up the routes for plague doctor microservice.
func SetUp() *gin.Engine {
	server := tyrgin.SetupRouter()

	server.MaxMultipartMemory = 50 << 20

	//server.Use(tyrgin.Logger())
	server.Use(gin.Recovery())

	var graderEndpoints = []tyrgin.APIAction{
		tyrgin.NewRoute(grader.New, ":subid/new", false, tyrgin.POST),
		tyrgin.NewRoute(grader.Results, ":subid/results", false, tyrgin.GET),
		tyrgin.NewRoute(grader.Status, ":subid/status", false, tyrgin.GET),
	}

	tyrgin.AddRoutes(server, nil, "1", "grader", graderEndpoints)

	server.NoRoute(tyrgin.NotFound)

	return server
}
