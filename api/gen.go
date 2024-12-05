package api

import "mime/multipart"

type GenChartReq struct {
	Token string `json:"token" query:"token" form:"token" binding:"required"`

	ChartName string                `json:"chartName" form:"chartName" binding:"required"`
	ChartType string                `json:"chartType" form:"chartType" binding:"required"`
	Goal      string                `json:"goal" form:"goal" binding:"required"`
	File      *multipart.FileHeader `json:"file" form:"file" binding:"required"`
}

type GenChartResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`

	GenChart  string `json:"genChart" form:"genChart"`
	GenResult string `json:"genResult" form:"genResult"`
}
