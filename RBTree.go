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
	if r.root == nil {
		z := &node{
			parent: &sentinel, 
			left: &sentinel,
			right: &sentinel,
			Key: k,
			Value: v,
		}
		r.Root = z
	} else {
		r.insert(k, v)
	}
}

func (r *RBTree) insert(k int, v int){
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
			} else if z == z.parent.right {
				z = z.parent
				r.left_rotate(z)
			} else {
				z.parent.colour = BLACK
				z.parent.parent.colour = RED
				r.right_rotate(z.parent.parent)
			}
		} else {
			
		}
	}
	r.root.colour = BLACK
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
