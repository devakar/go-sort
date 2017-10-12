package main

import (
	"fmt"

	"github.com/devd/srt"
	"github.com/devd/work"
)

func main() {
	inp := []int{7, 3, 7, 0, 1}
	sorter1 := &srt.InsertionSort{}
	arr, err := work.ListTransformer(inp, sorter1)
	fmt.Println(arr, err)

	inp = []int{7, 3, 0, 1}
	sorter2 := &srt.MergeSort{}
	arr, err = work.ListTransformer(inp, sorter2)
	fmt.Println(arr, err)
}
