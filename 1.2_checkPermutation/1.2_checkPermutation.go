package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "strings"
)


func quickSort(b []byte) []byte {
  if len(b) <= 1 {
    return b
  }
  var pivot int = len(b) - 1
  lowmark := 0
  for i := 0; i < len(b) - 1; i++ {
    if b[i] <= b[pivot] {
      temp := b[lowmark]
      b[lowmark] = b[i]
      b[i] = temp
      lowmark++
    }
  }
  temp := b[lowmark]
  b[lowmark] = b[pivot]
  b[pivot] = temp
  bSorted := append(quickSort(b[0:lowmark]), b[lowmark])
  bSorted = append(bSorted, quickSort(b[lowmark+1:])...)
  return bSorted
}

func checkPermutation(s1 string, s2 string) bool {
  // abababcd  abbaabcd
  if len(s1) != len(s2) {
    return false
  }
  var b1 = []byte(s1)
  var b2 = []byte(s2)
  b1Sorted := quickSort(b1)
  b2Sorted := quickSort(b2)
  fmt.Printf("String1: %v\n", string(b1Sorted))
  fmt.Printf("String2: %v\n", string(b2Sorted))

  if string(b1Sorted) == string(b2Sorted) {
    return true
  } else {
    return false
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  
  fmt.Println("Enter string 1")
  s1, err := reader.ReadString('\n')
  if err != nil {
    log.Fatalf("Reading string 1 failed: ", err)
  }
  s1 = strings.TrimRight(s1, "\n")
  
  fmt.Println("Enter string 2")
  s2, err := reader.ReadString('\n')
  if err != nil {
    log.Fatalf("Reading string 2 failed: ", err)
  }
  s2 = strings.TrimRight(s2, "\n")

  var isPerm bool
  isPerm = checkPermutation(s1, s2)

  if isPerm {
    fmt.Println("String 2 is a permutation of string 1.")
  } else {
    fmt.Println("String 2 is not a permutation of string 1.")
  }
  return
}
  
 
    