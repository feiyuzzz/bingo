package framework

import (
	"errors"
	"strings"
)

// 构建一个前缀树来实现动态路由匹配

type Tree struct {
	root *node
}

type node struct {
	isLast  bool
	segment string
	handler ControllerHandler
	childs  []*node // 此节点的子节点
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root}
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

// 判断一个 segment 是否是一个通用的 segment

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足 segment 规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// 如果 segment 是通配符，那么所有下一层子节点都符合要求
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	// 过滤所有的下一层子节点
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			// 如果下一层子节点有通配符，则满足要求
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			// 下一层子节点不是通配符，但是文本完全匹配，也满足要求
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

// 判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	// 用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层子节点
	cnodes := n.filterChildNodes(segment)
	// 没有一个符合，说明这个 uri 一定不存，直接返回 nil
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}
	// 只有一个 segment，则是最后一个
	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}

	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// 添加路由

func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist:" + uri)
	}

	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1
		var objNode *node
		childNodes := n.filterChildNodes(segment)
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}
		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode
	}
	return nil
}

func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
