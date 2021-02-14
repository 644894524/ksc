package article

import (
	"ksc/common"
	"ksc/data"
)

const LIMIT = 20

func List(page int) []data.Article {
	offset := page * LIMIT
	var articles []data.Article
	db := common.GetDb()
	db.Where("status = ?", 0).Order("update_time DESC").Offset(offset).Limit(LIMIT).Find(&articles)
	return articles
}