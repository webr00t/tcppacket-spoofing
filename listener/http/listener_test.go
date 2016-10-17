package http_test

import (
	. "github.com/DennisDenuto/wifi-redirector/listener/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/google/gopacket"
	"fmt"
)

type FakePacketReader struct {
}

func (httpPacketReader FakePacketReader) Packets(deviceName string) (chan gopacket.Packet, error) {
	packets := make(chan gopacket.Packet, 1)

	packets <- nil
	return packets, nil
}

var _ = Describe("Listener", func() {

	Context("Listen to all http traffic", func() {

		BeforeEach(func() {

		})

		It("Should listen to only http traffic on an interface", func() {
			listener := HttpListener{DeviceName: "en1"}
			packets, err := listener.Listen(HttpPacketReader{})

			Expect(err).ToNot(HaveOccurred())

			//packet := <-packets

			for value := range packets {
				fmt.Println(string(value.ApplicationLayer().Payload()))
				fmt.Println(string(value.ApplicationLayer().LayerPayload()))
			}
		})
	})
})
