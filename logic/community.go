package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 查找到所有的Community并返回
func GetCommunityList() (data []*models.Community, err error) {
	// 查找数据库
	data, err = mysql.GetCommunityList()
	return data, err
}

func GetCommunityDetail(id int64) (*models.Community, error) {
	return mysql.GetCommunityDetailByID(id)
}
