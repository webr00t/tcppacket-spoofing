package wifi

import (
	"github.com/google/gopacket/pcap"
	"fmt"
	"github.com/DennisDenuto/wifi-redirector/interfaces"
)

type Device struct {
	Name string
}

func (device Device) GetInterface(detector interfaces.InterfaceTypeDetector, ifs []pcap.Interface) (pcap.Interface, error) {
	for key, value := range ifs {
		if detector.IsType(value.Name)&& len(value.Addresses) > 1 {
			fmt.Println(key)
			fmt.Println(value)
			fmt.Println("WIRELESS")
			return value, nil
		}
	}

	return pcap.Interface{}, fmt.Errorf("")
}
