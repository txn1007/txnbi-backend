package handle

import (
	"net/http"
	"txnbi-backend/api"
	"txnbi-backend/errs"
	biz "txnbi-backend/internal/biz/admin"
	"txnbi-backend/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// AdminListUsers godoc
//
//	@Summary		管理员获取用户列表
//	@Description	管理员获取用户列表
//	@Tags			admin
//	@Produce		json
//	@Param			page		query		int		true	"页码"
//	@Param			pageSize	query		int		true	"每页数量"
//	@Param			keyword		query		string	false	"搜索关键词"
//	@Success		200			{object}	api.AdminUserListResp
//	@Router			/admin/user/list [get]
func AdminListUsers(ctx *gin.Context) {
	var req api.AdminUserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminUserListResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	var users []*model.User
	var total int64
	var err error

	// 如果有关键词，则搜索
	if req.Keyword != "" {
		users, total, err = biz.SearchUsers(ctx, req.Keyword, req.Page, req.PageSize)
	} else {
		users, total, err = biz.ListUsers(ctx, req.Page, req.PageSize)
	}

	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("获取用户列表失败")
		ctx.JSON(http.StatusOK, api.AdminUserListResp{StatusCode: 1, Message: "获取用户列表失败"})
		return
	}

	// 转换为响应格式
	userInfos := make([]api.UserInfoV1, 0, len(users))
	for _, user := range users {
		userInfos = append(userInfos, api.UserInfoV1{
			ID:          user.ID,
			UserAccount: user.UserAccount,
			UserName:    user.UserName,
			UserAvatar:  user.UserAvatar,
			UserRole:    user.UserRole,
			UserStatus:  user.UserStatus,
			CreateTime:  user.CreateTime,
			UpdateTime:  user.UpdateTime,
		})
	}

	ctx.JSON(http.StatusOK, api.AdminUserListResp{
		StatusCode: 0,
		Message:    "获取用户列表成功",
		Total:      total,
		Users:      userInfos,
	})
}

// AdminGetUserDetail godoc
//
//	@Summary		管理员获取用户详情
//	@Description	管理员获取用户详情
//	@Tags			admin
//	@Produce		json
//	@Param			userId	query		int	true	"用户ID"
//	@Success		200		{object}	api.AdminUserDetailResp
//	@Router			/admin/user/detail [get]
func AdminGetUserDetail(ctx *gin.Context) {
	var req api.AdminUserDetailReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminUserDetailResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	user, err := biz.GetUserDetail(ctx, req.UserID)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("获取用户详情失败")
		ctx.JSON(http.StatusOK, api.AdminUserDetailResp{StatusCode: 1, Message: "获取用户详情失败"})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminUserDetailResp{
		StatusCode: 0,
		Message:    "获取用户详情成功",
		User: api.UserInfoV1{
			ID:          user.ID,
			UserAccount: user.UserAccount,
			UserName:    user.UserName,
			UserAvatar:  user.UserAvatar,
			UserRole:    user.UserRole,
			UserStatus:  user.UserStatus,
			CreateTime:  user.CreateTime,
			UpdateTime:  user.UpdateTime,
		},
	})
}

// AdminCreateUser godoc
//
//	@Summary		管理员创建用户
//	@Description	管理员创建用户
//	@Tags			admin
//	@Produce		json
//	@Param			user	body		api.AdminCreateUserReq	true	"用户信息"
//	@Success		200		{object}	api.AdminCreateUserResp
//	@Router			/admin/user/create [post]
func AdminCreateUser(ctx *gin.Context) {
	var req api.AdminCreateUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminCreateUserResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	userID, err := biz.CreateUser(ctx, req.UserAccount, req.UserPassword, req.UserRole)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("创建用户失败")
		ctx.JSON(http.StatusOK, api.AdminCreateUserResp{StatusCode: 1, Message: "创建用户失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminCreateUserResp{
		StatusCode: 0,
		Message:    "创建用户成功",
		UserID:     userID,
	})
}

// AdminUpdateUser godoc
//
//	@Summary		管理员更新用户
//	@Description	管理员更新用户
//	@Tags			admin
//	@Produce		json
//	@Param			user	body		api.AdminUpdateUserReq	true	"用户信息"
//	@Success		200		{object}	api.AdminUpdateUserResp
//	@Router			/admin/user/update [post]
func AdminUpdateUser(ctx *gin.Context) {
	var req api.AdminUpdateUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminUpdateUserResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	user := &model.User{
		ID:           req.UserID,
		UserAccount:  req.UserAccount,
		UserPassword: req.UserPassword,
		UserName:     req.UserName,
		UserAvatar:   req.UserAvatar,
		UserRole:     req.UserRole,
	}

	err := biz.UpdateUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("更新用户失败")
		ctx.JSON(http.StatusOK, api.AdminUpdateUserResp{StatusCode: 1, Message: "更新用户失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminUpdateUserResp{
		StatusCode: 0,
		Message:    "更新用户成功",
	})
}

// AdminDeleteUser godoc
//
//	@Summary		管理员删除用户
//	@Description	管理员删除用户
//	@Tags			admin
//	@Produce		json
//	@Param			user	body		api.AdminDeleteUserReq	true	"用户ID"
//	@Success		200		{object}	api.AdminDeleteUserResp
//	@Router			/admin/user/delete [post]
func AdminDeleteUser(ctx *gin.Context) {
	var req api.AdminDeleteUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminDeleteUserResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	err := biz.DeleteUser(ctx, req.UserID)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("删除用户失败")
		ctx.JSON(http.StatusOK, api.AdminDeleteUserResp{StatusCode: 1, Message: "删除用户失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminDeleteUserResp{
		StatusCode: 0,
		Message:    "删除用户成功",
	})
}

// AdminDisableUser godoc
//
//	@Summary		管理员禁用/启用用户
//	@Description	管理员禁用/启用用户
//	@Tags			admin
//	@Produce		json
//	@Param			user	body		api.AdminDisableUserReq	true	"用户ID和状态"
//	@Success		200		{object}	api.AdminDisableUserResp
//	@Router			/admin/user/disable [post]
func AdminDisableUser(ctx *gin.Context) {
	var req api.AdminDisableUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminDisableUserResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	err := biz.DisableUser(ctx, req.UserID, req.Status)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("禁用/启用用户失败")
		ctx.JSON(http.StatusOK, api.AdminDisableUserResp{StatusCode: 1, Message: "禁用/启用用户失败: " + err.Error()})
		return
	}

	statusText := "启用"
	if req.Status == 1 {
		statusText = "禁用"
	}

	ctx.JSON(http.StatusOK, api.AdminDisableUserResp{
		StatusCode: 0,
		Message:    statusText + "用户成功",
	})
}
