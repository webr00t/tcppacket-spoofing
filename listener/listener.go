package listener

import "github.com/google/gopacket"

type Listener interface {
	Listen() chan gopacket.Packet
}