package main

import (
	"log"
	"os"

	"realizr.io/graph/model"
)

func main() {
	args := os.Args[1:]
	graphname := getStringOrNil(args, 0)
	g := model.Load(graphname)
	log.Printf("Graph name: %s", graphname)
	if g == nil {
		g = &model.Graph{Name: graphname}
	}

	switch getStringOrNil(args, 1) {
	case "add":
		// graph "graphname" "add" "name" "stereotype" "description"
		g.AddNode(getStringOrNil(args, 2), getStringOrNil(args, 3), getStringOrNil(args, 4))
		g.Save()
		break
	case "ren":
		// graph "graphname" "ren" "name" "name"
		g.RenameNode(getStringOrNil(args, 2), getStringOrNil(args, 3))
		g.Save()
		break
	case "del":
		// graph "graphname" "del" "name"
		g.DeleteNode(getStringOrNil(args, 2))
		g.Save()
		break
	case "link":
		// graph "graphname" "link" "name" "reltype" "name"
		g.Connect(getStringOrNil(args, 2), getStringOrNil(args, 3), getStringOrNil(args, 4))
		g.Save()
		break
	case "unlink":
		// graph "graphname" "unlink" "name" "reltype" "name"
		g.Disconnect(getStringOrNil(args, 2), getStringOrNil(args, 3), getStringOrNil(args, 4))
		g.Save()
		break
	default:
		log.Printf("Unknown command %s", getStringOrNil(args, 1))
	}

}

func getStringOrNil(args []string, pos int) string {
	if pos >= len(args) {
		return ""
	} else {
		return args[pos]
	}
}

// graph "graphname" "tag" "name" "tags"
// graph "graphname" "taglink" "name" "name" "tags"
// graph "graphname" "comment" "name"
