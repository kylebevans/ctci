package main

import (
  "testing"
)

func TestRouteBetweenNodesBFS(t *testing.T) {
  var g *Graph
  g = NewGraph()

  g.Insert("a", []string{"b", "c"})
  g.Insert("b", []string{"d"})
  g.Insert("c", []string{"e"})
  g.Insert("d", []string{"f"})
  g.Insert("e", []string{"f"})
  g.Insert("f", []string{})
  found, err := g.RouteBetweenNodesBFS("a", "f")
  if err != nil {
    t.Errorf("%v\n", err)
  }
  if !found {
    t.Errorf("RouteBetweenNodes should be true but it is false\n")
  }
}

func TestRouteBetweenNodesDFS(t *testing.T) {
  var g *Graph
  g = NewGraph()

  g.Insert("a", []string{"b", "c"})
  g.Insert("b", []string{"d"})
  g.Insert("c", []string{"e"})
  g.Insert("d", []string{"f"})
  g.Insert("e", []string{"f"})
  g.Insert("f", []string{})
  found, err := g.RouteBetweenNodesDFS("a", "f")
  if err != nil {
    t.Errorf("%v\n", err)
  }
  if !found {
    t.Errorf("RouteBetweenNodes should be true but it is false\n")
  }
}