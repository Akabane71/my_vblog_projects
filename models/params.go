package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// 定义请求的参数结构体

// binding 为 gin框架默认的validator库 读取的字段名
// validate 为 validator 库的 字段名

// 参看validate文档去方便开发

type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required,email"`
}

type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamsVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID string `json:"post_id,string" binding:"required"`
	// oneof = -1 0 1
	Direction int8 `json:"direction,string" binding:"required,oneof=1 0 -1"` // 赞成: 1 / 反对票 -1 / 取消投票 0

}

// ParamPostList 获取帖子列表参数 query string 参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` // 可以为空
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}

type ParamCommunityPostList struct {
	*ParamPostList
}
