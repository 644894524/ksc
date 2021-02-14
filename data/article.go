package data

const tableName = "articles"

//定义模型
type Article struct {
	ID          int			`gorm:"column:id"`
	Name 		string		`gorm:"column:name"`
	Type		int			`gorm:"column:type;default:0"`
	Like		int			`gorm:"column:like;default:0"`
	Collection  int			`gorm:"column:collection;default:0"`
	CreateTime  int			`gorm:"column:create_time"`
	UpdateTime	int         `grom:"column:update_time"`
	Status		int			`gorm:"column:status;default:0"`
}

//表名
func (Article) TableName() string{
	return tableName
}