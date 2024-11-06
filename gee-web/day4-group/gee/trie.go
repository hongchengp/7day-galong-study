package gee

import "strings"

type node struct {
	pattern  string // 从 起始到结束，so 只有终结节点有, 用户注册的路径，可以实现动态路由
	part     string // 节点匹配的部分
	children []*node
	isWild   bool // 如果是 * or : 那 就是true
}

func newNode(part string) *node {
	return &node{
		part:   part,
		isWild: len(part) > 0 && (part[0] == '*' || part[0] == ':'),
	}
}

// 插入，只需要一个就可以了
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 用来 search，要将所有的符合的节点找
func (n *node) matchChildren(part string) []*node {
	children := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}
	return children
}

// 插入，只用插入一个 子节点就可以了 因为 
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	child := n.matchChild(parts[height])
	if child == nil {
		child = newNode(parts[height])
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height + 1)
}

// 传入 http url路径 找到对应的pattern then pattern 找到 handler
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	children := n.matchChildren(parts[height])
	for _, child := range children {
		result := child.search(parts, height + 1)
		if result != nil {
			return result
		}
	}
	return nil
}