package api

import "mime/multipart"

type GenChartReq struct {
	Token string `json:"token" query:"token" form:"token"`

	ChartName string                `json:"chartName" form:"chartName"`
	ChartType string                `json:"chartType" form:"chartType"`
	Goal      string                `json:"goal" form:"goal"`
	ChartData *multipart.FileHeader `json:"chartData" form:"chartData"`
}

type GenChartResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`

	GenChart  string `json:"genChart" form:"genChart"`
	GenResult string `json:"genResult" form:"genResult"`
}
