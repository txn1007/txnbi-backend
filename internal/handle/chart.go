package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"slices"
	"txnbi-backend/api"
	"txnbi-backend/errs"
	"txnbi-backend/internal/biz"
	"txnbi-backend/pkg/tlog"
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
		tlog.L.Debug().Msgf("生成图表失败，原因：%s", err.Error())
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrGenerateChartFailed.Error()})
		return
	}

	// 校验参数
	// 检查文件大小是否超过 16MB 大小限制
	fileSize := req.File.Size
	if fileSize < 0 || fileSize > MAX_USER_UPLOAD_FILE_SIZE {
		tlog.L.Debug().Msgf("生成图表失败，原因：%s，原始数据：文件大小：%d", "文件大小不合法", fileSize)
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrFileSizeInvalid.Error()})
		return
	}
	// 	检查文件后缀格式是否合法
	fileName := req.File.Filename
	ext := filepath.Ext(fileName)
	if ext != ".xlsx" && ext != ".xls" && ext != ".csv" {
		tlog.L.Debug().Msgf("生成图表失败，原因：%s，原始数据：文件名：%s,文件后缀：%s", "文件后缀格式不支持", fileName, ext)
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrFileExtensionNotSupported.Error()})
		return
	}

	// 检查用户表名、表类型、分析目标的值长度是否在合法范围内
	// 有 gorm 已经有参数化查询，所以在这就不针对 SQL注入 做检查
	goal, chartName := req.Goal, req.ChartName
	goalLen, chartNameLen := len(goal), len(chartName)
	if goalLen < 2 || goalLen > 255 {
		tlog.L.Debug().Msgf("生成图表失败，原因：%s，原始数据：分析目标：%s，分析目标长度：%d", "分析目标字符串长度不合法", goal, goalLen)
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrGoalCharacterCountOutOfRange.Error()})
		return
	}
	if chartNameLen < 1 || chartNameLen > 127 {
		tlog.L.Debug().Msgf("生成图表失败，原因：%s，原始数据：表名：%s，表名长度：%d", "表名长度不合法", chartName, chartNameLen)
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrTableNameLengthOutOfRange.Error()})
		return
	}

	chartType := req.ChartType
	allChartSupportType := []string{"折线图", "柱状图", "堆叠图", "饼图", "雷达图"}
	if !slices.Contains(allChartSupportType, chartType) {
		tlog.L.Debug().Msgf("生成图表失败，原因：%s，原始数据：图表类型：%s", "图表类型不支持", chartType)
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrUnsupportedChartType.Error()})
		return
	}

	data, analysis, err := biz.GenChart(ctx, req.ChartName, req.ChartType, req.Goal, req.File, ctx.GetInt64("userID"))
	if err != nil {
		tlog.L.Debug().Msgf("生成图表失败，原因：%s，原始数据：%v", err.Error(), req)
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: errs.ErrGenerateChartFailed.Error()})
		return
	}
	tlog.L.Debug().Msgf("生成图表成功，原始数据：%s", req)
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
		tlog.L.Debug().Msgf("查找我的图表失败，原因：%s", err.Error())
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrFindMyChartFailed.Error()})
		return
	}
	// 校验参数
	if req.PageSize < 1 || req.PageSize > 32 {
		tlog.L.Debug().Msgf("查找我的图表失败，原因：%s，原始数据：页面大小：%d", "页面大小不合法！", req.PageSize)
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrInvalidPageSize.Error()})
		return
	}
	chartName := req.ChartName
	chartNameLen := len(chartName)
	if chartNameLen < 0 || chartNameLen > 127 {
		tlog.L.Debug().Msgf("查找我的图表失败，原因：%s,表名：%s，表名长度：%d", "表名长度不合法，", chartName, chartNameLen)
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrTableNameLengthOutOfRange.Error()})
		return
	}
	currentPage := req.CurrentPage
	if currentPage < 0 {
		tlog.L.Debug().Msgf("查找我的图表失败，原因：%s,当前页面参数：%d", "当前页面参数不合法", currentPage)
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrInvalidPageParameters.Error()})
		return
	}

	chart, total, err := biz.ListMyChart(ctx, ctx.GetInt64("userID"), req.ChartName, req.CurrentPage, req.PageSize)
	if err != nil {
		tlog.L.Debug().Msgf("查找我的图表失败，原因：%s，原始数据：%v", err.Error(), req)
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: errs.ErrFindMyChartFailed.Error()})
		return
	}
	tlog.L.Debug().Msgf("查找我的图表成功，初始数据：%v", req)
	ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 0, Message: "查询成功！", Charts: chart, Total: total})
	return
}

// DeleteMyChart godoc
//
//	@Summary		删除图表数据接口
//	@Description	删除图表数据接口
//	@Tags			chart
//	@Produce		json
//	@Param			Info	formData	api.DeleteMyChartReq	true	"查询信息"
//	@Success		200		{object}	api.DeleteMyChartResp
//	@Router			/chart/auth/myChartDel [post]
func DeleteMyChart(ctx *gin.Context) {
	var req api.DeleteMyChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		tlog.L.Debug().Msgf("删除我的图表失败，原因：%s", err.Error())
		ctx.JSON(http.StatusOK, api.DeleteMyChartResp{StatusCode: 1, Message: errs.ErrDeleteMyChartFailed.Error()})
		return
	}
	userID := ctx.GetInt64("userID")
	err := biz.DeleteMyChart(ctx, req.ChartID, userID)
	if err != nil {
		tlog.L.Debug().Msgf("删除我的图表失败，原因：%s,原始数据:%v", err.Error(), req)
		ctx.JSON(http.StatusOK, api.DeleteMyChartResp{StatusCode: 1, Message: errs.ErrDeleteMyChartFailed.Error()})
		return
	}
	tlog.L.Debug().Msgf("删除我的图表成功，原始数据：%v", req)
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
		tlog.L.Debug().Msgf("获取示例图表失败，原因：%s", err.Error())
		ctx.JSON(http.StatusOK, api.ExampleChartResp{StatusCode: 1, Message: errs.ErrGetExampleChartFailed.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.ExampleChartResp{StatusCode: 0, Message: "查询成功！", Charts: charts, Total: total})
	return
}
