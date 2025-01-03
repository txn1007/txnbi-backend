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

type FindMyChartReq struct {
	Token string `json:"token" query:"token" form:"token" binding:"required"`

	ChartName   string `json:"chartName" form:"chartName" query:"chartName"`
	CurrentPage int    `json:"currentPage" form:"currentPage" query:"currentPage"`
	PageSize    int    `json:"pageSize" form:"pageSize" query:"pageSize"`
}

type FindMyChartResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`

	Total  int64         `json:"total" form:"total"`
	Charts []ChartInfoV0 `json:"charts"`
}

type ChartInfoV0 struct {
	ChartID     int64  `json:"chartID" form:"chartID"`
	ChartName   string `json:"chartName" form:"chartName"`
	ChartGoal   string `json:"chartGoal" form:"chartGoal"`
	ChartType   string `json:"chartType" form:"chartType"`
	ChartCode   string `json:"chartCode" form:"chartCode"`
	ChartResult string `json:"chartResult" form:"chartResult"`

	UpdateTime string `json:"updateTime" form:"updateTime"`
}

type DeleteMyChartReq struct {
	Token string `json:"token" query:"token" form:"token" binding:"required"`

	ChartID int64 `json:"chartID" form:"chartID" binding:"required"`
}

type DeleteMyChartResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`
}

type ExampleChartReq struct {
}

type ExampleChartResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`

	Total  int64         `json:"total" form:"total"`
	Charts []ChartInfoV0 `json:"charts"`
}

type UpdateChartReq struct {
	Token string `json:"token" query:"token" form:"token" binding:"required"`

	ChartID   int64  `json:"chartID" query:"chartID" form:"chartID" binding:"required"`
	ChartName string `json:"chartName" form:"chartName" binding:"required"`
	ChartGoal string `json:"chartGoal" form:"chartGoal"`
	GenResult string `json:"genResult" form:"genResult"`
}

type UpdateChartResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`
}

type ShareChartReq struct {
	Token string `json:"token" query:"token" form:"token" binding:"required"`

	ChartID int64 `json:"chartID" query:"chartID" form:"chartID" binding:"required"`
}

type ShareChartResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`

	AccessCode string `json:"accessCode" form:"accessCode"`
}
