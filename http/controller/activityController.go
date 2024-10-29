package controller

import (
	"lottery_annual/config/globalConst"
	"lottery_annual/models"
	activityService "lottery_annual/service/activity/crud"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
}

func (index ActivityController) QueryList(c *gin.Context, code, startTime, endTime string, status, limit, pageSize int) {
	activity := &models.Activity{
		Code:      code,
		Status:    status,
		StartTime: time.Time{},
		EndTime:   time.Time{},
	}
	if startTime != "" {

		t1, _ := time.Parse(globalConst.DateFormat, startTime) //这里按照当前时区转
		activity.StartTime = t1
	}
	if endTime != "" {
		t2, _ := time.Parse(globalConst.DateFormat, endTime) //这里按照当前时区转
		activity.EndTime = t2
	}
	activityService.List(activity, pageSize, limit, c)
}

func (index ActivityController) Add(c *gin.Context, title, desc, startTime, endTime string) {
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

func (index ActivityController) Update(c *gin.Context, title, desc, startTime, endTime string, id int) {
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

func (index ActivityController) Delete(c *gin.Context, ids []int) {
	activityService.Delete(ids, c)
}

func (index ActivityController) UpdateStatus(c *gin.Context, id, status int) {
	activityService.UpdateStatus(id, status, c)
}

// 生成固定位的随机数
func randNum(i int) string {
	r := ""
	for v := 0; v < i; v++ {
		r += strconv.Itoa(rand.Intn(10))
	}
	return r
}
