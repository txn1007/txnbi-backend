package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"txnbi-backend/api"
	"txnbi-backend/errs"
	"txnbi-backend/internal/biz"
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
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	// 校验参数
	accountLen, passwordLen := len(req.Account), len(req.Password)
	if accountLen < 6 || accountLen > 16 {
		log.Info().Err(errs.ErrUsernameLengthOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: errs.ErrUsernameLengthOutOfRange.Error()})
		return
	}
	if passwordLen < 8 || passwordLen > 24 {
		log.Info().Err(errs.ErrPasswordLengthOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: errs.ErrPasswordLengthOutOfRange.Error()})
		return
	}

	token, err := biz.UserLogin(req.Account, req.Password)
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: errs.ErrUserLoginFailed.Error()})
		return
	}
	log.Info().Interface("req", req).Msg("登陆成功")
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
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: errs.ErrUserRegistrationFailed.Error()})
		return
	}

	// 校验参数
	accountLen, passwordLen, inviteCodeLen := len(req.Account), len(req.Password), len(req.InviteCode)
	if accountLen < 6 || accountLen > 16 {
		log.Info().Err(errs.ErrUsernameLengthOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: errs.ErrUsernameLengthOutOfRange.Error()})
		return
	}
	if passwordLen < 8 || passwordLen > 24 {
		log.Info().Err(errs.ErrPasswordLengthOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: errs.ErrPasswordLengthOutOfRange.Error()})
		return
	}
	if inviteCodeLen < 2 || inviteCodeLen > 16 {
		log.Info().Err(errs.ErrInviteCodeLengthOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: errs.ErrInviteCodeLengthOutOfRange.Error()})
		return
	}

	err := biz.UserRegister(ctx, req.Account, req.Password, req.InviteCode)
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: errs.ErrUserRegistrationFailed.Error()})
		return
	}

	log.Info().Interface("req", req).Msg("登陆成功")
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
//	@Router			/user/auth/currentUserDetail [get]
func CurrentUserDetail(ctx *gin.Context) {
	var req api.CurrentUserDetailReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: errs.ErrGetCurrentUserDetailsFailed.Error()})
		return
	}

	userID := ctx.GetInt64("userID")

	user, err := biz.CurrentUserDetail(userID)
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UserRegisterResp{StatusCode: 1, Message: errs.ErrGetCurrentUserDetailsFailed.Error()})
		return
	}
	// 成功
	log.Info().Interface("req", req).Msg("用户获取自身信息成功")
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
//	@Router			/user/auth/loginOut [post]
func UserLoginOut(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	err := biz.UserLoginOut(userID)
	if err != nil {
		log.Info().Err(err).Int64("req", userID).Msg("")
		ctx.JSON(http.StatusOK, api.UserLoginOutResp{StatusCode: 1, Message: errs.ErrLogoutFailed.Error()})
		return
	}
	log.Info().Int64("userID", userID).Msg("用户登陆成功")
	ctx.JSON(http.StatusOK, api.UserLoginOutResp{StatusCode: 0, Message: "退出登陆成功！"})
	return
}
