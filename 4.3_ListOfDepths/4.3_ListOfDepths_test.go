package main

import (
  "testing"
)

func TestListOfDepths(t *testing.T) {
  testTree := NewGraph()
  testTree.Insert(10, []int{5, 20})
  testTree.Insert(5, []int{10, 12})
  testTree.Insert(20, []int{10, 3, 7})
  testTree.Insert(12, []int{5})
  testTree.Insert(3, []int{9, 18})
  testTree.Insert(7, []int{20})
  testTree.Insert(9, []int{3})
  testTree.Insert(18, []int{3})
  testTree.root = 10

  var zeroLL, oneLL, twoLL, threeLL LinkedList
  zeroLL.Append(10)
  oneLL.Append(5)
  oneLL.Append(20)
  twoLL.Append(12)
  twoLL.Append(3)
  twoLL.Append(7)
  threeLL.Append(9)
  threeLL.Append(18)
  
  listOfDepths := make(map[int]*LinkedList)
  listOfDepths = ListOfDepths(*testTree)
  if len(listOfDepths) == 0 {
    t.Errorf("listOfDepths is empty\n")
  }
  for k, v := range listOfDepths {
    switch k {
    case 0:
      if !LLCompare(*v, zeroLL) {
        t.Errorf("level 0 failed, needed: %v, had: %v\n", v, zeroLL)
      }
    case 1:
      if !LLCompare(*v, oneLL) {
        t.Errorf("level 1 failed, needed: %v, had: %v\n", v, oneLL)
      }
    case 2:
      if !LLCompare(*v, twoLL) {
        t.Errorf("level 2 failed, needed: %v, had: %v\n", v, twoLL)
      }
    case 3:
      if !LLCompare(*v, threeLL) {
        t.Errorf("level 3 failed, needed: %v, had: %v\n", v, threeLL)
      }
    }
  }
}