package passageDao

import "time"

type Passage struct {
	ID       uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Content  string     `json:"content" gorm:"type:longtext;unique_index" `
	TypeId   string     `json:"typeId" gorm:"type:int;unique_index" `
	PageView int        `json:"pageView" gorm:"type:int;unique_index" `
	CreateAt time.Time  `json:"createAt" gorm:"column:create_at"`
	UpdateAt time.Time  `json:"updateAt" gorm:"column:update_at"`
	DeleteAt *time.Time `json:"deleteAt" gorm:"column:delete_at"`
}

func (p *Passage) TableName() string {
	return "passage"
}
