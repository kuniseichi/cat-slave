package passage

import (
	"cat-slave/model"
	"time"
)

type Passage struct {
	ID       uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Content  string     `json:"content" gorm:"type:longtext;unique_index" `
	TypeId   string     `json:"typeId" gorm:"type:int;unique_index" `
	PageView int        `json:"pageView" gorm:"type:int;unique_index" `
	CreateAt time.Time  `json:"createAt" gorm:"column:create_at"`
	UpdateAt time.Time  `json:"updateAt" gorm:"column:update_at"`
	DeleteAt *time.Time `json:"deleteAt" gorm:"column:delete_at"`
}

type PassageListDto struct {
	Content  string     `json:"content" gorm:"type:longtext;unique_index" `
	PageView int        `json:"pageView" gorm:"type:int;unique_index" `
	CreateAt time.Time  `json:"createAt" gorm:"column:create_at"`
	UpdateAt time.Time  `json:"updateAt" gorm:"column:update_at"`
	DeleteAt *time.Time `json:"deleteAt" gorm:"column:delete_at"`
}

func (p *Passage) TableName() string {
	return "passage"
}

func Get(id int) (*Passage, error) {
	p := &Passage{}
	db := model.DB.Mysql.Where("id = ?", id).First(&p)
	return p, db.Error
}

func List() ([]*Passage, error) {
	passages := make([]*Passage, 0)
	db := model.DB.Mysql.Raw("select * from passage left join ").Scan(&passages)
	return passages, db.Error
}
