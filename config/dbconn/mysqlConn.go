package dbconn

import (
	"fmt"
	"friends_ranking/config/errorMsg"
	"friends_ranking/config/variable"
	"strings"

	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB   `gorm:"-" json:"-"`
	Id         int64  `gorm:"primaryKey" json:"id"`
	CreateTime string `json:"createTime" gorm:"column:createTime"`
	EditTime   string `json:"editTime" gorm:"column:editTime"`
	Creator    string `json:"creator" gorm:"column:creator"`
	Editor     string `json:"editor" gorm:"column:editor"`
	IsDelete   int64  `json:"isDelete" gorm:"column:isDelete"`
}

func UseDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	sqlType = strings.Trim(sqlType, " ")

	if sqlType == "" {
		sqlType = variable.YamlConfig.GetString("Gormv2.UseDbType")
	}
	switch strings.ToLower(sqlType) {
	case "mysql":
		if variable.GormDbMysql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(errorMsg.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbMysql
	}
	return db
}

func ReConnectInterval(sqlType string) *gorm.DB {
	var db *gorm.DB
	sqlType = strings.Trim(sqlType, " ")
	if sqlType == "" {
		sqlType = variable.YamlConfig.GetString("Gormv2.UseDbType")
	}
	switch strings.ToLower(sqlType) {
	case "mysql":
		if variable.GormDbMysql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(errorMsg.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbMysql
	default:
		variable.ZapLog.Error(errorMsg.ErrorsDbDriverNotExists + sqlType)
	}
	return db
}
