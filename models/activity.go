package models

import (
	"fmt"
	"lottery_annual/config/dbconn"
	"lottery_annual/config/globalConst"
	"time"
)

type Activity struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code" gorm:"column:code"`
	Title     string    `json:"title" gorm:"column:title"`
	Desc      string    `json:"desc"  gorm:"column:ac_desc"`
	StartTime time.Time `json:"startTime" gorm:"column:startTime"`
	EndTime   time.Time `json:"endTime" gorm:"column:endTime" `
	Status    int       `json:"status"  gorm:"column:status" `
	dbconn.BaseModel
}

func (u *Activity) TableName() string {
	return "sys_activity"
}

func CreateActivityFactory(sqlType string) *Activity {
	now := time.Now()
	nowTime := now.Format(globalConst.DateFormat)
	creator := globalConst.SysAccount
	return &Activity{BaseModel: dbconn.BaseModel{
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
func (activity *Activity) Show(code string, startTime, endTime time.Time, status, pageSize, limit int) (counts int64, temp []Activity) {
	// 计算总记录数并执行分页查询
	sql := activity.Model(&Activity{}).Where("isDelete = ?", 0)
	if code != "" {
		codeStr := "%" + code + "%"
		sql.Where("code like ? or title like?", codeStr, codeStr)
	}

	if status != 0 {
		sql.Where("status = ?", status)
	}

	if !startTime.IsZero() && !endTime.IsZero() {
		fmt.Println("时间范围灵芝")
		sql.Where(" startTime>= ? and startTime<=?", startTime, endTime)
	}
	sql.Count(&counts).Offset((pageSize - 1) * limit).Limit(limit).Order("editTime desc").Find(&temp)
	return
}

func (activity *Activity) Add(ac *Activity) bool {
	ac.Creator = activity.Creator
	ac.Editor = activity.Editor
	ac.CreateTime = activity.CreateTime
	ac.EditTime = activity.EditTime
	if err := activity.Create(&ac).Error; err != nil {
		return false
	}
	fmt.Println("新增 成功")
	return true
}

func (activity *Activity) Update(ac *Activity) bool {
	if err := activity.DB.Model(&Activity{}).Where("id = ?", ac.Id).Updates(ac); err.RowsAffected >= 0 {
		return true
	}
	return false
}

func (activity *Activity) Delete(ids []int) bool {
	if res := activity.Table("sys_activity").Where("id in(?)", ids).Delete(nil); res.RowsAffected >= 0 {
		return true
	}
	return false
}

func (activity *Activity) UpdateStatus(id, status int) bool {
	if res := activity.Table("sys_activity").Where("id = ?", id).Update("status", status); res.RowsAffected >= 0 {
		return true
	}
	return false
}
