package api

type UserLoginReq struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`
	Token      string `json:"token" form:"token"`
}

type UserRegisterReq struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserResisterResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`
	Token      string `json:"token" form:"token"`
}
