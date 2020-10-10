package gee

import (
	"strings"
)

type node struct {
	pattern  string // 待匹配路由
	value    string
	children []*node
	isWild   bool // 是否精确匹配
}

func (n *node) insert(pattern string, values []string, height int) {
	if len(values) == height {
		n.pattern = pattern
		return
	}

	value := values[height]
	child := n.matchChild(value)
	if child == nil {
		child = &node{value: value, isWild: value[0] == ':' || value[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, values, height+1)
}

func (n *node) search(values []string, height int) *node {
	if len(values) == height || strings.HasPrefix(n.value, "*") {
		if n.pattern == "" {
			return nil
		} else {
			return n
		}
	}

	value := values[height]
	children := n.matchChildren(value)

	for _, child := range children {
		result := child.search(values, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) matchChild(value string) *node {
	for _, child := range n.children {
		if child.value == value || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(value string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.value == value || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
