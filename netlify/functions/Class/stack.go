package Class

import "fmt"

type Stack struct {
  // Stack that can contain any type of value
  items []string
}

func (s* Stack) IsEmpty() bool {
  return len(s.items) == 0
}

func (s* Stack) Size() int {
  return len(s.items)
}

// Function to see the top of the stack
func (s* Stack) Peek() (string, error) {
  if (s.IsEmpty()) {
    return "Invalid Syntax", fmt.Errorf("Invalid Syntax")
  }
  return s.items[len(s.items)-1], nil
}

func (s* Stack) Push(item string) {
  s.items = append(s.items, item)
}

func (s* Stack) Pop() (string, error) {
  if (s.IsEmpty()) {
    return "Invalid Syntax", fmt.Errorf("Invalid Syntax")
  }
  item := s.items[len(s.items)-1]
  // slicing the array
  s.items = s.items[:len(s.items)-1]
  return item, nil
}
