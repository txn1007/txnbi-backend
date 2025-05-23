package handle

import (
	"net/http"
	"path/filepath"
	"slices"
	"txnbi-backend/api"
	"txnbi-backend/errs"
	biz "txnbi-backend/internal/biz/user"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// GenChart godoc
//
//	@Summary		AI生成图表数据接口
//	@Description	AI生成图表数据接口
//	@Tags			chart
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			token		formData	string	true	"用户token"
//	@Param			chartName	formData	string	true	"表名"
//	@Param			chartType	formData	string	true	"表类型"
//	@Param			goal		formData	string	true	"查询目标"
//	@Param			file		formData	file	true	"用户上传的文件"
//	@Success		200			{object}	api.GenChartResp
//	@Router			/chart/auth/gen [post]
func GenChart(ctx *gin.Context) {
	var req api.GenChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrGenerateChartFailed.Error()})
		return
	}

	// 校验参数
	// 检查文件大小是否超过 16MB 大小限制
	fileSize := req.File.Size
	if fileSize < 0 || fileSize > MAX_USER_UPLOAD_FILE_SIZE {
		log.Info().Err(errs.ErrFileSizeInvalid).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrFileSizeInvalid.Error()})
		return
	}
	// 	检查文件后缀格式是否合法
	fileName := req.File.Filename
	ext := filepath.Ext(fileName)
	if ext != ".xlsx" && ext != ".xls" && ext != ".csv" {
		log.Info().Err(errs.ErrFileExtensionNotSupported).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrFileExtensionNotSupported.Error()})
		return
	}

	// 检查用户表名、表类型、分析目标的值长度是否在合法范围内
	// 有 gorm 已经有参数化查询，所以在这就不针对 SQL注入 做检查
	goal, chartName := req.Goal, req.ChartName
	goalLen, chartNameLen := len(goal), len(chartName)
	if goalLen < 2 || goalLen > 255 {
		log.Info().Err(errs.ErrGoalCharacterCountOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrGoalCharacterCountOutOfRange.Error()})
		return
	}
	if chartNameLen < 1 || chartNameLen > 127 {
		log.Info().Err(errs.ErrTableNameLengthOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrTableNameLengthOutOfRange.Error()})
		return
	}

	chartType := req.ChartType
	allChartSupportType := []string{"折线图", "柱状图", "堆叠图", "饼图", "雷达图"}
	if !slices.Contains(allChartSupportType, chartType) {
		log.Info().Err(errs.ErrUnsupportedChartType).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrUnsupportedChartType.Error()})
		return
	}

	data, analysis, err := biz.GenChart(ctx, req.ChartName, req.ChartType, req.Goal, req.File, ctx.GetInt64("userID"))
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrGenerateChartFailed.Error()})
		return
	}
	log.Info().Interface("req", req).Msg("生成图表成功")
	ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 0, Message: "生成成功！", GenChart: data, GenResult: analysis})
	return
}

// FindMyChart godoc
//
//	@Summary		用户获取自己的图表数据接口
//	@Description	用户获取自己的图表数据接口
//	@Tags			chart
//	@Produce		json
//	@Param			Info	query		api.FindMyChartReq	true	"查询信息"
//	@Success		200		{object}	api.FindMyChartResp
//	@Router			/chart/auth/findMyChart [get]
func FindMyChart(ctx *gin.Context) {
	var req api.FindMyChartReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrFindMyChartFailed.Error()})
		return
	}
	// 校验参数
	if req.PageSize < 1 || req.PageSize > 32 {
		log.Info().Err(errs.ErrInvalidPageSize).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrInvalidPageSize.Error()})
		return
	}
	chartName := req.ChartName
	chartNameLen := len(chartName)
	if chartNameLen < 0 || chartNameLen > 127 {
		log.Info().Err(errs.ErrTableNameLengthOutOfRange).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrTableNameLengthOutOfRange.Error()})
		return
	}
	currentPage := req.CurrentPage
	if currentPage < 0 {
		log.Info().Err(errs.ErrInvalidPageParameters).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrInvalidPageParameters.Error()})
		return
	}

	chart, total, err := biz.ListMyChart(ctx, ctx.GetInt64("userID"), req.ChartName, req.CurrentPage, req.PageSize)
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrFindMyChartFailed.Error()})
		return
	}
	log.Info().Interface("req", req).Msg("查找我的图表成功")
	ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 0, Message: "查询成功！", Charts: chart, Total: total})
	return
}

// DeleteMyChart godoc
//
//	@Summary		删除图表数据接口
//	@Description	删除图表数据接口
//	@Tags			chart
//	@Produce		json
//	@Accept			multipart/form-data
//	@Param			token	formData	string	true	"用户token"
//	@Param			userID	formData	string	true	"userID"
//	@Success		200		{object}	api.DeleteMyChartResp
//	@Router			/chart/auth/myChartDel [post]
func DeleteMyChart(ctx *gin.Context) {
	var req api.DeleteMyChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.DeleteMyChartResp{StatusCode: 1, Message: errs.ErrDeleteMyChartFailed.Error()})
		return
	}
	userID := ctx.GetInt64("userID")
	err := biz.DeleteMyChart(ctx, req.ChartID, userID)
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.DeleteMyChartResp{StatusCode: 1, Message: errs.ErrDeleteMyChartFailed.Error()})
		return
	}
	log.Info().Interface("req", req).Msg("删除图表成功")
	ctx.JSON(http.StatusOK, api.DeleteMyChartResp{StatusCode: 0, Message: "删除成功！"})
	return
}

// ExampleChart godoc
//
//	@Summary		用户获取自己的图表数据接口
//	@Description	用户获取自己的图表数据接口
//	@Tags			chart
//	@Produce		json
//	@Param			Info	query		api.ExampleChartReq	true	"查询信息"
//	@Success		200		{object}	api.ExampleChartResp
//	@Router			/chart/exampleChart [get]
func ExampleChart(ctx *gin.Context) {
	charts, total, err := biz.ExampleChart(ctx)
	if err != nil {
		log.Info().Err(err).Msg("")
		ctx.JSON(http.StatusOK, api.ExampleChartResp{StatusCode: 1, Message: errs.ErrGetExampleChartFailed.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.ExampleChartResp{StatusCode: 0, Message: "查询成功！", Charts: charts, Total: total})
	return
}

// UpdateChart godoc
//
//	@Summary		用户修改自己的图表数据接口
//	@Description	用户修改自己的图表数据接口
//	@Tags			chart
//	@Produce		json
//	@Accept			multipart/form-data
//	@Param			token		formData	string	true	"用户token"
//	@Param			chartID		formData	string	true	"图表ID"
//	@Param			chartName	formData	string	true	"图表名"
//	@Param			chartGoal	formData	string	true	"分析目标"
//	@Param			genResult	formData	string	true	"分析结果"
//	@Success		200			{object}	api.UpdateChartResp
//	@Router			/chart/auth/update [post]
func UpdateChart(ctx *gin.Context) {
	var req api.UpdateChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UpdateChartResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}

	userID := ctx.GetInt64("userID")
	err := biz.UpdateChart(ctx, req.ChartID, userID, req.ChartName, req.ChartGoal, req.GenResult)
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UpdateChartResp{StatusCode: 1, Message: errs.ErrUpdateChartFailed.Error()})
		return
	}

	log.Info().Interface("req", req).Msg("")
	ctx.JSON(http.StatusOK, api.UpdateChartResp{StatusCode: 0, Message: "更新图表成功！"})
	return
}

// ShareChart godoc
//
//	@Summary		用户生成分享自己的图表邀请码接口
//	@Description	用户生成分享自己的图表邀请码接口
//	@Tags			chart
//	@Produce		json
//	@Accept			multipart/form-data
//	@Param			token	formData	string	true	"用户token"
//	@Param			chartID	formData	string	true	"图表ID"
//	@Success		200		{object}	api.ShareChartResp
//	@Router			/chart/auth/share [post]
func ShareChart(ctx *gin.Context) {
	var req api.ShareChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.UpdateChartResp{StatusCode: 1, Message: errs.ErrInvalidInputParameters.Error()})
		return
	}
	userID := ctx.GetInt64("userID")
	accessCode, err := biz.ShareChart(ctx, req.ChartID, userID, req.Token)
	if err != nil {
		log.Info().Err(err).Interface("req", req).Msg("")
		ctx.JSON(http.StatusOK, api.ShareChartResp{StatusCode: 1, Message: errs.ErrShareChartFailed.Error()})
		return
	}
	log.Info().Interface("req", req).Msg("")
	ctx.JSON(http.StatusOK, api.ShareChartResp{StatusCode: 0, Message: "生成分享链接成功！", AccessCode: accessCode})
	return
}
