package main

import (
  "testing"
)

func TestIsRotation(t *testing.T) {
  isRot := IsRotation("erbottlewat", "waterbottle")
  if !isRot {
    t.Errorf("IsRotation(\"erbottlewat\", \"waterbottle\") = %v, want true.\n", isRot)
  }
}
