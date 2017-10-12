package srt

import "sort"

type MergeSort struct {
}

func (i *MergeSort) Sort(inp []int) error {
  sort.Ints(inp)
  return nil
}
