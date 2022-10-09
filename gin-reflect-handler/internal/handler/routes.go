package handler

import (
	"gin-reflect-handler/internal/logic"
	"gin-reflect-handler/internal/svc"
	"gin-reflect-handler/internal/types"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func addApi(r *gin.Engine, serverCtx *svc.ServiceContext) {

	v1 := r.Group("/v1/api/cmdb")
	{

		h := Handler{svcCtx: serverCtx}
		targetRouter := v1.Group("/target")
		{
			targetHandler := TargetHandler{H: &h}
			targetRouter.GET("/attr", targetHandler.H.ReflectHandler(
				logic.NewTargetLogic,
				&types.AttrListReq{},
				"GetAttrList",
			))
		}
	}

}

// 增加 Recovery 中间件
func addMiddlewareRecovery(r *gin.Engine) {
	r.Use(gin.Recovery())
}

// 增加 cors 中间件
func addMiddlewareCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "DELETE", "PATCH", "PATCH"},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}

//CreateEngine gin
func CreateEngine(serverCtx *svc.ServiceContext) (*gin.Engine, error) {
	if !serverCtx.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	addMiddlewareRecovery(r)
	addMiddlewareCors(r)
	addApi(r, serverCtx)

	return r, nil
}
