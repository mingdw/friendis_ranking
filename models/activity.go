package models

import (
	"fmt"
	"friends_ranking/config/dbconn"
	"friends_ranking/config/globalConst"
	"friends_ranking/config/variable"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Activity struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Code      string `json:"code" gorm:"column:code"`
	Title     string `json:"title" gorm:"column:title"`
	Desc      string `json:"desc"  gorm:"column:desc"`
	StartTime string `json:"startTime" gorm:"column:startTime"`
	EndTime   string `json:"endTime" gorm:"column:endTime" `
	Status    int    `json:"status"  gorm:"column:status" `
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

// 查询数据之前统计条数
func (activity *Activity) counts(code, startTime, endTime string, status int) (counts int64) {
	sql := "SELECT  count(*) as counts  FROM  sys_activity  WHERE isDelete=0 "
	if code != "" {
		sql += " and (code like '%" + code + "%' or title like '%" + code + "')"
	}

	if status != 0 {
		sql += " and status =" + strconv.Itoa(status)
	}

	if startTime != "" && endTime != "" {
		sql += " and startTime >=" + startTime + " and  startTime <=" + endTime
	}
	if res := activity.Raw(sql).First(&counts); res.Error != nil {
		variable.ZapLog.Error("Activity - counts 查询数据条数出错", zap.Error(res.Error))
	}
	return counts
}

// 查询（根据关键词模糊查询）
func (activity *Activity) Show(code, startTime, endTime string, status, pageSize, limit int) (counts int64, temp []Activity) {
	if counts = activity.counts(code, startTime, endTime, status); counts > 0 {
		sql := "SELECT  *  FROM  sys_activity  WHERE isDelete=0 "
		if code != "" {
			sql += " and (code like '%" + code + "%' or title like '%" + code + "')"
		}

		if status != 0 {
			sql += " and status =" + strconv.Itoa(status)
		}

		if startTime != "" && endTime != "" {
			sql += " and startTime >=" + startTime + " and  startTime <=" + endTime
		}
		sql += " LIMIT ?,?"
		limitStart := (pageSize - 1) * limit
		limitItems := limitStart + limit
		if res := activity.Raw(sql, limitStart, limitItems).Find(&temp); res.RowsAffected > 0 {
			return counts, temp
		}
	}

	return 0, nil
}

func (activity *Activity) Add(ac *Activity) bool {
	sql := "insert into sys_activity(code,title,desc,startTime,endTime,status,isDelete,creator,editor) values(?,?,?,?,?,?,?,?)"
	if res := activity.Exec(sql, ac.Code, ac.Title, ac.Desc, ac.StartTime, ac.EndTime, 0, 0, ac.Creator, ac.Editor); res.RowsAffected >= 0 {
		return true
	}
	return false
}

func (activity *Activity) Update(ac *Activity) bool {
	sql := "select * from sys_activity where id=? and isDelete=0 "
	rersult := activity.Raw(sql, ac.Id).First(activity)
	if rersult.Error == nil {
		if ac.Title != "" {
			activity.Title = ac.Title
		}
		if ac.Desc != "" {
			activity.Desc = ac.Desc
		}
		if ac.StartTime != "" {
			activity.StartTime = ac.StartTime
		}
		if ac.EndTime != "" {
			activity.EndTime = ac.EndTime
		}

		err := activity.DB.Model(&Activity{}).Where("id = ?", ac.Id).Updates(activity).Error
		if err == nil {
			return true
		}

	}
	return false
}

func (activity *Activity) Delete(ids []int) bool {
	sql := "delete from  sys_activity where id in(?) "
	strs := make([]string, len(ids))
	for k, v := range ids {
		strs[k] = fmt.Sprintf("%d", v)
	}
	idsStr := strings.Join(strs, ",")
	fmt.Println("ids: ", idsStr)
	if res := activity.Raw(sql, idsStr); res.RowsAffected >= 0 {
		return true
	}
	return false

}
