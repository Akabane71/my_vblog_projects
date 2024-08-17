package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

func InsertPost(p *models.Post) (err error) {
	sqlStr := `insert into post 
	(post_id,title,content,author_id,community_id)
	values(?,?,?,?,?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostByID 根据id查询单个帖子数据
func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询 组合接口我们接口思想想要的数据
	post := new(models.Post) // 手动分配，如果不分配会出现栈消亡的情况
	sqlStr := `select post_id,title,content,author_id,community_id from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	if err != nil {
		zap.L().Error("GetPostByID", zap.Error(err))
		return
	}
	return

}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (data []*models.Post, err error) {
	data = make([]*models.Post, 2)
	sqlStr := `select 
    post_id,title,content,author_id,community_id,create_time
    from post
    ORDER BY create_time DESC 
	limit ?,?`

	err = db.Select(&data, sqlStr, page, size)
	return data, err
}

// GetPostListByIDs 根据给定的id列表查询帖子列表
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time 
	from post 
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)` // FIND_IN_SET为MySQL自带的
	// 使用sqlx 拼接处mysql语句
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	// 拼接
	query = db.Rebind(query)
	db.Select(&postList, query, args)
	return postList, err
}
