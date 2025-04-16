package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Route
	router := gin.Default()

	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
