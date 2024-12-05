package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"txnbi-backend/api"
	"txnbi-backend/internal/biz"
)

// GenChart godoc
//
//	@Summary		AI生成图表数据接口
//	@Description	AI生成图表数据接口
//	@Tags			chart
//	@Produce		json
//	@Param			token		formData	string	true	"用户token"
//	@Param			chartName	formData	string	true	"表名"
//	@Param			chartType	formData	string	true	"表类型"
//	@Param			goal		formData	string	true	"查询目标"
//	@Param			chartData	formData	file	true	"用户上传的文件"
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

	data, analysis, err := biz.GenChart(req.ChartName, req.ChartType, req.Goal, req.File, ctx.GetInt64("userID"))
	if err != nil {
		ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 1, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.GenChartResp{StatusCode: 0, Message: "生成成功！", GenChart: data, GenResult: analysis})
	return
}
