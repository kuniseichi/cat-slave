package passage

import (
	"cat-slave/model"
	"strings"
	"time"
)

type Passage struct {
	ID       uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Title 	 string		`json:"title"`
	Content  string     `json:"content" gorm:"type:longtext;unique_index" `
	TypeId   string     `json:"typeId" gorm:"type:int;unique_index" `
	PageView int        `json:"pageView" gorm:"type:int;unique_index" `
	CreateAt time.Time  `json:"createAt" gorm:"column:create_at"`
	UpdateAt time.Time  `json:"updateAt" gorm:"column:update_at"`
	DeleteAt *time.Time `json:"deleteAt" gorm:"column:delete_at"`
	Weight 	 int 		`json:"weight"`
	Url 	 string 	`json:"url" gorm:"column:url"`
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
	db := model.DB.Mysql.Raw("select * from passage").Scan(&passages)
	return passages, db.Error
}

func Index(keyword string) ([]*Passage, error) {
	passages := make([]*Passage, 0)
	db := model.DB.Mysql.Raw("select * from passage where title like ? or content like ?", "%" + keyword + "%", "%" + keyword + "%").Scan(&passages)

	for _, value := range passages {
		// 1. 清除\n
		value.Content = strings.Replace(value.Content, "\n", "", -1)
		// 2. 加入权重
		value.Weight = strings.Count(value.Content, keyword)
		if strings.Contains(value.Title, keyword) {
			value.Weight += 100
		}
	}

	// 根据权重排序
	for i := 0; i < len(passages); i++ {
		for j := i; j < len(passages); j++ {
			if passages[j].Weight > passages[i].Weight {
				p := passages[j]
				passages[j] = passages[i]
				passages[i] = p
			}
		}
	}

	// 截取关键部分
	for _, value := range passages {
		begin := UnicodeIndex(value.Content, keyword)
		value.Content = SubString(value.Content, begin - 20, 40)
	}

	return passages, db.Error
}

func UnicodeIndex(str,substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str,substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func SubString(str string,begin,length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

