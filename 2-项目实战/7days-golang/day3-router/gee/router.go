package gee

import "strings"

// router 结构体包含两个字段：roots 和 handlers
// roots 用于存储每种 HTTP 方法的 trie 树的根节点
// handlers 用于存储每个路由对应的处理函数
// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// newRouter 函数创建并返回一个新的 router 实例
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 函数接收一个 URL 路径模式，将其分割成多个部分，并返回这些部分
// 如果路径模式中包含通配符 "*"，那么只返回到 "*" 为止的部分
func parsePattern(pattren string) []string {
	vs := strings.Split(pattren, "/")
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

// addRoute 方法接收一个 HTTP 方法、一个 URL 路径模式和一个处理函数，将它们添加到路由器中
// 首先，它将 URL 路径模式分割成多个部分，然后在对应 HTTP 方法的 trie 树中插入这些部分
// 最后，它将 URL 路径模式和处理函数添加到 handlers 中
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// getRoute 方法接收一个 HTTP 方法和一个 URL 路径，然后在对应 HTTP 方法的 trie 树中搜索这个路径
// 如果找到了匹配的节点，那么它将返回这个节点和路径中的参数
// 如果没有找到匹配的节点，那么它将返回 nil
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
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

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(404, "404 NOT FOUND: %s\n", c.Path)
	}
}
