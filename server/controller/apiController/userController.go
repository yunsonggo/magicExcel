package apiController

import (
	"2021/magicExcel/server/common"
	"2021/magicExcel/server/dao"
	"2021/magicExcel/server/model"
	"2021/magicExcel/server/param"
	"2021/magicExcel/server/until"
	"2021/magicExcel/server/until/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var ud = dao.NewUserDao()

// 用户注册
func RegisterController(ctx *gin.Context) {
	var registerParam param.RegisterParam
	err := ctx.ShouldBind(&registerParam)
	if err != nil {
		zap.L().Error("获取注册参数失败", zap.Error(err))
		common.Failed(ctx, err, common.ParamFailed, "获取注册参数失败")
		return
	}
	// 验证验证码
	var code common.VerifyCaptchaBody
	code.Id = registerParam.CaptchaId
	code.VerifyValue = registerParam.Captcha
	if !common.VerifyCaptchaCode(code) {
		common.Failed(ctx, err, common.CaptchaFailed, "验证码错误")
		return
	}
	if registerParam.CheckPass != registerParam.Password {
		err := errors.New("两次密码不一致")
		common.Failed(ctx, err, common.FailedCode, "两次密码不一致")
		return
	}
	//查询是否注册
	res, resErr := ud.QueryUserByPassword(registerParam.Password)
	if resErr != nil {
		zap.L().Error("密码查询用户失败", zap.Error(err))
		common.Failed(ctx, err, common.QueryDBFailed, "密码查询用户失败")
		return
	}
	if res.ID > 0 || res.Name == registerParam.Name {
		common.Failed(ctx, err, common.FailedCode, "用户已经注册")
		return
	}
	var user = new(model.UserModel)
	user.Name = registerParam.Name
	password := until.EncodeSha256(registerParam.Password)
	user.Password = password
	id, insertErr := ud.InsertUserByPassword(user)
	if insertErr != nil {
		zap.L().Error("新建用户失败", zap.Error(err))
		common.Failed(ctx, err, common.InsertDBFailed, "新建用户失败")
	}
	fmt.Printf("id:%d\n", id)
	common.Success(ctx, "注册成功", "ok")
	return
}

// 用户登录
func LoginController(ctx *gin.Context) {
	var loginParam param.LoginParam
	err := ctx.ShouldBind(&loginParam)
	if err != nil {
		zap.L().Error("获取登录参数失败", zap.Error(err))
		common.Failed(ctx, err, common.ParamFailed, "获取登录参数失败")
		return
	}
	// 验证验证码
	var code common.VerifyCaptchaBody
	code.Id = loginParam.CaptchaId
	code.VerifyValue = loginParam.Captcha
	if !common.VerifyCaptchaCode(code) {
		err = errors.New("验证码校验错误")
		common.Failed(ctx, err, common.CaptchaFailed, "验证码错误")
		return
	}
	// 验证用户
	res, resErr := ud.QueryUserByName(loginParam.Name)
	if resErr != nil {
		zap.L().Error("查询用户名失败", zap.Error(err))
		common.Failed(ctx, err, common.QueryDBFailed, "查询用户名失败")
		return
	}
	if res.ID == 0 {
		err = errors.New("用户名或密码错误")
		common.Failed(ctx, err, common.FailedCode, "用户名或密码错误")
		return
	}
	password := until.EncodeSha256(loginParam.Password)
	if password != res.Password {
		err = errors.New("用户名或密码错误")
		common.Failed(ctx, err, common.FailedCode, "用户名或密码错误")
		return
	}
	token, tokenErr := jwt.ReleaseToken(res.ID)
	if tokenErr != nil {
		zap.L().Error("生成token失败", zap.Error(tokenErr))
		common.Failed(ctx, tokenErr, common.TokenFailed, "生成token失败")
		return
	}
	common.Success(ctx, token, "ok")
	return
}

// 修改密码
func ResetUserPass(ctx *gin.Context) {
	ctxId, ok := ctx.Get("ctxUserId")
	if !ok {
		err := errors.New("token无效,请重新登录")
		common.Failed(ctx, err, common.TokenFailed, "token无效,请重新登录")
		return
	}
	id := ctxId.(int64)
	fmt.Printf("id type: %T,%v\n", id, id)
	var resetParam param.ResetParam
	err := ctx.ShouldBind(&resetParam)
	if err != nil {
		zap.L().Error("修改密码参数错误", zap.Error(err))
		common.Failed(ctx, err, common.ParamFailed, "修改密码参数错误")
		return
	}
	if resetParam.Pass != resetParam.CheckPass {
		err = errors.New("密码不一致")
		common.Failed(ctx, err, common.ParamFailed, "密码不一致")
		return
	}
	password := until.EncodeSha256(resetParam.Pass)
	oldPassword := until.EncodeSha256(resetParam.OldPass)
	userInfo, userErr := ud.QueryUserByPassword(oldPassword)
	if userErr != nil || userInfo.ID == 0 {
		zap.L().Error("用户不存在或用户密码错误", zap.Error(userErr))
		common.Failed(ctx, userErr, common.QueryDBFailed, "用户不存在或用户密码错误")
		return
	}
	err = ud.EditUserPass(id, password)
	if err != nil {
		zap.L().Error("更新用户密码失败", zap.Error(err))
		common.Failed(ctx, err, common.InsertDBFailed, "更新用户密码失败")
		return
	}
	common.Success(ctx, "", "修改密码成功")
	return
}
