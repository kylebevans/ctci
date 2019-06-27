package main

import (
  "fmt"
  "log"
)

type ListNode struct {
  prev *ListNode
  next *ListNode
  data int
}

type LinkedList struct {
  head *ListNode
  tail *ListNode
  length int
}

func (l LinkedList) Len() int {
  return l.length
}

func (l *LinkedList) Append(i int) {
  newNode := new(ListNode)
  newNode.data = i
  if l.Len() == 0 {
    l.head = newNode
  } else {
    newNode.prev = l.tail
    l.tail.next = newNode
  }
  l.tail = newNode
  l.length++
}

func (l LinkedList) Print() {
  currentNode := l.head
  for currentNode != l.tail {
    fmt.Printf("%v", currentNode.data)
  }
  if currentNode == l.tail {
    fmt.Printf("%v\n", currentNode.data)
  }
  return
}

func LLCompare(l1, l2 LinkedList) bool {
  if l1.Len() != l2.Len() {
    return false
  }
  currentNode1 := l1.head
  currentNode2 := l2.head

  for currentNode1 != nil && currentNode2 != nil {
    if currentNode1.data != currentNode2.data {
      return false
    }
    currentNode1 = currentNode1.next
    currentNode2 = currentNode2.next
  }
  return true
}

type Queue struct {
  queue []int
}

func (q Queue) Len() int {
  return len(q.queue)
}

func (q *Queue) Enqueue(i int) {
  q.queue = append(q.queue, i)
  return
}

func (q *Queue) Dequeue() int {
  i := q.queue[0]
  q.queue = q.queue[1:]
  return i
}

type Graph struct {
  numNodes int
  root int
  edges map[int][]int
}

func NewGraph() *Graph {
  g := new(Graph)
  g.edges = make(map[int][]int)
  return g
}

func (g *Graph) Insert(node int, edges []int) {
  g.numNodes++
  g.edges[node] = edges
  return
}

func (g *Graph) Delete(node int) {
  g.numNodes--
  delete(g.edges, node)
  return
}

func (g Graph) IsInAdjacency(parent int, current int) bool {
  for _, v := range g.edges[parent] {
    if v == current {
      return true
    }
  }
  return false
}

func ListOfDepths(tree Graph) map[int]*LinkedList {
  m := make(map[int]*LinkedList)
  visited := make(map[int]bool)
  var depth int

  kyoo := new(Queue)
  depthq := new(Queue)

  kyoo.Enqueue(tree.root)
  depthq.Enqueue(depth)

  for kyoo.Len() > 0 {
    nodey := kyoo.Dequeue()
    depth = depthq.Dequeue()
    if m[depth] == nil {
      m[depth] = new(LinkedList)
    }
    m[depth].Append(nodey)
    visited[nodey] = true
    for _, v := range tree.edges[nodey] {
      if !visited[v] {
        kyoo.Enqueue(v)
        depthq.Enqueue(depth+1)
      } else if !tree.IsInAdjacency(v, nodey) {
        log.Fatalf("cycle detected: %v\n", v)
      }
    }
  }
  return m
}
