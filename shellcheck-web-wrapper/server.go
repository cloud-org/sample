package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SC struct {
	Script string `json:"script"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Resp(data interface{}) *Response {
	return &Response{Code: 0, Data: data}
}

func FailedResp(msg string) *Response {
	return &Response{Code: -1, Msg: msg}
}

func addMiddleware(e *echo.Echo) {
	// 增加 cors 中间件
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func SCHandler(c echo.Context) error {
	var sc SC
	err := c.Bind(&sc)
	if err != nil {
		return c.JSON(http.StatusOK, FailedResp("请输入正确的请求体"))
	}
	if sc.Script == "" {
		return c.JSON(http.StatusOK, FailedResp("script 不能为空"))
	}
	json1, err := ShellCheck(sc.Script)
	if err != nil {
		return c.JSON(http.StatusOK, FailedResp(err.Error()))
	}
	return c.JSON(http.StatusOK, json1)
}

func addApi(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		v1.POST("/shellcheck", SCHandler)
	}
}

//CreateEngine echo 实例
func CreateEngine() (*echo.Echo, error) {
	e := echo.New()

	addMiddleware(e)
	addApi(e)

	return e, nil
}
