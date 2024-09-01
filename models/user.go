package models

import (
	"fmt"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Age        int    `json:"age"`
	Sex        int    `json:"sex"`
	Phone      string `json:"phone"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"CreateTime"`
	EditTime   int64  `json:"EditTime"`
	Creator    string `json:"Creator"`
	Editor     string `json:"Editor"`
}

func (User) TableName() string {
	return "sys_user"
}

func GetUserInfoById(id int) User {
	return initUser(id)
}

func initUser(id int) User {
	user := User{
		Id:         id,
		Username:   "张三",
		Password:   "pwdzhangsan",
		Age:        20,
		Sex:        0,
		Phone:      "18710127234",
		CreateTime: time.Now().Unix(),
		EditTime:   time.Now().Unix(),
		Creator:    "系统创建",
		Editor:     "系统更新",
	}
	fmt.Println("user info: ", user)
	return user
}
