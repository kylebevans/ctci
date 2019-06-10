package main

import (
  "testing"
)

func TestAppend(t *testing.T) {
  var test *LinkedList
  test = new(LinkedList)
  for i := 1; i <= 5; i++ {
    test.Append(i)
  }
  if test.head.data != 1 && test.head.next.data != 2 && test.head.next.next.data != 3 && test.head.next.next.next.data != 4 && test.head.next.next.next.next.data != 5 {
    t.Errorf("list should contain 12345, actual: %v%v%v%v%v\n", test.head.data, test.head.next.data, test.head.next.next.data, test.head.next.next.next.data, test.head.next.next.next.next.data)
  }
}

func TestDelete(t *testing.T) {
  var test *LinkedList
  test = new(LinkedList)
  for i := 1; i <= 5; i++ {
    test.Append(i)
  }
  test.Delete(5)
  if test.head.data != 1 && test.head.next.data != 2 && test.head.next.next.data != 3 && test.head.next.next.next.data != 4 && test.head.next.next.next.next != nil {
    t.Errorf("LinkedList should be 1234 with next = nil, actual: %v%v%v%v with next = %v\n", test.head.data, test.head.next.data, test.head.next.next.data, test.head.next.next.next.data, test.head.next.next.next.next)
  }
}

func TestRemoveDups(t *testing.T) {
  var test *LinkedList
  test = new(LinkedList)
  for i := 1; i <= 5; i++ {
    test.Append(i)
  }
  test.Append(3)
  test.Append(3)
  test.Append(4)
  test.Append(5)
  test.Delete(5)
  if test.head.data != 1 && test.head.next.data != 2 && test.head.next.next.data != 3 && test.head.next.next.next.data != 4 && test.head.next.next.next.next.data != 5 && test.head.next.next.next.next.next != nil {
    t.Errorf("list should contain 12345 with next = nil, actual: %v%v%v%v%v with next = %v\n", test.head.data, test.head.next.data, test.head.next.next.data, test.head.next.next.next.data, test.head.next.next.next.next.data, test.head.next.next.next.next.next)
  }
}