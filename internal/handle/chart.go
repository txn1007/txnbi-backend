package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"txnbi-backend/api"
	"txnbi-backend/internal/biz"
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
//	@Router			/chart/gen [post]
func GenChart(ctx *gin.Context) {
	var req api.GenChartReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: err.Error()})
		return
	}

	// 检查文件大小是否超过 16MB 大小限制
	if req.File.Size > 16*1024*1024 {
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: "file size too big"})
		return
	}
	// 	检查文件后缀格式是否合法
	ext := filepath.Ext(req.File.Filename)
	if ext != ".xlsx" && ext != ".xls" && ext != ".csv" {
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: "file type not supported"})
		return
	}

	data, analysis, err := biz.GenChart(req.ChartName, req.ChartType, req.Goal, req.File, ctx.GetInt64("userID"))
	if err != nil {
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: err.Error()})
		return
	}
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
//	@Router			/chart/findMyChart [get]
func FindMyChart(ctx *gin.Context) {
	var req api.FindMyChartReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: err.Error()})
		return
	}
	// 校验参数
	if req.PageSize < 1 || req.PageSize > 32 {
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: "pageSize不合法"})
		return
	}

	chart, total, err := biz.ListMyChart(ctx.GetInt64("userID"), req.ChartName, req.CurrentPage, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 1, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.FindMyChartResp{StatusCode: 0, Message: "查询成功！", Charts: chart, Total: total})
	return
}