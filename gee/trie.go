// author: ashing
// time: 2020/3/28 11:51 下午
// mail: axingfly@gmail.com
// Less is more.

package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool // part 含有 : 或 * 时为true
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	// 递归终止条件
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil { // 为 nil 则新建一个 child
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	// 继续下一个
	child.insert(pattern, parts, height+1)

}

// 查找路由
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
