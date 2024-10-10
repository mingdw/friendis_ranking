package controller

import (
	"friends_ranking/models"
	activityService "friends_ranking/service/activity/crud"
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
		StartTime: startTime,
		EndTime:   endTime,
	}
	activityService.List(activity, pageSize, limit, c)
}

func (index ActivityController) Add(c *gin.Context, title, desc, startTime, endTime string) {
	code := "AC" + time.Now().Format("20060102150405") + randNum(4)
	activity := &models.Activity{
		Code:      code,
		Title:     title,
		Desc:      desc,
		Status:    0,
		StartTime: startTime,
		EndTime:   endTime,
	}
	activityService.Add(activity, c)
}

func (index ActivityController) Update(c *gin.Context, title, desc, startTime, endTime string, id int) {
	activity := &models.Activity{
		Id:        id,
		Title:     title,
		Desc:      desc,
		StartTime: startTime,
		EndTime:   endTime,
	}
	activityService.Update(activity, c)
}

func (index ActivityController) Delete(c *gin.Context, ids []int) {
	activityService.Delete(ids, c)
}

// 生成固定位的随机数
func randNum(i int) string {
	r := ""
	for v := 0; v < i; v++ {
		r += strconv.Itoa(rand.Intn(10))
	}
	return r
}
