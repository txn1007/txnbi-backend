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

type UserRegisterResp struct {
	// required: true
	// example: 0
	StatusCode int `json:"statusCode" form:"statusCode"`
	// required: true
	// example: 登陆成功
	Message string `json:"message" form:"message"`
}
