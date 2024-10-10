package activityService

import (
	"friends_ranking/models"
	"friends_ranking/utils/response"

	"github.com/gin-gonic/gin"
)

/*
*
分页查询活动详情列表
*/
func List(activity *models.Activity, limit, pageSize int, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	count, details := activityModels.Show(activity.Code, activity.StartTime, activity.EndTime, activity.Status, limit, pageSize)
	response.ReturnSuccess(c, details, count)
}

func Add(activity *models.Activity, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	if activityModels.Add(activity) {
		response.ReturnDefaultOk(c)
	}
	response.ReturnDeFaultFail(c)
}

func Update(activity *models.Activity, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	if activityModels.Update(activity) {
		response.ReturnDefaultOk(c)
	}
	response.ReturnDeFaultFail(c)
}

func Delete(ids []int, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	if activityModels.Delete(ids) {
		response.ReturnDefaultOk(c)
	}
	response.ReturnDeFaultFail(c)
}
