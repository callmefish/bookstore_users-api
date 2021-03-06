package app

import (
	"github.com/callmefish/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about tot start the application...")
	router.Run(":8080")
}
