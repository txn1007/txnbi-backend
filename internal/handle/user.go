package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"txnbi-backend/api"
	"txnbi-backend/internal/biz"
	"txnbi-backend/pkg/tlog"
)

// UserLogin godoc
//
//	@Summary		用户登陆接口
//	@Description	登陆界面中的用户登陆接口
//	@Tags			user
//	@Produce		json
//	@Param			LoginInfo	body		api.UserLoginReq	true	"登陆信息"
//	@Success		200			{object}	api.UserLoginResp
//	@Router			/user/login [post]
func UserLogin(ctx *gin.Context) {
	var req api.UserLoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		tlog.L.Debug().Msgf("用户登陆失败，原因: %s,原始数据: 输入的账号: %s 输入的密码: %s", err.Error(), req.Account, req.Password)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: "参数不合法！"})
		return
	}

	// 校验参数
	accountLen, passwordLen := len(req.Account), len(req.Password)
	if accountLen < 6 || accountLen > 16 {
		tlog.L.Debug().Msgf("用户登陆失败，原因: %s,输入的账号: %s 输入的密码: %s", "输入的账号长度不合法", req.Account, req.Password)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: "用户名长度超出要求范围，长度应在6 ~ 16位"})
		return
	}
	if passwordLen < 8 || passwordLen > 24 {
		tlog.L.Debug().Msgf("用户登陆失败，: %s,输入的账号: %s 输入的密码: %s", "输入的密码长度不合法", req.Account, req.Password)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: "密码长度超出要求范围，长度应在8 ~ 24位"})
		return
	}

	token, err := biz.UserLogin(req.Account, req.Password)
	if err != nil {
		tlog.L.Debug().Msgf("用户登陆失败，: %s,输入的账号: %s 输入的密码: %s", err.Error(), req.Account, req.Password)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: err.Error()})
		return
	}
	tlog.L.Debug().Msgf("用户登陆成功，输入的账号: %s 输入的密码: %s", req.Account, req.Password)
	ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 0, Message: "登陆成功", Token: token})
}

// UserRegister godoc
//
//	@Summary		用户注册接口
//	@Description	登陆界面中的用户注册接口
//	@Tags			user
//	@Produce		json
//	@Param			RegisterInfo	body		api.UserRegisterReq	true	"注册信息"
//	@Success		200				{object}	api.UserRegisterResp
//	@Router			/user/register [post]
func UserRegister(ctx *gin.Context) {
	var req api.UserRegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		tlog.L.Debug().Msgf("用户注册失败，原因: %s,原始数据: 输入的账号: %s 输入的密码: %s, 输入的邀请码: %s", err.Error(), req.Account, req.Password, req.InviteCode)
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: err.Error()})
		return
	}

	// 校验参数
	accountLen, passwordLen, inviteCodeLen := len(req.Account), len(req.Password), len(req.InviteCode)
	if accountLen < 6 || accountLen > 16 {
		tlog.L.Debug().Msgf("用户注册失败，原因: %s,输入的账号: %s 输入的密码: %s, 输入的邀请码: %s", "输入的账号长度不合法", req.Account, req.Password, req.InviteCode)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: "用户名长度超出要求范围，长度应在6 ~ 16位"})
		return
	}
	if passwordLen < 8 || passwordLen > 24 {
		tlog.L.Debug().Msgf("用户注册失败，: %s,输入的账号: %s 输入的密码: %s, 输入的邀请码: %s", "输入的密码长度不合法", req.Account, req.Password, req.InviteCode)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: "密码长度超出要求范围，长度应在8 ~ 24位"})
		return
	}
	if inviteCodeLen < 2 || inviteCodeLen > 16 {
		tlog.L.Debug().Msgf("用户注册失败，: %s,输入的账号: %s 输入的密码: %s, 输入的邀请码: %s", "输入的邀请码长度不合法", req.Account, req.Password, req.InviteCode)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: "邀请码长度超出要求范围，长度应在 2 ~ 24位"})
		return
	}

	err := biz.UserRegister(ctx, req.Account, req.Password, req.InviteCode)
	if err != nil {
		tlog.L.Debug().Msgf("用户注册失败，原因: %s,原始数据: 输入的账号: %s 输入的密码: %s,输入的邀请码: %s", err.Error(), req.Account, req.Password, req.InviteCode)
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: err.Error()})
		return
	}
	tlog.L.Debug().Msgf("用户注册成功，原始数据: 输入的账号: %s 输入的密码: %s ,输入的邀请码: %s", req.Account, req.Password, req.InviteCode)
	ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 0, Message: "注册成功！"})
}

// CurrentUserDetail godoc
//
//	@Summary		用户自身详情接口
//	@Description	用户自身详情接口
//	@Tags			user
//	@Produce		json
//	@Param			Info	query		api.CurrentUserDetailReq	true "查询参数"
//	@Success		200		{object}	api.CurrentUserDetailResp
//	@Router			/user/CurrentUserDetail [get]
func CurrentUserDetail(ctx *gin.Context) {
	var req api.CurrentUserDetailReq
	if err := ctx.ShouldBind(&req); err != nil {
		tlog.L.Debug().Msgf("获取当前用户详细信息失败，原因：%s", err.Error())
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: "获取当前用户详细信息失败！"})
		return
	}
	userID := ctx.GetInt64("userID")
	// 校验参数
	if userID <= 0 {
		tlog.L.Debug().Msgf("获取当前用户详细信息失败，原因：%s,原始数据: userID = %d ", "userID不合法", userID)
		ctx.JSON(http.StatusOK, api.UserLoginOutResp{StatusCode: 1, Message: "获取当前用户详细信息失败！"})
		return
	}

	user, err := biz.CurrentUserDetail(userID)
	if err != nil {
		tlog.L.Debug().Msgf("获取当前用户详细信息失败，原因：%s,原始数据: userID = %d ", err.Error(), userID)
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: "获取当前用户详细信息失败"})
		return
	}
	// 成功
	tlog.L.Debug().Msgf("获取当前用户详细信息成功，原始数据: userID = %d ", userID)
	ctx.JSON(http.StatusOK, api.CurrentUserDetailResp{StatusCode: 0, Message: "获取用户本人信息成功！", UserInfoV0: api.UserInfoV0{
		ID: user.ID, UserAccount: user.UserAccount, UserName: user.UserName, UserAvatar: user.UserAvatar,
		UserRole: user.UserRole, CreateTime: user.CreateTime, UpdateTime: user.UpdateTime,
	},
	})
	return
}

// UserLoginOut godoc
//
//	@Summary		用户退出登陆接口
//	@Description	用户自身详情接口
//	@Tags			user
//	@Produce		json
//	@Param			Info	query		api.UserLoginOutReq	true "参数"
//	@Success		200		{object}	api.UserLoginOutResp
//	@Router			/user/loginOut [post]
func UserLoginOut(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	// 校验参数
	if userID <= 0 {
		tlog.L.Debug().Msgf("用户登出失败，原因：%s,原始数据: userID = %d ", "userID不合法", userID)
		ctx.JSON(http.StatusOK, api.UserLoginOutResp{StatusCode: 1, Message: "退出登陆失败！"})
		return
	}

	err := biz.UserLoginOut(userID)
	if err != nil {
		tlog.L.Debug().Msgf("用户登出失败，原因：%s,原始数据: userID = %d ", err.Error(), userID)
		ctx.JSON(http.StatusOK, api.UserLoginOutResp{StatusCode: 1, Message: "退出登陆失败！"})
		return
	}
	tlog.L.Debug().Msgf("用户登出成功，原始数据: userID = %d ", userID)
	ctx.JSON(http.StatusOK, api.UserLoginOutResp{StatusCode: 0, Message: "退出登陆成功！"})
	return
}
