package controller

import (
	"bluebell/logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 与社区相关

func CommunityHandler(c *gin.Context) {
	// 期望是查询到所有的社区(community_id,community_name) 以列表的形式返回
	date, err := logic.GetCommunityList()
	if err != nil {
		fmt.Println(err)
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		// 不要对外暴露一些详细的信息
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}

// CommunityDetailHandler 根据ID查询社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id") // 获取路径参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据id获取社区信息
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)

}
