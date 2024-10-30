package prizeService

import (
	"lottery_annual/models"
	"lottery_annual/utils/response"

	"github.com/gin-gonic/gin"
)

/*
*
分页查询活动详情列表
*/
func DetailList(code, prizeName string, status, limit, pageSize int, c *gin.Context) {
	prizeModels := models.CreatePrizeFactory("")
	count, details := prizeModels.SelectList(code, prizeName, status, pageSize, limit)
	response.ReturnSuccess(c, details, count)
}

func Add(activity *models.Activity, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	if activityModels.Add(activity) {
		response.ReturnDefaultOk(c)
		return
	}
	response.ReturnDeFaultFail(c)
}

func Update(activity *models.Activity, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	if activityModels.Update(activity) {
		response.ReturnDefaultOk(c)
		return
	}
	response.ReturnDeFaultFail(c)
}

func Delete(ids []int, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	if activityModels.Delete(ids) {
		response.ReturnDefaultOk(c)
		return
	}
	response.ReturnDeFaultFail(c)
}

func UpdateStatus(id, status int, c *gin.Context) {
	activityModels := models.CreateActivityFactory("")
	if activityModels.UpdateStatus(id, status) {
		response.ReturnDefaultOk(c)
		return
	}
	response.ReturnDeFaultFail(c)
}
