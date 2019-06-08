package main

import (
  "strings"
)

func IsRotation(s1 string, s2 string) bool {
  if len(s1) != len(s2) {
    return false
  }
  
  var sb strings.Builder
  sb.WriteString(s2)
  sb.WriteString(s2)

  return strings.Contains(sb.String(), s1)
}
