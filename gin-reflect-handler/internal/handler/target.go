package handler

//TargetHandler 嵌入 H 可以自定义自己的路由 handler 如果有需要的话
type TargetHandler struct {
	H *Handler
}
