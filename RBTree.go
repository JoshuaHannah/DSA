// ief
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
	Parent *Node
	Left *Node
	Right *Node
	//TODO: Change Key and Value to be any ordered type.
	Key int
	Value int
}

type RBTree struct {
	Root *Node
	Size int
}

func New() *RBTree {
	return &RBTree{}
}

func (r *RBTree) Insert(){
	
}

func (r *RBTree) Delete(key int){
	
}

func (r *RBTree) Lookup(ket int){
	
}

func (r *RBTree) Clear(){
	r.Root = nil
	r.Size = 0
}

func (r *RBTree) Empty() {
	return r.Root == nil
}
