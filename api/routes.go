package api

import (
	"court-herald/api/grader"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	tyrgin "github.com/stevens-tyr/tyr-gin"
)

// SetUp is a function to set up the routes for plague doctor microservice.
func SetUp() *gin.Engine {
	server := tyrgin.SetupRouter()

	server.MaxMultipartMemory = 50 << 20

	// server.Use(tyrgin.Logger())
	server.Use(gin.Recovery())
	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	var graderEndpoints = []tyrgin.APIAction{
		tyrgin.NewRoute(grader.New, ":subid/new", tyrgin.POST),
		tyrgin.NewRoute(grader.Status, ":subid/status", tyrgin.GET),
	}

	tyrgin.AddRoutes(server, false, nil, "1", "grader", graderEndpoints)

	server.NoRoute(tyrgin.NotFound)

	return server
}
