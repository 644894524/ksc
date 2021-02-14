package data

const tableName = "articles"

//定义模型
type Article struct {
	ID          int			`gorm:"column:id"`
	CreateTime  int			`gorm:"column:create_time"`
	UpdateTime	int         `grom:"column:update_time"`
	Name 		string		`gorm:"column:name"`
	Type		int			`gorm:"column:type;default:0"`
	Like		int			`gorm:"column:like;default:0"`
	Collection  int			`gorm:"column:collection;default:0"`
	Status		int			`gorm:"column:status;default:0"`
}

//表名
func (Article) TableName() string{
	return tableName
}