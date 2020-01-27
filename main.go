package main

import (
	"fmt"
)

type Type int 
type Space struct {
	Types map[Type][]Type
	Terminal map[Type]bool
	Commuter [][2][]int
}

type PathTree struct {
	t Type 
	c []*PathTree
}

func (pt *PathTree) Gnos() {
	var GnosAux func(pt *PathTree, tabs int)
	GnosAux = func(pt *PathTree, tabs int) {
		for i := 0; i < tabs; i ++ {
			fmt.Printf("\t")
		}
		fmt.Printf("%d\n", pt.t)
		for _, child := range pt.c {
			GnosAux(child, tabs + 1)
		}
	}
	GnosAux(pt, 0)
}

func (pt *PathTree) Sow() { //for sowing together common pathes
	
}

func (sp *Space) Pathes(d int, t Type) *PathTree {
	var c []*PathTree
	if d > 0 {
		c = make([]*PathTree, len(sp.Types[t]))
		for i, t := range sp.Types[t] {
			c[i] = sp.Pathes(d-1, t)
		}
	} else {
		c = []*PathTree{}
	}
	return &PathTree {
		t: t,
		c: c,
	}
}

func main() {
	tree := Space {
		Types: map[Type][]Type{
			1: []Type{0, 1, 1, 1, 1}, // l, r, u, d
		},
		Terminal: map[Type]bool{
			0: true,
			1: false,
		},
		Commuter: [][2][]int{
			[2][]int{[]int{1, 2}, []int{2, 1}},
			[2][]int{[]int{1, 2}, []int{2, 1}},
		},
	}
	tree.Pathes(3, 1).Gnos()
}