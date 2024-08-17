package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
	"go.uber.org/zap"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamsSignUp) (err error) {
	// 1. 判断用户存不存在
	var exist bool
	if exist, err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询失败
		return errors.New("查询失败")
	}

	if exist {
		// 数据已经存在
		return errors.New("数据已经存在")
	}

	// 2. 生成ID
	userID := snowflake.GenID()

	// 构造一个User实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 4. 保存进数据库
	if err = mysql.InsertUser(&user); err != nil {
		zap.L().Error("insert user fail", zap.Error(err))
		return err
	}

	return
}

func Login(p *models.ParamsLogin) (err error) {
	// 1. 判读用户是否存在
	var exist bool
	if exist, err = mysql.CheckUserExist(p.Username); err != nil {
		zap.L().Error("user login error", zap.Error(err))
		return
	}

	// 2. 判断用户密码是否正确
	var user *models.User
	if exist {
		user, err = mysql.QueryUserByUserName(p.Username)
		if err != nil {
			zap.L().Error("user login error", zap.Error(err))
			return
		}
		// 密码相同 且正确
		if mysql.EncryptPassword(p.Password) == user.Password {
			return
		} else {
			return errors.New("password error")
		}
	}

	return
}

func UserLogin(p *models.ParamsLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针,就能拿到user.UserID
	if err = mysql.Login(user); err != nil {
		return token, err
	}
	// 生成JWT
	token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
