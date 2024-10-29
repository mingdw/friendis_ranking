package controller

import (
	"lottery_annual/config/globalConst"
	"lottery_annual/models"
	activityService "lottery_annual/service/activity/crud"
	"time"

	"github.com/gin-gonic/gin"
)

type PrizeController struct {
}

func (index PrizeController) QueryList(c *gin.Context, code, prizeName string, status, limit, pageSize int) {
	details := &models.PrizeDetails{}

	details.SelectList(code, prizeName, status, pageSize, limit)
}

func (index PrizeController) Add(c *gin.Context, title, desc, startTime, endTime string) {
	code := "AC" + time.Now().Format("20060102150405") + randNum(4)
	t1, _ := time.ParseInLocation(globalConst.DateFormat, startTime, time.Local) //这里按照当前时区转
	t2, _ := time.ParseInLocation(globalConst.DateFormat, endTime, time.Local)   //这里按照当前时区转
	activity := &models.Activity{
		Code:      code,
		Title:     title,
		Desc:      desc,
		Status:    1,
		StartTime: t1,
		EndTime:   t2,
	}
	activityService.Add(activity, c)
}

func (index PrizeController) Update(c *gin.Context, title, desc, startTime, endTime string, id int) {
	activity := &models.Activity{
		Id:        id,
		Title:     title,
		Desc:      desc,
		StartTime: time.Time{},
		EndTime:   time.Time{},
	}
	if startTime != "" {
		t1, _ := time.ParseInLocation(globalConst.DateFormat, startTime, time.Local) //这里按照当前时区转
		activity.StartTime = t1
	}
	if endTime != "" {
		t2, _ := time.ParseInLocation(globalConst.DateFormat, endTime, time.Local) //这里按照当前时区转
		activity.EndTime = t2
	}
	activityService.Update(activity, c)
}

func (index PrizeController) Delete(c *gin.Context, ids []int) {
	activityService.Delete(ids, c)
}

func (index PrizeController) UpdateStatus(c *gin.Context, id, status int) {
	activityService.UpdateStatus(id, status, c)
}
