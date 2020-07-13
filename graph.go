package main

import (
	"encoding/json"
	"log"
	"os"
)

type Graph struct {
	Nodes []Node
	Links []Link
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

func main() {
	argsWithoutProg := os.Args[1:]
	log.Print("Hello world\n")
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

// graph "graphname" "add" "stereotype" "name" "description"
// graph "graphname" "ren" "name" "name"
// graph "graphname" "link" "name" "reltype" "name"
// graph "graphname" "del" "name"
// graph "graphname" "unlink" "name" "reltype" "name"
// graph "graphname" "tag" "name" "tags"
// graph "graphname" "taglink" "name" "name" "tags"
// graph "graphname" "comment" "name"
