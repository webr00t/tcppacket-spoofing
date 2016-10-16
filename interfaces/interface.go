package interfaces

import "github.com/google/gopacket/pcap"

type ListInterfaces interface {
	List() []pcap.Interface
}

type InterfaceFinder interface {
	GetNextFreeInterface(detector InterfaceTypeDetector, listInterfaces ListInterfaces) (pcap.Interface, error)
	GetInterfaceByName(listInterfaces ListInterfaces, deviceName string) (pcap.Interface, error)
}