package srt

import "sort"

type InsertionSort struct {
}

func (i *InsertionSort) Sort(inp []int) error {
  sort.Ints(inp)
  return nil
}
