package work

import "github.com/devd/srt"

func ListTransformer(inp []int, sorter srt.Sorter) ([]int, error){
  newArr := make([]int,len(inp))
  copy(newArr[:], inp)
  err := sorter.Sort(newArr)
  return newArr, err
}
