package handle

import (
	"net/http"
	"txnbi-backend/api"
	"txnbi-backend/errs"
	biz "txnbi-backend/internal/biz/admin"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// AdminListCharts godoc
//
//	@Summary		管理员获取图表列表接口
//	@Description	管理员获取图表列表接口
//	@Tags			admin-chart
//	@Produce		json
//	@Param			Info	query		api.AdminChartListReq	true	"查询信息"
//	@Success		200		{object}	api.AdminChartListResp
//	@Router			/admin/chart/list [get]
func AdminListCharts(ctx *gin.Context) {
	var req api.AdminChartListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminChartListResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	// 校验参数
	if req.PageSize < 1 || req.PageSize > 50 {
		log.Info().Err(errs.ErrInvalidPageSize).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.AdminChartListResp{StatusCode: 1, Message: errs.ErrInvalidPageSize.Error()})
		return
	}

	if req.Page < 1 {
		log.Info().Err(errs.ErrInvalidPageParameters).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.AdminChartListResp{StatusCode: 1, Message: errs.ErrInvalidPageParameters.Error()})
		return
	}

	charts, total, err := biz.AdminListCharts(ctx, req.Page, req.PageSize, req.Keyword, req.UserID)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("获取图表列表失败")
		ctx.JSON(http.StatusOK, api.AdminChartListResp{StatusCode: 1, Message: "获取图表列表失败"})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminChartListResp{
		StatusCode: 0,
		Message:    "获取图表列表成功",
		Total:      total,
		Charts:     charts,
	})
}

// AdminGetChartDetail godoc
//
//	@Summary		管理员获取图表详情接口
//	@Description	管理员获取图表详情接口
//	@Tags			admin-chart
//	@Produce		json
//	@Param			chartID	query		int64	true	"图表ID"
//	@Success		200		{object}	api.AdminChartDetailResp
//	@Router			/admin/chart/detail [get]
func AdminGetChartDetail(ctx *gin.Context) {
	var req api.AdminChartDetailReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminChartDetailResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	if req.ChartID <= 0 {
		log.Info().Err(errs.ErrInvalidInputParameters).Interface("req", req).Msg("图表ID无效")
		ctx.JSON(http.StatusOK, api.AdminChartDetailResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	chart, err := biz.AdminGetChartDetail(ctx, req.ChartID)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("获取图表详情失败")
		ctx.JSON(http.StatusOK, api.AdminChartDetailResp{StatusCode: 1, Message: "获取图表详情失败"})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminChartDetailResp{
		StatusCode: 0,
		Message:    "获取图表详情成功",
		Chart:      *chart,
	})
}

// AdminUpdateChart godoc
//
//	@Summary		管理员更新图表接口
//	@Description	管理员更新图表接口
//	@Tags			admin-chart
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			chartID		formData	int64	true	"图表ID"
//	@Param			chartName	formData	string	false	"图表名称"
//	@Param			chartGoal	formData	string	false	"分析目标"
//	@Param			genResult	formData	string	false	"分析结果"
//	@Success		200			{object}	api.AdminUpdateChartResp
//	@Router			/admin/chart/update [post]
func AdminUpdateChart(ctx *gin.Context) {
	var req api.AdminUpdateChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminUpdateChartResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	if req.ChartID <= 0 {
		log.Info().Err(errs.ErrInvalidInputParameters).Interface("req", req).Msg("图表ID无效")
		ctx.JSON(http.StatusOK, api.AdminUpdateChartResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	err := biz.AdminUpdateChart(ctx, req.ChartID, req.ChartName, req.ChartGoal, req.ChartResult)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("更新图表失败")
		ctx.JSON(http.StatusOK, api.AdminUpdateChartResp{StatusCode: 1, Message: "更新图表失败"})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminUpdateChartResp{
		StatusCode: 0,
		Message:    "更新图表成功",
	})
}

// AdminDeleteChart godoc
//
//	@Summary		管理员删除图表接口
//	@Description	管理员删除图表接口
//	@Tags			admin-chart
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			chartID	formData	int64	true	"图表ID"
//	@Success		200		{object}	api.AdminDeleteChartResp
//	@Router			/admin/chart/delete [post]
func AdminDeleteChart(ctx *gin.Context) {
	var req api.AdminDeleteChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("参数绑定失败")
		ctx.JSON(http.StatusOK, api.AdminDeleteChartResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	if req.ChartID <= 0 {
		log.Info().Err(errs.ErrInvalidInputParameters).Interface("req", req).Msg("图表ID无效")
		ctx.JSON(http.StatusOK, api.AdminDeleteChartResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	err := biz.AdminDeleteChart(ctx, req.ChartID)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("删除图表失败")
		ctx.JSON(http.StatusOK, api.AdminDeleteChartResp{StatusCode: 1, Message: "删除图表失败"})
		return
	}

	ctx.JSON(http.StatusOK, api.AdminDeleteChartResp{
		StatusCode: 0,
		Message:    "删除图表成功",
	})
}
