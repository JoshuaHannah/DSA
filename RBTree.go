package DSA

import (
	"fmt"
)

type colour int

const (
	RED colour = iota
	BLACK
)

type node struct {
	parent *node
	left *node
	right *node
	colour Colour
	//TODO: Change Key and Value to be any ordered type.
	Key int
	Value int
}

type RBTree struct {
	Root *Node
	Size int
}

var sentinel = &node{}

func New() *RBTree {
	tree := &RBTree{
		Root: &sentinel,
	}
	sentinel.parent = tree
	return tree
}

func (r *RBTree) Insert(k int, v int){
	y := &sentinel
	x := r.Root
	for x != &sentinel {
		y = x
		if k < x.Key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z := &node{
		parent: y,
		left: &sentinel,
		right: &sentinel,
		Key: k,
		Value: v,
	}
	if y == &sentinel {
		r.Root = z
	} else if k < y.key {
		y.left = z
	} else {
		y.right = z
	}
	z.left = &sentinel
	z.right = &sentinel
	z.colour = RED
	r.Size = r.Size + 1
	insert-fixup(z)
}

func (r *RBTree) insert-fixup(z *node){
	for z.parent.colour == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.colour == RED {
				z.parent.colour = BLACK
				y.colour = BLACK
				z.parent.parent.colour = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					r.left_rotate(z)
				}
				z.parent.colour = BLACK
				z.parent.parent.colour = RED
				r.right_rotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.colour == RED {
				z.parent.colour = BLACK
				y.colour = BLACK
				z.parent.parent.colour = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					r.right_rotate(z)
				}
				z.parent.colour = BLACK
				z.parent.parent.colour = RED
				r.left_rotate(z.parent.parent)
			}
		}
	}
	r.Root.colour = BLACK
}

func (r *RBTree) Delete(key int){
	
}

func (r *RBTree) Lookup(key int){
	
}

func (r *RBTree) Clear(){
	r.Root = &sentinel
	r.Size = 0
}

func (r *RBTree) Empty() {
	return r.Root == nil
}

func (r *RBTree) left-rotate(z *node){
	y := z.right
	z.right = y.left
	if y.left != &sentinel{
		y.left.parent = z
	}
	y.parent = z.parent
	if z.parent == &sentinel {
		r.Root = y
	} else if z == z.parent.left {
		z.parent.left = y
	} else {
		z.parent.right = y
	}
	y.left = z
	z.parent = y
}

func (r *RBTree) right-rotate(y *node){
	z := y.left
	y.left = z.right
	z.right = y
	z.parent = y.parent
	if y.parent == &sentinel {
		r.Root = z
	} else y.parent.left == y {
		y.parent.left = z
	} else {
		y.parent.right = z
	}
	y.parent = z
}