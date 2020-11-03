package spec_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	GinDefaultWriter = gin.DefaultWriter
)

func TestSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Suite")
}

var _ = BeforeSuite(func() {
	gin.DefaultWriter = GinkgoWriter
})

var _ = AfterSuite(func() {
	gin.DefaultWriter = GinDefaultWriter
})
