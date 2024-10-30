package models

import (
	"fmt"
	"lottery_annual/config/dbconn"
	"lottery_annual/config/globalConst"
	"time"
)

type ActivityAndPrizeDetails struct {
	Id           int             `json:"id" gorm:"primaryKey"`
	Code         string          `json:"code" gorm:"column:code"`
	Title        string          `json:"title" gorm:"column:title"`
	Status       int             `json:"status"  gorm:"column:status" `
	PrizeDetails []*PrizeDetails `json :"prizeDetails"   gorm:"foreignKey:AcId"`
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

func (p *ActivityAndPrizeDetails) TableName() string {
	return "sys_activity"
}

func (p *PrizeDetails) TableName() string {
	return "sys_prize"
}

func CreatePrizeFactory(sqlType string) *ActivityAndPrizeDetails {
	now := time.Now()
	nowTime := now.Format(globalConst.DateFormat)
	creator := globalConst.SysAccount
	return &ActivityAndPrizeDetails{BaseModel: dbconn.BaseModel{
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
func (prize *ActivityAndPrizeDetails) SelectList(code, prizeName string, status, pageSize, limit int) (counts int64, temp []ActivityAndPrizeDetails) {

	// 计算总记录数并执行分页查询
	fmt.Println("开启查询，code；", code, "; prizeName: ", prizeName, "; status: ", status, "; pageSize: ", pageSize, "; limit: ", limit)
	sql := prize.Where("isDelete = ?", 0)
	if code != "" {
		codeStr := "%" + code + "%"
		sql.Where("code like ? or title like?", codeStr, codeStr)
	}
	if status != 0 {
		sql.Where("status = ?", status)
	}
	sql.Model(&prize).Count(&counts).Offset((pageSize - 1) * limit).Limit(limit).Order("editTime desc").First(&prize).Preload("PrizeDetails").Find(&temp)
	return
}
