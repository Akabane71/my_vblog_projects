package models

import "time"

// 内存对齐，相近的放一块

type Post struct {
	ID          int64 `json:"post_id" dbL:"post_id" binding:"required"`
	AuthorID    int64 `json:"author_id" dbL:"author_id" binding:"required"`
	CommunityID int64 `json:"community_id" dbL:"community_id" binding:"required"`
	Status      int32 `json:"status" dbL:"status"`
	// 内存对齐的需要; 尽量将类型相近的放一起
	Title     string    `json:"title" dbL:"title" binding:"required"`
	Content   string    `json:"content" dbL:"content" binding:"required"`
	CreatedAt time.Time `json:"created_time" dbL:"created_time"`
	UpdatedAt time.Time `json:"updated_time" dbL:"updated_time"`
}

// ApiPostDetail 帖子详情结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	*Post
	*Community
}
