package work_test

import (
  . "github.com/onsi/gomega"
  . "github.com/onsi/ginkgo"
  "github.com/devd/work"
  "github.com/devd/srt/fakesorters"
  "sort"
  "errors"
)

var _ = Describe("work", func(){
  Context("ListTransformer call",func(){
    var fs *fakesorters.FakeSorter
    BeforeEach(func(){
      fs = &fakesorters.FakeSorter{}
    })
    It("returns error if unable to sort", func(){
      fs.SortCall.Returns.AnError = errors.New("")
      _,err:=work.ListTransformer([]int{1,2,3}, fs)
      Expect(fs.SortCall.Receives.IntArr).To(Equal([]int{1,2,3}))
      Expect(fs.SortCall.Called).To(BeTrue())
      Expect(err).To(HaveOccurred())
    })

    It("returns sorted array if no error", func(){
      fs.SortCall.Returns.AnError = nil
      fs.SortCall.DoSort = true
      arr,err:=work.ListTransformer([]int{3,2,3}, fs)
      Expect(fs.SortCall.Called).To(BeTrue())
      Expect(err).NotTo(HaveOccurred())
      Expect(sort.IntsAreSorted(arr)).To(BeTrue())
      Expect(fs.SortCall.Receives.Orig).To(Equal([]int{3,2,3}))
    })
  })
})
