package userDao

import (
	"cat-slave/model"
	"time"
)

type User struct {
	ID       uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	OpenID   string     `json:"openId" gorm:"type:varchar(100);unique_index" `
	CreateAt time.Time  `json:"createAt" gorm:"column:create_at"`
	UpdateAt time.Time  `json:"updateAt" gorm:"column:update_at"`
	DeleteAt *time.Time `json:"deleteAt" gorm:"column:delete_at"`
}

func (User) TableName() string {
	return "user"
}

func Create() {

}

func GetUserList() []*User {
	users := make([]*User, 0)
	model.DB.Mysql.Raw("select * from user").Scan(&users)
	return users
}
