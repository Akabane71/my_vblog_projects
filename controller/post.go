package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 3. 从c中取到当前发请求的用户ID值
	id := c.GetInt64(CtxUserIDKey)
	p.ID = int64(id)

	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	// 3. 返回响应
	ResponseSuccess(c, p)
}

// GetPostDetailHandler 从贴子中获取详解信息
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据pid取出帖子数据(查数据库)
	data, err := logic.GetPostByID(pid)
	if err != nil {
		// 日志你想记录就记录;但你一定要能定位到出问题的地方
		zap.L().Error("logic.GetPostByID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取数据列表
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)

	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	// 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 根据前端传来的参数，动态地去获取帖子列表
// 按创建时间排序，或者按照分数排序
// 1. 获取参数 query string 残水
// 2. 去redis查询id值
// 3. 根据id去数据库查询帖子仔细信息
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数: /api/v1/post2?page=1&size=10&order=time
	// 建议写成道默认的配置文件中的参数
	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page: 1,
			Size: 10,
			//Order: "time", // 不要出现这种硬编码的词
			Order: models.OrderTime,
		},
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	} // 使用gin框架利用反射获取
	// shouldBind()  根据请求的类型选择相应的方法去获取参数
	// 如果请求中携带的是json格式的数据，才能使用这个方法获取到

	data, err := logic.GetPostListNew(p) // 更新: 合二为一

	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// GetCommunityPostListHandler 根据社区去查询帖子列表
func GetCommunityPostListHandler(c *gin.Context) {
	// GET请求参数: /api/v1/post2?page=1&size=10&order=time
	// 建议写成道默认的配置文件中的参数
	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderTime,
		},
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	} // 使用gin框架利用反射获取
	// shouldBind()  根据请求的类型选择相应的方法去获取参数
	// 如果请求中携带的是json格式的数据，才能使用这个方法获取到

	data, err := logic.GetCommunityPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
