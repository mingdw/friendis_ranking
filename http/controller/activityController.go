package controller

import (
	"fmt"
	"friends_ranking/config/globalConst"
	"friends_ranking/models"
	activityService "friends_ranking/service/activity/crud"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
}

func (index ActivityController) QueryList(c *gin.Context) {
	code := c.GetString(globalConst.ValidatorPrefix + "code")
	status := c.GetInt(globalConst.ValidatorPrefix + "status")
	startTime := c.GetString(globalConst.ValidatorPrefix + "startTime")
	endTime := c.GetString(globalConst.ValidatorPrefix + "endTime")
	limit := c.GetInt(globalConst.ValidatorPrefix + "limit")
	pageSize := c.GetInt(globalConst.ValidatorPrefix + "pageSize")
	fmt.Println("activity pageSize： ", pageSize, "; limit: ", limit, ", code: ", code)
	activity := &models.Activity{
		Code:      code,
		Status:    status,
		StartTime: startTime,
		EndTime:   endTime,
	}
	activityService.List(activity, 0, 10, c)

}
