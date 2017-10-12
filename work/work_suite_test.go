package work_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Work Suite")
}
