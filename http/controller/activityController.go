package controller

import (
	"friends_ranking/config/globalConst"
	"friends_ranking/models"
	activityService "friends_ranking/service/activity/crud"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
}

func (index IndexController) QueryList(c *gin.Context) {
	code := c.GetString(globalConst.ValidatorPrefix + "code")
	status := c.GetInt(globalConst.ValidatorPrefix + "status")
	startTime := c.GetString(globalConst.ValidatorPrefix + "startTime")
	endTime := c.GetString(globalConst.ValidatorPrefix + "endTime")

	limit := c.GetInt(globalConst.ValidatorPrefix + "limit")
	pageSize := c.GetInt(globalConst.ValidatorPrefix + "pageSize")
	activity := &models.Activity{
		Code:      code,
		Status:    status,
		StartTime: startTime,
		EndTime:   endTime,
	}
	activityService.List(activity, limit, pageSize, c)

}
