package models

import "time"

// 评论的数据结构

type Community struct {
	ID           int64     `json:"community_id" db:"community_id"`
	Name         string    `json:"community_name" db:"community_name"`
	Introduction string    `json:"introduction" db:"introduction"`
	CreatedAt    time.Time `json:"create_time" db:"create_time"`
	UpdatedAt    time.Time `json:"update_time" db:"update_time"`
}
