package http

import (
	"github.com/google/gopacket"
	"time"
	"github.com/google/gopacket/pcap"
	"fmt"
	"github.com/DennisDenuto/wifi-redirector/listener"
	"strings"
)

type HttpListener struct {
	DeviceName string
}

type HttpPacketReader struct {

}

func (httpPacketReader HttpPacketReader) Packets(deviceName string) (chan gopacket.Packet, error) {
	handle, err := pcap.OpenLive(deviceName,
		int32(65535),
		true,
		-1 * time.Second)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	handle.SetBPFFilter("tcp and (port 80)")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	return packetSource.Packets(), nil
}

func (listener HttpListener) Listen(packetReader listener.PacketReader) (chan gopacket.Packet, error) {
	httpPackets := make(chan gopacket.Packet)
	tcpPackets, err := packetReader.Packets(listener.DeviceName)

	if err != nil {
		return nil, err
	}

	go func() {
		for tcpPacket := range tcpPackets {
			if tcpPacket.ApplicationLayer() == nil {
				continue
			}

			if strings.Contains(string(tcpPacket.ApplicationLayer().Payload()), "HTTP/1.1") {
				httpPackets <- tcpPacket
			}
		}
	}()

	return httpPackets, nil
}

