package fakesorters

import "sort"

type FakeSorter struct {
  SortCall struct {
    Called bool
    Receives struct {
      Orig []int
      IntArr []int
    }
    Returns struct {
      AnError error
    }
    DoSort bool
  }
}

func (fs *FakeSorter) Sort(inp []int) error {
  fs.SortCall.Called = true
  fs.SortCall.Receives.Orig= make([]int, len(inp))
  copy(fs.SortCall.Receives.Orig[:], inp)

  fs.SortCall.Receives.IntArr = inp

  if fs.SortCall.DoSort {
    sort.Ints(fs.SortCall.Receives.IntArr)
  }
  return fs.SortCall.Returns.AnError
}


