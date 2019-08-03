package remind

import (
	"cat-slave/model"
	"time"
)

type Remind struct {
	ID       uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserId 	 uint64		`json:"userId" gorm:"column:user_id"`
	Content  string     `json:"content"`
	Interval   uint64     `json:"interval"`
	BeginAt time.Time        `json:"beginAt"`
	IsDone uint64  `json:"isDone" gorm:"column:is_done"`
	CreateAt time.Time  `json:"createAt" gorm:"column:create_at"`
}

func (r *Remind) TableName() string {
	return "remind"
}

func List(userId int) ([]*Remind, error) {
	reminds := make([]*Remind, 0)
	db := model.DB.Mysql.Raw("select * from remind where user_id = ?", userId).Scan(&reminds)
	return reminds, db.Error
}