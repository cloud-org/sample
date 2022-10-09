// author: ashing
// time: 2020/3/28 10:59 下午
// mail: axingfly@gmail.com
// Less is more.

package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				// path = /hello/panda
				// pattern = /hello/:name
				// 赋值 :name 例如 name = panda
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	log.Printf("%+v %d\n", nodes, len(nodes))
	return nodes
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		// print route
		log.Printf("Method: %s\tPattern: %s\tPath: %s\n", c.Method, n.pattern, c.Path)
		c.Params = params
		key := c.Method + "-" + n.pattern
		//r.handlers[key](c)
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		// print route
		log.Printf("NotFound Method: %s\tPath: %s\n", c.Method, c.Path)
		//c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}

	c.Next()
}
