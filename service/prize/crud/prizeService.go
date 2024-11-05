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

func Delete(ids []int, c *gin.Context) {
	prizeModels := models.CreatePrizeFactory("")
	if prizeModels.Delete(ids) {
		response.ReturnDefaultOk(c)
		return
	}
	response.ReturnDeFaultFail(c)
}

func SelectByAcId(acId int, c *gin.Context) {
	prizeModels := models.CreatePrizeFactory("")
	count, result := prizeModels.SelectByAcId(acId)
	response.ReturnSuccess(c, result, count)
}

func Add(c *gin.Context, acid, itemCode, prizeNum, isRepeat int, prizeName string) {
	prizeModels := models.CreatePrizeFactory("")
	if prizeModels.Add(acid, itemCode, prizeNum, isRepeat, prizeName) {
		response.ReturnDefaultOk(c)
		return
	}
	response.ReturnDeFaultFail(c)
}

func Update(c *gin.Context, id, acid, itemCode, prizeNum, isRepeat int, prizeName string) {
	prizeModels := models.CreatePrizeFactory("")
	if prizeModels.Update(id, acid, itemCode, prizeNum, isRepeat, prizeName) {
		response.ReturnDefaultOk(c)
		return
	}
	response.ReturnDeFaultFail(c)
}
