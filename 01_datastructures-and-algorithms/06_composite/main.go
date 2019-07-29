package main

import (
	"fmt"
)

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("leaflet", leaf.name)
}

type Branch struct {
	leafs []Leaflet
	name string
	branches []Branch
}

func (br *Branch) perform() {
	fmt.Println("Branch", br.name)
	for _, leaf := range br.leafs {
		leaf.perform()
	}

	for _, branch := range br.branches {
		branch.perform()
	}
}

func (br *Branch) add(leaf Leaflet) {
	br.leafs = append(br.leafs, leaf)
}

func (br *Branch) addBranch(newBranch Branch) {
	br.branches = append(br.branches, newBranch)
}

func (br *Branch) getLeaflets() []Leaflet {
	return br.leafs
}

func main() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}
	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perform()

}
