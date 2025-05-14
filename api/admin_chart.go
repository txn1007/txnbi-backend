package api

// AdminChartListReq 管理员获取图表列表请求
type AdminChartListReq struct {
	Page     int    `form:"page" json:"page" binding:"required,min=1"`
	PageSize int    `form:"pageSize" json:"pageSize" binding:"required,min=1,max=100"`
	Keyword  string `form:"keyword" json:"keyword"`
	UserID   int64  `form:"userId" json:"userId"`
}

// AdminChartListResp 管理员获取图表列表响应
type AdminChartListResp struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Total      int64         `json:"total"`
	Charts     []ChartInfoV1 `json:"charts"`
}

// ChartInfoV1 图表信息（管理员视角）
type ChartInfoV1 struct {
	ChartID     int64  `json:"chartId"`
	UserID      int64  `json:"userId"`
	UserAccount string `json:"userAccount"`
	ChartName   string `json:"chartName"`
	ChartType   string `json:"chartType"`
	ChartGoal   string `json:"chartGoal"`
	ChartCode   string `json:"chartCode"`
	ChartResult string `json:"chartResult"`
	Status      string `json:"status"`
	UpdateTime  string `json:"updateTime"`
}

// AdminChartDetailReq 管理员获取图表详情请求
type AdminChartDetailReq struct {
	ChartID int64 `form:"chartId" json:"chartId" binding:"required,min=1"`
}

// AdminChartDetailResp 管理员获取图表详情响应
type AdminChartDetailResp struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Chart      ChartInfoV1 `json:"chart"`
}

// AdminUpdateChartReq 管理员更新图表请求
type AdminUpdateChartReq struct {
	ChartID     int64  `json:"chartId" binding:"required,min=1"`
	ChartName   string `json:"chartName"`
	ChartGoal   string `json:"chartGoal"`
	ChartResult string `json:"chartResult"`
}

// AdminUpdateChartResp 管理员更新图表响应
type AdminUpdateChartResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// AdminDeleteChartReq 管理员删除图表请求
type AdminDeleteChartReq struct {
	ChartID int64 `json:"chartId" binding:"required,min=1"`
}

// AdminDeleteChartResp 管理员删除图表响应
type AdminDeleteChartResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}