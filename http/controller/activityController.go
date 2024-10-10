package controller

import (
	"fmt"
	"friends_ranking/models"
	activityService "friends_ranking/service/activity/crud"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
}

func (index ActivityController) QueryList(c *gin.Context, code, startTime, endTime string, status, limit, pageSize int) {
	fmt.Println("code: ", code, "; status: ", status, "; startTime: ", startTime, "; endTime: ", endTime, "; limit: ", limit, "; pageSize: ", pageSize)
	activity := &models.Activity{
		Code:      code,
		Status:    status,
		StartTime: startTime,
		EndTime:   endTime,
	}
	activityService.List(activity, pageSize, limit, c)
}

func (index ActivityController) Add(c *gin.Context, title, desc, startTime, endTime string) {
	//time := time.Now().Format("20060102150405")

}

func (index ActivityController) Update(c *gin.Context, title, desc, startTime, endTime string, id int) {
	//gengxin
}

func (index ActivityController) Delete(c *gin.Context, ids []int) {

}
