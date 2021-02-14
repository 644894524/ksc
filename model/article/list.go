package article

import (
	"fmt"
	"ksc/common"
	"ksc/data"
)

const LIMIT = 20

func List(page int){
	offset := page * LIMIT
	var articles data.Article
	db := common.GetDb()
	db.Where("status = ?", 0).Order("update_time DESC").Offset(offset).Limit(10).Find(&articles)
	fmt.Println(articles)
}