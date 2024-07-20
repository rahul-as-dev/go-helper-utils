package utils

type BinaryTree[T comparable] struct {
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
	Value T
}

func NewBinaryTree[T comparable](val T) *BinaryTree[T] {
	return &BinaryTree[T]{Value: val}
}

// IsEquivalent determines whether the trees
// t1 and t2 contain the same values.
func IsEquivalent[T comparable](t1, t2 *BinaryTree[T]) bool {
	ch1 := make(chan T)
	ch2 := make(chan T)
	go walking(t1, ch1)
	go walking(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func (t *BinaryTree[T]) NodeCount() int {
	if t == nil {
		return 0
	}
	return 1 + t.Left.NodeCount() + t.Right.NodeCount()
}

func (t *BinaryTree[T]) inOrder(f func(T)) {
	if t == nil {
		return
	}
	t.Left.inOrder(f)
	f(t.Value)
	t.Right.inOrder(f)
}
func (t *BinaryTree[T]) preOrder(f func(T)) {
	if t == nil {
		return
	}
	f(t.Value)
	t.Left.preOrder(f)
	t.Right.preOrder(f)
}
func (t *BinaryTree[T]) postOrder(f func(T)) {
	if t == nil {
		return
	}
	t.Left.postOrder(f)
	t.Right.postOrder(f)
	f(t.Value)
}

func (t *BinaryTree[T]) search(val T) bool {
	if t == nil {
		return false
	}
	if val == t.Value {
		return true
	}
	return t.Left.search(val) || t.Right.search(val)
}

func (t *BinaryTree[T]) searchByComparator(val T, compare func(t, val T) bool) bool {
	if t == nil {
		return false
	}
	if compare(t.Value, val) {
		return true
	}
	return t.Left.search(val) || t.Right.search(val)
}

func walking[T comparable](t *BinaryTree[T], ch chan T) {
	walk(t, ch)
	defer close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk[T comparable](t *BinaryTree[T], ch chan T) {
	// inorder traversal (left,val,right)
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}
