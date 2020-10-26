package spec_test

import (
	"net/http"
	"os"

	"github.com/keitam0/agility/spec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/metrics", func() {
	Describe("GET", func() {
		Context("with an agility instance", func() {
			var (
				ins *os.Process
			)

			BeforeEach(func() {
				i, err := spec.StartAgility(map[string]string{})
				if err != nil {
					panic(err)
				}
				ins = i
			})

			AfterEach(func() {
				ins.Kill()
				ins.Wait()
			})

			It("successes", func() {
				res, err := http.Get("http://localhost:9090/metrics")
				if err != nil {
					Fail(err.Error())
				}
				defer res.Body.Close()
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})
