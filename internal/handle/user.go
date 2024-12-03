package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"txnbi-backend/api"
	"txnbi-backend/internal/biz"
)

// UserLogin godoc
//
//	@Summary		用户登陆接口
//	@Description	登陆界面中的用户登陆接口
//	@Tags			user
//	@Produce		json
//	@Param			account		formData	string	true	"用户账号"
//	@Param			password	formData	string	true	"用户密码"
//	@Success		200			{string}	{string}
//	@Router			/user/login [post]
func UserLogin(ctx *gin.Context) {
	var req api.UserLoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: "参数不合法！"})
		return
	}

	err := biz.UserLogin(req.Account, req.Password)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 1, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.UserLoginResp{StatusCode: 0, Message: "登陆成功"})
}

// UserRegister godoc
//
//	@Summary		用户注册接口
//	@Description	登陆界面中的用户注册接口
//	@Tags			user
//	@Produce		json
//	@Param			account		formData	string	true	"用户账号"
//	@Param			password	formData	string	true	"用户密码"
//	@Success		200			{string}	{string}
//	@Router			/user/register [post]
func UserRegister(ctx *gin.Context) {
	var req api.UserRegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, api.UserResisterResp{StatusCode: 1, Message: err.Error()})
		return
	}

	err := biz.UserRegister(req.Account, req.Password)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, api.UserResisterResp{StatusCode: 1, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.UserResisterResp{StatusCode: 0, Message: "注册成功！"})
}
