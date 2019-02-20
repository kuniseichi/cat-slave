package userDao

import (
	"cat-slave/model"
	"fmt"
	"time"
)

type User struct {
	ID       uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	OpenID   string     `gorm:"type:varchar(100);unique_index"`
	CreateAt time.Time  `gorm:"column:create_at"`
	UpdateAt time.Time  `gorm:"column:update_at"`
	DeleteAt *time.Time `gorm:"column:delete_at"`
}

func (User) TableName() string {
	users := make([]*User, 0)
	model.DB.Mysql.Raw("select * from user").Scan(&users)
	// model.DB.mysql
	fmt.Print(users)

	// 新增 自动添加createAt和updateAt
	// u := &User{
	// 	OpenID: "232422232",
	// }
	// db.mysql.Create(u)
	// fmt.Print(u.ID)

	//删
	// db.mysql.Where("", "").Delete(&User{
	// 	ID: 3,
	// })

	//改 全改(没给的变空)
	// db.mysql.Save(&User{
	// 	ID:     2,
	// 	OpenID: "11111",
	// })
	// 改字段 批量修改
	// db.mysql.Model(&User{
	// 	ID: 2,
	// }).Where("open_id", "jinzhu").Update("", "")
	// struct 修改 "" 0 false 不做修改
	// db.mysql.Model(&User{
	// 	ID: 2,
	// }).Updates(&UserDto{
	// 	OpenID: "123232112323",
	// })

	return "user"
}

func Create() {

}

func GetUserList() []*User {
	users := make([]*User, 0)
	model.DB.Mysql.Raw("select * from user").Scan(&users)
	// model.DB.mysql
	fmt.Print(users)
	return users
}
