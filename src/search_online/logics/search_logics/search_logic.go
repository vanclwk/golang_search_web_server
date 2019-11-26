package search_logics

import (
	"github.com/gin-gonic/gin"
	"go_search_online/src/search_online/logics/parameter_logics"
)

type SearchResult struct {
	Status string `json:"status"`
	ResultCount int `json:"result_count"`
	ResultData []map[string]string `json:"result_data"`
}

func SearchLogic(ginContext *gin.Context) *SearchResult{
	parameter_logics.ParseParameters(ginContext)
	result := &SearchResult{
		Status:"WARN",
		ResultCount:0,
		ResultData:[]map[string]string{{"itemid":"1", "title": "test1"}, {"itemid":"2", "title": "test2"}},
	}
	return result
}
