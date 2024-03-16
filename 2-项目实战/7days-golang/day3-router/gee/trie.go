package gee

import "strings"

// 定义 node 结构体，代表 trie 树的一个节点
type node struct {
	pattern  string  // 待匹配路由，例如 "p/:lang/doc"
	part     string  // 路由中的一部分，例如 "p", ":lang", "doc"
	children []*node // 子节点，例如 [doc, blog]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// matchChild 函数在子节点中找到第一个匹配 "part" 的节点
// 如果子节点的 part 与输入的 part 相同，或者子节点是通配符节点（isWild 为 true），则返回该子节点
// 如果没有找到匹配的子节点，则返回 nil
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || n.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 函数在子节点中找到所有匹配 "part" 的节点
// 如果子节点的 part 与输入的 part 相同，或者子节点是通配符节点（isWild 为 true），则将该子节点添加到结果列表中
// 最后返回包含所有匹配子节点的列表
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 方法用于向 trie 树中插入一个节点
// pattern 是待插入的完整路径，parts 是路径分割后的各部分，height 是当前处理到 parts 的哪一部分
func (n *node) insert(pattern string, parts []string, height int) {
	// 如果所有的 parts 都已处理完，那么在当前节点存储 pattern
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 取出当前要处理的 part
	part := parts[height]
	// 在当前节点的子节点中查找是否已存在该 part
	child := n.matchChild(part)
	// 如果不存在，则新建一个子节点
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 在找到的（或新建的）子节点上，递归地插入剩余的 parts
	child.insert(pattern, parts, height+1)
}

// search 方法用于在 trie 树中搜索一个路径
// parts 是待搜索路径分割后的各部分，height 是当前处理到 parts 的哪一部分
func (n *node) search(parts []string, height int) *node {
	// 如果所有的 parts 都已处理完，或者当前节点包含通配符
	// 那么如果当前节点存储了 pattern，则返回当前节点，否则返回 nil
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	// 取出当前要处理的 part
	part := parts[height]
	// 在当前节点的子节点中查找所有匹配的节点
	children := n.matchChildren(part)
	// 在找到的所有子节点上，递归地搜索剩余的 parts
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	// 如果没有找到匹配的节点，返回 nil
	return nil
}
