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

// AdminListLogs godoc
//
//	@Summary		管理员获取日志列表
//	@Description	管理员获取日志列表
//	@Tags			admin-log
//	@Produce		json
//	@Param			page		query		int		true	"页码"
//	@Param			pageSize	query		int		true	"每页数量"
//	@Param			keyword		query		string	false	"搜索关键词"
//	@Param			startTime	query		string	false	"开始时间"
//	@Param			endTime		query		string	false	"结束时间"
//	@Success		200			{object}	api.AdminLogListResp
//	@Router			/admin/log/list [get]
func AdminListLogs(ctx *gin.Context) {
	var req api.AdminLogListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminLogListResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	// 校验参数
	if req.PageSize < 1 || req.PageSize > 100 {
		log.Info().Err(errs.ErrInvalidPageSize).Interface("req", req).Msg("页面大小无效")
		ctx.JSON(http.StatusOK, api.AdminLogListResp{StatusCode: 1, Message: errs.ErrInvalidPageSize.Error()})
		return
	}

	if req.Page < 1 {
		log.Info().Err(errs.ErrInvalidPageParameters).Interface("req", req).Msg("页码无效")
		ctx.JSON(http.StatusOK, api.AdminLogListResp{StatusCode: 1, Message: errs.ErrInvalidPageParameters.Error()})
		return
	}

	logs, total, err := biz.AdminListLogs(ctx, req.Page, req.PageSize, req.Keyword, req.StartTime, req.EndTime)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("获取日志列表失败")
		ctx.JSON(http.StatusOK, api.AdminLogListResp{StatusCode: 1, Message: "获取日志列表失败"})
		return
	}

	// 转换为响应格式
	logInfos := make([]api.OperationLog, 0, len(logs))
	for _, l := range logs {
		logInfos = append(logInfos, api.OperationLog{
			ID:          l.ID,
			UserID:      l.UserID,
			UserName:    l.UserName,
			UserAccount: l.UserAccount,
			Operation:   l.Operation,
			Method:      l.Method,
			Path:        l.Path,
			IP:          l.IP,
			CreateTime:  l.CreateTime,
		})
	}

	ctx.JSON(http.StatusOK, api.AdminLogListResp{
		StatusCode: 0,
		Message:    "获取日志列表成功",
		Total:      total,
		Logs:       logInfos,
	})
}

// AdminGetLogDetail godoc
//
//	@Summary		管理员获取日志详情
//	@Description	管理员获取日志详情
//	@Tags			admin-log
//	@Produce		json
//	@Param			logId	query		int	true	"日志ID"
//	@Success		200		{object}	api.AdminLogDetailResp
//	@Router			/admin/log/detail [get]
func AdminGetLogDetail(ctx *gin.Context) {
	var req api.AdminLogDetailReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminLogDetailResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	logDetail, err := biz.AdminGetLogDetail(ctx, req.LogID)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("获取日志详情失败")
		ctx.JSON(http.StatusOK, api.AdminLogDetailResp{StatusCode: 1, Message: "获取日志详情失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminLogDetailResp{
		StatusCode: 0,
		Message:    "获取日志详情成功",
		Log: api.OperationLog{
			ID:          logDetail.ID,
			UserID:      logDetail.UserID,
			UserName:    logDetail.UserName,
			UserAccount: logDetail.UserAccount,
			Operation:   logDetail.Operation,
			Method:      logDetail.Method,
			Path:        logDetail.Path,
			IP:          logDetail.IP,
			CreateTime:  logDetail.CreateTime,
		},
	})
}

// AdminCreateLog godoc
//
//	@Summary		管理员创建日志
//	@Description	管理员创建日志
//	@Tags			admin-log
//	@Produce		json
//	@Param			log	body		api.AdminCreateLogReq	true	"日志信息"
//	@Success		200	{object}	api.AdminCreateLogResp
//	@Router			/admin/log/create [post]
func AdminCreateLog(ctx *gin.Context) {
	var req api.AdminCreateLogReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminCreateLogResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	logModel := &model.OperationLog{
		UserID:    req.UserID,
		Operation: req.Operation,
		Method:    req.Method,
		Path:      req.Path,
		IP:        req.IP,
	}

	logID, err := biz.AdminCreateLog(ctx, logModel)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("创建日志失败")
		ctx.JSON(http.StatusOK, api.AdminCreateLogResp{StatusCode: 1, Message: "创建日志失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminCreateLogResp{
		StatusCode: 0,
		Message:    "创建日志成功",
		LogID:      logID,
	})
}

// AdminUpdateLog godoc
//
//	@Summary		管理员更新日志
//	@Description	管理员更新日志
//	@Tags			admin-log
//	@Produce		json
//	@Param			log	body		api.AdminUpdateLogReq	true	"日志信息"
//	@Success		200	{object}	api.AdminUpdateLogResp
//	@Router			/admin/log/update [post]
func AdminUpdateLog(ctx *gin.Context) {
	var req api.AdminUpdateLogReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminUpdateLogResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	logModel := &model.OperationLog{
		ID:        req.LogID,
		Operation: req.Operation,
		Method:    req.Method,
		Path:      req.Path,
		IP:        req.IP,
	}

	err := biz.AdminUpdateLog(ctx, logModel)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("更新日志失败")
		ctx.JSON(http.StatusOK, api.AdminUpdateLogResp{StatusCode: 1, Message: "更新日志失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminUpdateLogResp{
		StatusCode: 0,
		Message:    "更新日志成功",
	})
}

// AdminDeleteLog godoc
//
//	@Summary		管理员删除日志
//	@Description	管理员删除日志
//	@Tags			admin-log
//	@Produce		json
//	@Param			log	body		api.AdminDeleteLogReq	true	"日志ID"
//	@Success		200	{object}	api.AdminDeleteLogResp
//	@Router			/admin/log/delete [post]
func AdminDeleteLog(ctx *gin.Context) {
	var req api.AdminDeleteLogReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminDeleteLogResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	err := biz.AdminDeleteLog(ctx, req.LogID)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("删除日志失败")
		ctx.JSON(http.StatusOK, api.AdminDeleteLogResp{StatusCode: 1, Message: "删除日志失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminDeleteLogResp{
		StatusCode: 0,
		Message:    "删除日志成功",
	})
}

// AdminBatchDeleteLogs godoc
//
//	@Summary		管理员批量删除日志
//	@Description	管理员批量删除日志
//	@Tags			admin-log
//	@Produce		json
//	@Param			log	body		api.AdminBatchDeleteLogReq	true	"日志ID列表"
//	@Success		200	{object}	api.AdminBatchDeleteLogResp
//	@Router			/admin/log/batchDelete [post]
func AdminBatchDeleteLogs(ctx *gin.Context) {
	var req api.AdminBatchDeleteLogReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminBatchDeleteLogResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	err := biz.AdminBatchDeleteLogs(ctx, req.LogIDs)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("批量删除日志失败")
		ctx.JSON(http.StatusOK, api.AdminBatchDeleteLogResp{StatusCode: 1, Message: "批量删除日志失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminBatchDeleteLogResp{
		StatusCode: 0,
		Message:    "批量删除日志成功",
	})
}
