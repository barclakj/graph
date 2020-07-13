package model

import (
	"encoding/json"
	"io/ioutil"
)

type Graph struct {
	Nodes []*Node
	Links []*Link
	Name  string
}

type Node struct {
	Name        string
	Stereotype  string
	Description string
	Tags        []string
	Comment     []string
}

type Link struct {
	FromName string
	ToName   string
	Reltype  string
	Tags     []string
}

func (g *Graph) GetLink(fromName string, reltype string, toName string) (int, *Link) {
	for idx, link := range g.Links {
		if link.FromName == fromName && link.ToName == toName && link.Reltype == reltype {
			return idx, link
		}
	}
	return -1, nil
}

func (g *Graph) Disconnect(fromName string, reltype string, toName string) {
	idx, _ := g.GetLink(fromName, reltype, toName)
	if idx >= 0 {
		links := append(g.Links[0:idx], g.Links[idx+1:]...)
		g.Links = links
	}
}

func (g *Graph) Connect(fromName string, reltype string, toName string) *Link {
	_, l := g.GetLink(fromName, reltype, toName)
	if l == nil {
		l = &Link{FromName: fromName, ToName: toName, Reltype: reltype}
		g.Links = append(g.Links, l)
	}
	return l
}

func (g *Graph) GetNode(name string) (int, *Node) {
	for idx, n := range g.Nodes {
		if n.Name == name {
			return idx, n
		}
	}
	return -1, nil
}

func (g *Graph) DeleteNode(name string) {
	idx, _ := g.GetNode(name)
	if idx >= 0 {
		nodes := append(g.Nodes[0:idx], g.Nodes[idx+1:]...)
		g.Nodes = nodes
	}
}

func (g *Graph) AddNode(name string, stereotype string, description string) *Node {
	_, n := g.GetNode(name)
	if n == nil {
		n = &Node{Name: name}
		g.Nodes = append(g.Nodes, n)
	}
	n.Stereotype = stereotype
	n.Description = description
	return n
}

func (g *Graph) RenameNode(name string, newname string) {
	_, node := g.GetNode(name)
	if node != nil {
		node.Name = newname
	}
}

func (g *Graph) Size() int {
	return len(g.Nodes)
}

func (g *Graph) LinkLength() int {
	return len(g.Links)
}

func (g *Graph) ToJSON() string {
	jsonMap, _ := json.Marshal(g)
	return string(jsonMap)
}

func FromJSON(jsonString []byte) *Graph {
	g := Graph{}
	json.Unmarshal(jsonString, &g)
	return &g
}

func (g *Graph) Save() {
	ioutil.WriteFile(g.Name+".graph", []byte(g.ToJSON()), 0644)
}

func Load(filename string) *Graph {
	dat, _ := ioutil.ReadFile(filename + ".graph")
	if len(dat) == 0 {
		return nil
	} else {
		return FromJSON(dat)
	}
}
