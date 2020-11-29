package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由 例如 /p/:lang
	part     string  // 路由中的一部分 :lang
	children []*node // 子节点,例如 [doc,tutorial, intro]
	isWild   bool    // 是否需要精确匹配, part 含有:或者*为true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	// 判断层数, 只有最后一层, 添加 pattern
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 根据层数,获取节点块
	part := parts[height]
	// 根据节点块 判断是否已经创建过节点(返回第一个命中的)
	child := n.matchChild(part)
	if child == nil {
		// 还没创建,则创建一个
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 递归插入下一层
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 如果已经判断到了最后一层, 或者本层的规则包含 *
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 如果 规则链还不是最后一层, 则返回空,命中失败
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

// 第一个匹配成功的节点,用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// 所有匹配成功的节点, 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
