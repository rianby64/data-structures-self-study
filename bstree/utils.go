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

func delete(t, parent *bstree) BStree {
	// case when: t.left == nil && t.right == nil -> leaf
	if t.left == nil && t.right == nil {
		if parent.left == t {
			parent.left = nil
		}
		if parent.right == t {
			parent.right = nil
		}

		return parent
	}

	// case else:

	// find max on left
	if t.left != nil {
		maxleft := findmax(t.left)
		if maxleft != nil {
			t.payload = maxleft.payload
			delete(maxleft, maxleft.parent)
			return t
		}
	}

	// find max on right
	if t.right != nil {
		maxright := findmax(t.right)
		if maxright != nil {
			t.payload = maxright.payload
			delete(maxright, maxright.parent)
			return t
		}
	}

	return nil
}

func find(a interface{}, t *bstree, matcher comparator) BStree {
	if matcher(a, t.Value()) {
		return t
	}

	if t.left != nil {
		if t.comparator(a, t.left.Value()) {
			found := find(a, t.left, matcher)
			if found != nil {
				return found
			}
		}
	}

	if t.right != nil {
		return find(a, t.right, matcher)
	}

	return nil
}

func findmax(t *bstree) *bstree {
	right := t.right
	if right != nil {
		return findmax(right)
	}

	return t
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
