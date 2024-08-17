package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票
type VoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    int64 `json:"post_id,string"`
	Direction int   `json:"direction,string"` // 赞成: 1 或者反对票 -1

}

// PostVoteController 投票
func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(models.ParamsVoteData)
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errDate := removeTopStruct(errs.Translate(trans)) // 翻译并去除错误提示中的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParam, errDate)
		return
	}
	// 获取当前请求用户的id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedKLogin)
	}
	if err = logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}
