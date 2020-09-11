package model

import (
	"encoding/json"
	"time"
)

type User struct {
	BaseModel
	Username string    `gorm:"size:50;default:'';comment:'登录名';not null;index" form:"username" binding:"required"`
	Password string    `gorm:"type:varchar(100);default:'';comment:'密码';not null" form:"pass"`
	Status   int       `gorm:"type:tinyint;default:1"`
	CreateAt time.Time `gorm:"column:ctime;comment:'创建时间';not null" json:'-'`
}

func (user *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(user)
}

func (user *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, user)
}

type Post struct {
	PostId  int    `gorm:"primary_key"`
	Title   string `gorm:"type:varchar(255);index;not null;default:'';comment:'标题'"`
	Content string `gorm:"type:"text";not null;default('')`
	UserId  int    `gorm:"column:user_id;type:int;size:11;default(0)"`
}
