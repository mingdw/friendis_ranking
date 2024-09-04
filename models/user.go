package models

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
	Db         dbconn
}

func (u *User) TableName() string {
	return "sys_user"
}

func insertUser()
