package bstree

import "github.com/rianby64/data-structures-self-study/list"

func insert(t *bstree, v interface{}, c func(a, b interface{}) bool) BStree {
	t.length++
	if t.payload == nil {
		t.SetValue(v)
		return t
	}

	node := &bstree{
		root:       t.root,
		comparator: c,
		parent:     t,
	}

	if c(t.payload.Value(), v) {
		if t.left == nil {
			t.left = node
		}
		return insert(t.left, v, c)
	}
	if t.right == nil {
		t.right = node
	}
	return insert(t.right, v, c)
}

func insertNode(t *bstree, node *bstree, c func(a, b interface{}) bool) BStree {
	t.length += node.Length()

	if c(t.payload.Value(), node.Value()) {
		if t.left == nil {
			t.left = node
			return node
		}
		return insert(t.left, node, c)
	}
	if t.right == nil {
		t.right = node
		return node
	}
	return insert(t.right, node, c)
}

func find(a interface{}, t *bstree, matcher, comparator comparator) BStree {
	if matcher(a, t.Value()) {
		return t
	}

	if t.left != nil {
		if t.comparator(a, t.left.Value()) {
			found := find(a, t.left, matcher, comparator)
			if found != nil {
				return found
			}
		}
	}

	if t.right != nil {
		return find(a, t.right, matcher, comparator)
	}

	return nil
}

func inorder(root *bstree, l list.List) {
	// la condicion || (root.root == root && root.Value() == nil) es por falta de sentinela
	if root == nil || (root.root == root && root.Value() == nil) {
		return
	}

	if root.left != nil {
		inorder(root.left, l)
	}

	l.Last().Insert(root.Value())

	if root.right != nil {
		inorder(root.right, l)
	}
}
