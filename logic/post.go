package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成POST 的 id
	id := int64(snowflake.GenID())
	p.ID = id

	// 2. 保存到数据库
	err = mysql.InsertPost(p)
	if err != nil {
		return err
	}

	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}

// GetPostByID 根据id查询帖子详情
func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	data = new(models.ApiPostDetail)

	// 查询并组合我们接口想要的数据
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}

	// 根据社区id查询社区的详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}

	// 拼装数据
	data.AuthorName = user.Username
	data.Community = community

	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	// 查找所有的帖子

	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return data, err
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	// 再查找相关的信息
	for _, post := range posts {
		// 根据作者id查阅作者信息
		user, err := mysql.GetUserByID(post.ID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", post.ID), zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", post.ID), zap.Error(err))
			continue
		}

		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			Post:       post,
			Community:  community,
		}

		data = append(data, postDetail)
	}

	//

	return data, err
}

func GetPostList2(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	// 2. 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p.ParamPostList)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder() failed", zap.Error(err))
		return
	}

	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 ")
		return
	}

	// 3. 根据id去MySQL查询数据库帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	// 提前查询好每天的
	voteDate, err := redis.GetPostVoteDate(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for index, post := range posts {
		// 根据作者id查阅作者信息
		user, err := mysql.GetUserByID(post.ID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", post.ID), zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", post.ID), zap.Error(err))
			continue
		}

		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum:    voteDate[index], // 按照对应的切片寻找对应的数据
			Post:       post,
			Community:  community,
		}

		data = append(data, postDetail)
	}

	return
}

func GetCommunityPostList2(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	// 2. 去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder() failed", zap.Error(err))
		return
	}

	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 ")
		return
	}

	// 3. 根据id去MySQL查询数据库帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	// 提前查询好每天的
	voteDate, err := redis.GetPostVoteDate(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for index, post := range posts {
		// 根据作者id查阅作者信息
		user, err := mysql.GetUserByID(post.ID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", post.ID), zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed", zap.Int64("pid", post.ID), zap.Error(err))
			continue
		}

		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum:    voteDate[index], // 按照对应的切片寻找对应的数据
			Post:       post,
			Community:  community,
		}

		data = append(data, postDetail)
	}

	return
}

// GetPostListNew 将两个查询帖子逻辑合二为一的函数
func GetPostListNew(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	// 根据请求参数的不同，执行不同的额逻辑
	if p.CommunityID == 0 {
		data, err = GetPostList2(p)
	} else {
		data, err = GetCommunityPostList2(p)
	}

	if err != nil {
		zap.L().Error("GetPostListNew() failed", zap.Error(err))
		return nil, err
	}
	return data, err
}
