package listener

import (
	"github.com/google/gopacket"
)

type Listener interface {
	Listen(packetReader PacketReader) chan gopacket.Packet
}


type PacketReader interface {
	Packets(deviceName string) (chan gopacket.Packet, error)
}
