package main

import (
	"fmt"
	"go_search_online/src/search_online/handlers"
	"github.com/gin-gonic/gin"
	"go_search_online/src/search_online/logics/config_logics"
	"go_search_online/src/search_online/logics/logging_logics"
)

func init() {
	fmt.Printf("Starting server on port %s...\n", config_logics.GlobalConfig.GetString("service.port"))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// search router
	logging_logics.ServiceLogger.Info("Starting server...")
	router.GET("/search/:app", handlers.SearchHandler)
	_ = router.Run(":" + config_logics.GlobalConfig.GetString("service.port"))
}