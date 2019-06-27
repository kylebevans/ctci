package main

import (
  "errors"
)

type Graph struct {
  numNodes int
  root int
  edges map[string][]string
}

func NewGraph() *Graph {
  g := new(Graph)
  g.edges = make(map[string][]string)
  return g
}

func (g *Graph) Insert(name string, e []string) {
  g.edges[name] = e
  g.numNodes++
}

func (g *Graph) Delete(name string) {
  delete(g.edges, name)
  g.numNodes--
}

type QueueNode struct {
  data string
}

type FIFOQueue struct {
  fq []QueueNode
}

func (f *FIFOQueue) Enqueue(name string) {
  //create queue node for string and append to fifoqueue
  var qn QueueNode
  qn.data = name
  f.fq = append(f.fq, qn)
}

func (f *FIFOQueue) Dequeue() string {
  //create a var and set it to the first value of the fifoqueue
  var qn QueueNode = f.fq[0]

  //slice the fifoqueue from 1 to the end
  f.fq = f.fq[1:]

  //return the var
  return qn.data
}

func (f *FIFOQueue) Len() int {
  return len(f.fq)
}

func (g Graph) RouteBetweenNodesUndirected(start string, finish string) bool {
  // perform bidirectional BFS search
  var startVisited map[string]bool
  var finishVisited map[string]bool

  startVisited = make(map[string]bool)
  startVisited[start] = true
  finishVisited = make(map[string]bool)
  finishVisited[finish] = true

  var startNodeTracker *FIFOQueue
  startNodeTracker = new(FIFOQueue)
  var finishNodeTracker *FIFOQueue
  finishNodeTracker = new(FIFOQueue)

  startNodeTracker.Enqueue(start)
  finishNodeTracker.Enqueue(finish)

  for startNodeTracker.Len() > 0 && finishNodeTracker.Len() > 0 {
    var startCurrentNode = startNodeTracker.Dequeue()
    for _, nodey := range g.edges[startCurrentNode] {
      if finishVisited[nodey] {
        return true
      }
      if !startVisited[nodey] {
        startVisited[nodey] = true
        startNodeTracker.Enqueue(nodey)
      }
    }

    var finishCurrentNode = finishNodeTracker.Dequeue()
    for _, nodey := range g.edges[finishCurrentNode] {
      if startVisited[nodey] {
        return true
      }
      if !finishVisited[nodey] {
        finishVisited[nodey] = true
        finishNodeTracker.Enqueue(nodey)
      }
    }
  }
  return false
}

func (g Graph) RouteBetweenNodesBFS(start string, finish string) (bool, error) {
  var visited map[string]bool
  visited = make(map[string]bool)

  _, ok := g.edges[start]
  if !ok {
    return false, errors.New("Start is not in graph")
  }

  _, ok = g.edges[finish]
  if !ok {
    return false, errors.New("Finish is not in graph")
  }

  if start == finish {
    return true, nil
  }

  var nodeTracker *FIFOQueue = new(FIFOQueue)
  visited[start] = true
  nodeTracker.Enqueue(start)
  
  for nodeTracker.Len() > 0 {
    currentNode := nodeTracker.Dequeue()
    for _, nodey := range g.edges[currentNode] {
      if nodey == finish {
        return true, nil
      }
      if !visited[nodey] {
        nodeTracker.Enqueue(nodey)
        visited[nodey] = true
      }
    }
  }
  return false, nil
}


func (g Graph) RouteBetweenNodesDFS(start string, finish string) (bool, error) {
  // create a visited map
  var visited map[string]bool
  visited = make(map[string]bool)
  
  _, ok := g.edges[start]
  if !ok {
    return false, errors.New("Start is not in graph")
  }

  _, ok = g.edges[finish]
  if !ok {
    return false, errors.New("Finish is not in graph")
  }

  if start == finish {
    return true, nil
  }

  // call the actual DFS Search
  return g.DFS(start, finish, visited), nil
}

func (g Graph) DFS(start string, finish string, visited map[string]bool) bool {
  //mark the current node as visited
  visited[start] = true
  var found bool
  
  //recurse through the rest of the current nodes edges
  for _, nodey := range g.edges[start] {
    if nodey == finish {
      return true
    }
    if !visited[nodey] {
      found = g.DFS(nodey, finish, visited)
    }
  }
  return found
}