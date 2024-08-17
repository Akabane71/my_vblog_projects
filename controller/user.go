package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验

	// 这一步是最基本的格式判断!
	var p *models.ParamsSignUp
	// 使用了带三方库后会帮助我们校验
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误,记录错误,返回结果
		zap.L().Error("SignUp with invalid param", zap.Any("err", err))
		// 使用自己定义的错误方法
		ResponseError(c, CodeInvalidParam)
		return
	}
	fmt.Println(p)

	// 低效率 ----> 这种重复的字段校验 请使用开源库来辅助校验
	// 手动对请求参数进行详细的业务规则校验 [防止脚本攻击;和用户禁用了JS]
	if len(p.Username) == 0 || len(p.RePassword) == 0 || len(p.Password) == 0 || p.Password != p.RePassword {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		ResponseError(c, CodeUserExist)
		return
	}

	// 3. 返回响应 [中间件 或者 自己写]
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数及验证参数
	p := new(models.ParamsLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Any("err", err))
		ResponseError(c, CodeInvalidParam)
	}

	// 2. 业务处理----> 检查参数是否正确
	//if err := logic.Login(p); err != nil {
	//	zap.L().Error("Login with invalid param", zap.Any("err", err))
	//	ResponseError(c, CodeUserNotExist)
	//}

	// 更好的方式
	token, err := logic.UserLogin(p)
	if err != nil {
		zap.L().Error("Login with invalid param", zap.Any("err", err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, token)
}

func Login(p *models.ParamsLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err = mysql.Login(user); err != nil {
		return nil, err
	}

	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		user.Token = token
	}
	return
}
