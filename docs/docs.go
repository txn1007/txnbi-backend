// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chart/exampleChart": {
            "get": {
                "description": "用户获取自己的图表数据接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chart"
                ],
                "summary": "用户获取自己的图表数据接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ExampleChartResp"
                        }
                    }
                }
            }
        },
        "/chart/findMyChart": {
            "get": {
                "description": "用户获取自己的图表数据接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chart"
                ],
                "summary": "用户获取自己的图表数据接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "chartName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "currentPage",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.FindMyChartResp"
                        }
                    }
                }
            }
        },
        "/chart/gen": {
            "post": {
                "description": "AI生成图表数据接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chart"
                ],
                "summary": "AI生成图表数据接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表名",
                        "name": "chartName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "表类型",
                        "name": "chartType",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询目标",
                        "name": "goal",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "用户上传的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GenChartResp"
                        }
                    }
                }
            }
        },
        "/chart/myChartDel": {
            "post": {
                "description": "删除图表数据接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chart"
                ],
                "summary": "删除图表数据接口",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "chartID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.DeleteMyChartResp"
                        }
                    }
                }
            }
        },
        "/user/CurrentUserDetail": {
            "get": {
                "description": "用户自身详情接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户自身详情接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.CurrentUserDetailResp"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "登陆界面中的用户登陆接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户登陆接口",
                "parameters": [
                    {
                        "description": "登陆信息",
                        "name": "LoginInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UserLoginResp"
                        }
                    }
                }
            }
        },
        "/user/loginOut": {
            "post": {
                "description": "用户自身详情接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户退出登陆接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UserLoginOutResp"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "登陆界面中的用户注册接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "注册信息",
                        "name": "RegisterInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UserRegisterResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ChartInfoV0": {
            "type": "object",
            "properties": {
                "chartCode": {
                    "type": "string"
                },
                "chartGoal": {
                    "type": "string"
                },
                "chartID": {
                    "type": "integer"
                },
                "chartName": {
                    "type": "string"
                },
                "chartResult": {
                    "type": "string"
                },
                "chartType": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                }
            }
        },
        "api.CurrentUserDetailResp": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "required: true",
                    "type": "string"
                },
                "statusCode": {
                    "description": "required: true",
                    "type": "integer"
                },
                "userInfoV0": {
                    "$ref": "#/definitions/api.UserInfoV0"
                }
            }
        },
        "api.DeleteMyChartResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "api.ExampleChartResp": {
            "type": "object",
            "properties": {
                "charts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ChartInfoV0"
                    }
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "api.FindMyChartResp": {
            "type": "object",
            "properties": {
                "charts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ChartInfoV0"
                    }
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "api.GenChartResp": {
            "type": "object",
            "properties": {
                "genChart": {
                    "type": "string"
                },
                "genResult": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "api.UserInfoV0": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userAccount": {
                    "description": "账号",
                    "type": "string"
                },
                "userAvatar": {
                    "description": "用户头像",
                    "type": "string"
                },
                "userName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "userRole": {
                    "description": "用户角色：user/admin",
                    "type": "string"
                }
            }
        },
        "api.UserLoginOutResp": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "required: true",
                    "type": "string"
                },
                "statusCode": {
                    "description": "required: true",
                    "type": "integer"
                }
            }
        },
        "api.UserLoginReq": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "api.UserLoginResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api.UserRegisterReq": {
            "type": "object",
            "required": [
                "account",
                "inviteCode",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "inviteCode": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "api.UserRegisterResp": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "required: true",
                    "type": "string"
                },
                "statusCode": {
                    "description": "required: true",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "txnbi API",
	Description:      "txnbi 的 API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
