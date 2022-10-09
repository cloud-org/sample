// author: ashing
// time: 2020/3/28 4:33 下午
// mail: axingfly@gmail.com
// Less is more.

package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// 原始对象
	Writer  http.ResponseWriter
	Request *http.Request

	Path   string
	Method string

	Params map[string]string

	StatusCode int

	// middleware
	handlers []HandlerFunc
	index    int

	engine *Engine
}

func newContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Writer:  writer,
		Request: request,
		Path:    request.URL.Path,
		Method:  request.Method,
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers) // 直接将索引移至最后 TODO: 感觉 len(c.handlers) - 1 即可
	c.JSON(code, H{"message": err})
}
