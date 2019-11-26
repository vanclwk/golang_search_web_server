package parameter_logics

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_search_online/src/search_online/logics/logging_logics"
)

type Params struct {
	AppName string `form:"appname" json:"appname" xml:"appname" binding:"required"`
	AppId string `form:"appid" json:"appid" xml:"appid" binding:"required"`
	Time string `form:"appid" json:"appid" xml:"appid" binding:"required"`
	Sign string `form:"appid" json:"appid" xml:"appid" binding:"required"`
	Query string `form:"query" json:"query" xml:"query" binding:"required"`
	Cnt string `form:"cnt" json:"cnt" xml:"cnt"`
	Pos string `form:"pos" json:"pos" xml:"pos"`
	Filter string `form:"filter" json:"filter" xml:"filter"`
	Not_Filter string `form:"not_filter" json:"not_filter" xml:"not_filter"`
	Fuzzy string `form:"fuzzy" json:"fuzzy" xml:"fuzzy"`
	Sort string `form:"sort" json:"sort" xml:"sort"`
	Revert string `form:"revert" json:"revert" xml:"revert"`
	Range string `form:"range" json:"range" xml:"range"`
	Not_Range string `form:"not_range" json:"not_range" xml:"not_range"`
	Aggregation string `form:"aggregation" json:"aggregation" xml:"aggregation"`
}

func ParseParameters(ginContect *gin.Context) (Params, error) {
	var params Params
	var errorString string
	if err := ginContect.ShouldBind(&params); err != nil {
		logging_logics.ServiceLogger.Warn("Missing parameter " + err.Error())
		errorString = "missing parameter"
	}
	return params, errors.New(errorString)
}