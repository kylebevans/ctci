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
    if currentNode.data == d {
      currentNode.prev.next = currentNode.next
      currentNode.next.prev = currentNode.prev
      return l
    }
    currentNode = currentNode.next
  }
  return l
}

func (l *LinkedList) RemoveDups() *LinkedList {
  var hits map[int][]*(Node)

  currentNode := l.head
  for currentNode.next != nil {
    hits[currentNode.data] = append(hits[currentNode.data], currentNode)
  }

  for _, v := range hits {
    for i := 0; i < len(v)-1; i++ {
      if v[i] == l.head {
        l.head = v[i].next
        v[i].next.prev = nil
        continue
      }
      if v[i] == l.tail {
        l.tail = v[i].prev
        v[i].prev.next = nil
        continue
      }
      v[i].prev.next = v[i].next
      v[i].next.prev = v[i].prev
    }
  }
  return l
}
 