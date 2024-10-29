package models

import (
	"lottery_annual/config/dbconn"
	"lottery_annual/config/globalConst"
	"time"
)

type ActivityAndPrizeDetails struct {
	Id           int             `json:"id" gorm:"primaryKey"`
	Code         string          `json:"code" gorm:"column:code"`
	Title        string          `json:"title" gorm:"column:title"`
	Status       int             `json:"status"  gorm:"column:status" `
	PrizeDetails []*PrizeDetails `gorm:"foreignKey:AcId"`
	dbconn.BaseModel
}

type PrizeDetails struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	AcId       int       `json:"acId" gorm:"column:ac_id"`
	Name       string    `json:"name" gorm:"column:name"`
	ItemCode   string    `json:"itemCode"  gorm:"column:item_code"`
	PrizeNum   int       `json:"prizeNum" gorm:"column:prize_num"`
	IsRepeat   int       `json:"isRepeat" gorm:"column:is_repeat"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	Creator    string    `json:"creator" gorm:"column:creator"`
	EditTime   time.Time `json:"editTime" gorm:"column:editTime"`
	Editor     string    `json:"editor" gorm:"column:editor"`
	IsDelete   int       `json:"isDelete" gorm:"column:isDelete"`
	dbconn.BaseModel
}

func (p *PrizeDetails) TableName() string {
	return "sys_activity"
}

func CreatePrizeFactory(sqlType string) *PrizeDetails {
	now := time.Now()
	nowTime := now.Format(globalConst.DateFormat)
	creator := globalConst.SysAccount
	return &PrizeDetails{BaseModel: dbconn.BaseModel{
		DB:         dbconn.UseDbConn(sqlType),
		Id:         0,
		CreateTime: nowTime,
		EditTime:   nowTime,
		Creator:    creator,
		Editor:     creator,
		IsDelete:   0,
	}}
}

// 查询（根据关键词模糊查询）
func (prize *PrizeDetails) SelectList(code, prizeName string, status, pageSize, limit int) (counts int64, temp []ActivityAndPrizeDetails) {

	acAndPrizes := ActivityAndPrizeDetails{}
	// 计算总记录数并执行分页查询
	sql := prize.Model(&ActivityAndPrizeDetails{}).Where("isDelete = ?", 0)
	if code != "" {
		codeStr := "%" + code + "%"
		sql.Where("code like ? or title like?", codeStr, codeStr)
	}
	if status != 0 {
		sql.Where("status = ?", status)
	}
	sql.Count(&counts).Offset((pageSize - 1) * limit).Limit(limit).Order("editTime desc").Find(&temp)

	return
}
