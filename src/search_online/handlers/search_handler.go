package handlers

import (
	"go_search_online/src/search_online/logics/search_logics"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchHandler(context *gin.Context) {
	result := search_logics.SearchLogic(context)
	context.JSON(http.StatusOK, result)
}