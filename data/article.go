package data

const tableName = "articles"

//定义模型
type Article struct {
	ID          int			`gorm:"column:id" json:"id"`
	Name 		string		`gorm:"column:name" json:"name"`
	Type		int			`gorm:"column:type;default:0" json:"type"`
	Like		int			`gorm:"column:like;default:0" json:"like"`
	Collection  int			`gorm:"column:collection;default:0" json:"collection"`
	CreateTime  int			`gorm:"column:create_time" json:"create_time"`
	UpdateTime	int         `grom:"column:update_time" json:"update_time"`
	Status		int			`gorm:"column:status;default:0" json:"status"`
	Content     string		`gorm:"column:content" json:"content"`
	Imgs		string		`gorm:"column:imgs" json:"imgs"`
}

//表名
func (Article) TableName() string{
	return tableName
}