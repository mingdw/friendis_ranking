package controller

import (
	prizeService "lottery_annual/service/prize/crud"

	"github.com/gin-gonic/gin"
)

type PrizeController struct {
}

func (index PrizeController) QueryList(c *gin.Context, code, prizeName string, status, limit, pageSize int) {
	prizeService.DetailList(code, prizeName, status, limit, pageSize, c)
}

func (index PrizeController) PrizeAddAndUpDate(c *gin.Context, id, acId, itemCode, isRepeat, prizeNum int, name string) {
	if id == -1 {
		prizeService.Add(c, acId, itemCode, prizeNum, isRepeat, name)
	} else {
		prizeService.Update(c, id, acId, itemCode, prizeNum, isRepeat, name)
	}
}

func (index PrizeController) Delete(c *gin.Context, ids []int) {
	prizeService.Delete(ids, c)
}

func (index PrizeController) SelectByAcId(c *gin.Context, acId int) {
	prizeService.SelectByAcId(acId, c)
}
