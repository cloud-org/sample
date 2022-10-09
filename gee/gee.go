// author: ashing
// time: 2020/3/28 11:45 上午
// mail: axingfly@gmail.com
// Less is more.

package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups

	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	engine      *Engine       // all groups share a Engine instance
}

// Default use Logger() & Recovery middlewares
func Default() *Engine {
	engine := New()
	engine.Use(Recovery(), Logger())
	return engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	// 注意：这里是 g.addRoute 并非 g.engine.addRoute, 下面 POST 同理
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	// 将各个与请求路径前缀相同的 group 的 middlewares 加入其中
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	c.engine = e
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	//for k, v := range e.router.handlers {
	//	fmt.Println(k, v)
	//}
	return http.ListenAndServe(addr, e)
}

func (g *RouterGroup) Static(relativePath, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	g.GET(urlPattern, handler)
}

// 创建静态文件服务器 handler
func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath) // absolutePath: /asserts
	// http.StripPrefix 用法: https://studygolang.com/articles/9197
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
