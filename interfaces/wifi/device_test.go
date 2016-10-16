package wifi_test

import (
	. "github.com/DennisDenuto/wifi-redirector/interfaces/wifi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/google/gopacket/pcap"
)

type FakeWirelessTypeDetector struct {
	WifiDeviceName string
}

func (test FakeWirelessTypeDetector) IsType(deviceName string) bool {
	if test.WifiDeviceName == deviceName {
		return true
	}
	return false
}

type FakeListInterfaces struct {
	Interfaces []pcap.Interface
}

func (test FakeListInterfaces) List() []pcap.Interface {
	return test.Interfaces
}

var _ = Describe("Device", func() {

	Context("a machine with a wifi network card", func() {
		var ifs []pcap.Interface
		BeforeEach(func() {
			ifs = []pcap.Interface{
				pcap.Interface{
					Name: "dev-1",
					Description: "desc",
					Addresses: []pcap.InterfaceAddress{pcap.InterfaceAddress{}, pcap.InterfaceAddress{}},
				},
			}
		})

		It("Should return back the wifi device", func() {
			fakeWirelessTypeDetector := FakeWirelessTypeDetector{WifiDeviceName: "dev-1"}
			wifiInterface := Device{}

			wifiDevice, err := wifiInterface.GetInterface(fakeWirelessTypeDetector, FakeListInterfaces{Interfaces: ifs})
			Expect(err).To(BeNil())
			Expect(wifiDevice).ToNot(BeNil())
			Expect(wifiDevice.Name).To(Equal("dev-1"))
		})
	})

	Context("a machine without a wifi network card", func() {
		var ifs []pcap.Interface
		BeforeEach(func() {
			ifs = []pcap.Interface{
				pcap.Interface{
					Name: "not-wireless",
					Description: "desc",
					Addresses: []pcap.InterfaceAddress{pcap.InterfaceAddress{}, pcap.InterfaceAddress{}},
				},
			}
		})

		It("Should return back nil", func() {
			fakeWirelessTypeDetector := FakeWirelessTypeDetector{WifiDeviceName: "dev-1"}
			wifiInterface := Device{}

			_, err := wifiInterface.GetInterface(fakeWirelessTypeDetector, FakeListInterfaces{Interfaces: ifs})
			Expect(err).ToNot(BeNil())
		})
	})
})
