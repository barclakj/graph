package model

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndLoad(t *testing.T) {
	g := Graph{}
	g.Name = "bob"

	g.AddNode("one", "type a", "some description")
	g.AddNode("two", "type a", "")
	g.AddNode("three", "type a", "three description")
	_, n := g.GetNode("three")
	assert.Equal(t, "three description", n.Description)

	g.Connect("one", "dep", "two")
	g.Connect("three", "assoc", "two")
	g.Connect("three", "assoc", "two")

	g.Save()
	assert.Equal(t, 2, g.LinkLength())
	assert.Equal(t, 3, g.Size())

	g2 := Load("bob")

	assert.Equal(t, 3, g2.Size())

	_, n = g2.GetNode("three")
	assert.Equal(t, "three description", n.Description)
	g2.AddNode("three", "type a", "three description UPDATED")
	_, n = g2.GetNode("three")
	assert.Equal(t, "three description UPDATED", n.Description)
}

func TestDeleteLink(t *testing.T) {
	g := Graph{}
	g.Name = "bob"

	g.Connect("one", "dep", "two")
	g.Connect("one", "dep", "three")

	assert.Equal(t, 2, g.LinkLength())
	log.Printf("A: %s", g.ToJSON())

	g.Disconnect("alpha", "dep", "two")
	assert.Equal(t, 2, g.LinkLength())

	g.Disconnect("one", "dep", "two")
	log.Printf("B: %s", g.ToJSON())
	assert.Equal(t, 1, g.LinkLength())

	g.Connect("one", "dep", "three")
	log.Printf("C: %s", g.ToJSON())
	assert.Equal(t, 0, g.Size())
}

func TestDeleteNode(t *testing.T) {
	g := Graph{}
	g.Name = "bob"

	g.AddNode("one", "type a", "some description")
	g.AddNode("two", "type a", "some description")
	g.AddNode("three", "type a", "three description")

	assert.Equal(t, 3, g.Size())
	log.Printf("A: %s", g.ToJSON())

	g.DeleteNode("two")
	log.Printf("B: %s", g.ToJSON())
	assert.Equal(t, 2, g.Size())
	_, n := g.GetNode("one")
	assert.NotNil(t, n)
	_, n = g.GetNode("two")
	assert.Nil(t, n)
	_, n = g.GetNode("three")
	assert.NotNil(t, n)

	g.DeleteNode("one")
	log.Printf("C: %s", g.ToJSON())
	assert.Equal(t, 1, g.Size())

	g.DeleteNode("three")
	log.Printf("D: %s", g.ToJSON())
	assert.Equal(t, 0, g.Size())
}

func TestRenameNode(t *testing.T) {
	g := Graph{}
	g.Name = "bob"

	g.AddNode("one", "type a", "some description")

	_, n := g.GetNode("one")
	assert.NotNil(t, n)
	g.RenameNode("one", "two")

	_, n = g.GetNode("two")
	assert.NotNil(t, n)

	_, n = g.GetNode("one")
	assert.Nil(t, n)
}
