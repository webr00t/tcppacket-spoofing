package interfaces

import "github.com/google/gopacket/pcap"

type Device interface {
	GetDevice(detector InterfaceTypeDetector, ifs []pcap.Interface) (pcap.Interface, error)
}