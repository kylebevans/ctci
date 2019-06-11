package main

type Node struct {
  prev *Node
  next *Node
  data int
}

type LinkedList struct {
  head *Node
  tail *Node
}

func (l *LinkedList) Append(d int) {
  lnew := &Node{data: d}

  if l.head == nil && l.tail == nil {
    l.head = lnew
    l.tail = lnew
    return
  }
  currentNode := l.head
  
  for currentNode.next != nil {
    currentNode = currentNode.next
  }
  currentNode.next = lnew
  lnew.prev = currentNode
  l.tail = lnew
  return
}

func (l *LinkedList) Delete(d int) *LinkedList {
  var currentNode *Node = l.head

  if currentNode.data == d {
    currentNode.next.prev = nil
    l.head = currentNode.next
    return l
  }

  for currentNode.next != nil {
    if currentNode.next.data == d {
      currentNode.next = currentNode.next.next
      if currentNode.next != nil {
        currentNode.next.prev = currentNode
      } else {
        l.tail = currentNode
      }
      return l
    }
    currentNode = currentNode.next
  }
  return l
}

func (l *LinkedList) RemoveDups() *LinkedList {
  var hits map[int][]*(Node)
  hits = make(map[int][]*(Node))

  currentNode := l.head

  hits[currentNode.data] = append(hits[currentNode.data], currentNode)

  for currentNode.next != nil {
    currentNode = currentNode.next
    if len(hits[currentNode.data]) > 0 {
      if currentNode == l.tail {
        l.tail = currentNode.prev
        currentNode.prev.next = nil
      } else {
        currentNode.prev.next = currentNode.next
        currentNode.next.prev = currentNode.prev
      }
    }
    hits[currentNode.data] = append(hits[currentNode.data], currentNode)
  }
  return l
}

func (l *LinkedList) RemoveDupsNoBuffer() *LinkedList {
  currentNode := l.head
  for currentNode.next != nil {
    innerNode := currentNode
    for innerNode.next != nil {
      if innerNode.next.data == currentNode.data {
        if innerNode.next == l.tail {
          l.tail = innerNode
          innerNode.next = nil
          continue
        } else {
          innerNode.next.next.prev = innerNode
          innerNode.next = innerNode.next.next
        }
        continue
      }
      innerNode = innerNode.next
    }
    if currentNode == l.tail {
      return l
    } else {
      currentNode = currentNode.next
    }
  }
  return l
}

 