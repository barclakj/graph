package model

import (
	"encoding/json"
	"fmt"
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
	FromName    string
	ToName      string
	Reltype     string
	Description string
	Tags        []string
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

func (g *Graph) Connect(fromName string, reltype string, toName string, description string) *Link {
	_, l := g.GetLink(fromName, reltype, toName)
	if l == nil {
		l = &Link{FromName: fromName, ToName: toName, Reltype: reltype}
		g.Links = append(g.Links, l)
	}
	l.Description = description
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
	for _, l := range g.Links {
		if l.FromName == name {
			l.FromName = newname
		}
		if l.ToName == name {
			l.ToName = newname
		}
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

func (g *Graph) Show() {
	fmt.Printf("Graph %s. Nodes %d. Links %d\n", g.Name, g.Size(), g.LinkLength())
	fmt.Print("------------------------\n")
	for idx, n := range g.Nodes {
		fmt.Printf("%d: %s [%s] == %s\n", idx, n.Name, n.Stereotype, n.Description)
	}
	fmt.Print("------------------------\n")
	for idx, l := range g.Links {
		fmt.Printf("%d: %s -(%s)-> %s == %s\n", idx, l.FromName, l.Reltype, l.ToName, l.Description)
	}
}

func (g *Graph) Graphviz() {
	fmt.Printf("digraph g { \n\tgraph [fontsize=30 labelloc=\"t\" label=\"\" splines=true overlap=false rankdir = \"LR\"];\n\n")

	for _, n := range g.Nodes {
		fmt.Printf("\t\"%s\" [ style=\"filled\"  penwidth = 1 fillcolor = \"white\" fontname = \"Courier New\" shape = \"Mrecord\" label = \"≪%s≫\\n%s\"];\n", n.Name, n.Stereotype, n.Name)
	}

	fmt.Print("\n")

	for _, l := range g.Links {
		if l.Description != "" {
			fmt.Printf("\t\"%s\" -> \"%s\" [ penwidth = 1 fontsize = 10 fontcolor = \"black\" label = \"≪%s≫\\n%s\" ];\n", l.FromName, l.ToName, l.Reltype, l.Description)
		} else {
			fmt.Printf("\t\"%s\" -> \"%s\" [ penwidth = 1 fontsize = 10 fontcolor = \"black\" label = \"≪%s≫\" ];\n", l.FromName, l.ToName, l.Reltype)
		}
	}
	fmt.Printf("}\n")
}
