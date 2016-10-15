package http_test

import (
	. "github.com/DennisDenuto/wifi-redirector/listener/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Listener", func() {

	Context("Listen to all http traffic", func() {
		It("Should listen to only http traffic on an interface", func() {
			listener := HttpListener{}
			packets := listener.Listen()

			Expect(packets).ToNot(BeNil())
		})
	})
})
