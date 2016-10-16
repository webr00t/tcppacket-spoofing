package wifi

import (
	"github.com/google/gopacket/pcap"
	"fmt"
	"github.com/DennisDenuto/wifi-redirector/interfaces"
)

type Device struct {
	Name string
}

func (device Device) GetInterface(detector interfaces.InterfaceTypeDetector, listInterfaces interfaces.ListInterfaces) (pcap.Interface, error) {
	for _, value := range listInterfaces.List() {
		if detector.IsType(value.Name)&& len(value.Addresses) > 1 {
			return value, nil
		}
	}

	return pcap.Interface{}, fmt.Errorf("")
}

func (device Device) GetDeviceByName(listInterfaces interfaces.ListInterfaces, deviceName string) (pcap.Interface, error) {
	return pcap.Interface{}, fmt.Errorf("")
}
