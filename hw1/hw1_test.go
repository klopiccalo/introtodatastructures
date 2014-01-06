package hw1

import (
  "encoding/base64"
  "testing"
)

type SumExample struct {
  Input []int
  Expected int
}

func TestSum(t *testing.T) {
  if !testing.Verbose() {
    t.Fatal("Please run with the verbose flag set (\"go test -test.v\")")
  }
  examples := []SumExample{
    {[]int{1, 2, 3, 4}, 10},
    {[]int{0, 0, 0, 0}, 0},
    {[]int{-1, -2, -3, -4}, -10},
    {[]int{4, 2, 6, -3}, 9},
  }
  for _, example := range examples {
    actual := Sum(example.Input)
    if actual != example.Expected {
      t.Fatalf("Sum(%v): Got %d, expected %d", example.Input, actual, example.Expected)
    }
  }
  s, _ := base64.StdEncoding.DecodeString("VGhlIG5leHQgc2VudGVuY2UgaXMgdHJ1ZS4gVGhlIHByZXZpb3VzIHNlbnRlbmNlIGlzIGZhbHNlLg==")
  t.Log(string(s))
}
