package handler

import (
	"fmt"
	"gin-reflect-handler/internal/svc"
	"log"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svcCtx *svc.ServiceContext
}

//ReflectHandler 反射通用 handler
//fn -> NewTargetLogic 实例化 logic 的函数签名
//req -> 请求体 需要传指针类型 .Elem 需要 Kind 为 Ptr
//callFn -> 调用的函数名称
func (h *Handler) ReflectHandler(fn, req interface{}, callFn string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//  bind req
		reqType := reflect.TypeOf(req).Elem()
		request := reflect.New(reqType)
		if err := c.ShouldBind(request.Interface()); err != nil {
			c.JSON(http.StatusOK, fmt.Sprintf("反序列化失败: %v", err))
			return
		}
		log.Printf("%+v\n", request.Interface())
		// 实例化 logic
		logicParams := []reflect.Value{reflect.ValueOf(c), reflect.ValueOf(h.svcCtx)}
		rets := reflect.ValueOf(fn).Call(logicParams)
		if len(rets) != 1 {
			c.JSON(http.StatusOK, fmt.Sprintf("实例化 logic 失败"))
			return
		}
		log.Printf("rets[0] is %+v\n", rets[0])
		//  Logic.Call func by method name
		reqParams := []reflect.Value{reflect.ValueOf(request.Interface())}
		res := rets[0].MethodByName(callFn).Call(reqParams)
		if len(res) != 2 {
			c.JSON(http.StatusOK, fmt.Sprintf("调用 %s func 失败", callFn))
			return
		}

		if !res[1].IsNil() { // 第二个返回参数是 error
			c.JSON(http.StatusOK, fmt.Sprintf(
				"fn: %+v, req: %+v, callFn: %s, err: %+v",
				fn,
				request.Interface(),
				callFn,
				res[1].Interface()))
			return
		}

		c.JSON(http.StatusOK, bson.M{"data": res[0].Interface()}) // 第一个返回参数是 data
		return
	}
}
