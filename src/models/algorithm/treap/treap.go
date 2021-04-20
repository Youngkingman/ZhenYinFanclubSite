package algorithm

import "math/rand"

type TNode struct {
	ch       [2]*TNode
	priority int
	val      int
}

func (o *TNode) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *TNode) rotate(d int) *TNode {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type Treap struct {
	root *TNode
}

func (t *Treap) _put(o *TNode, val int) *TNode {
	if o == nil {
		return &TNode{priority: rand.Int(), val: val}
	}
	d := o.cmp(val)
	o.ch[d] = t._put(o.ch[d], val)
	if o.ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	return o
}

func (t *Treap) Put(val int) {
	t.root = t._put(t.root, val)
}

func (t *Treap) _delete(o *TNode, val int) *TNode {
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *Treap) Delete(val int) {
	t.root = t._delete(t.root, val)
}

func (t *Treap) LowerBound(val int) (lb *TNode) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.ch[0]
		case c > 0:
			o = o.ch[1]
		default:
			return o
		}
	}
	return
}

func ContainsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	set := &Treap{}
	for i, v := range nums {
		if lb := set.LowerBound(v - t); lb != nil && lb.val <= v+t {
			return true
		}
		set.Put(v)
		if i >= k {
			set.Delete(nums[i-k])
		}
	}
	return false
}
