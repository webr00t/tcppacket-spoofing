package sender

import "github.com/google/gopacket"

type PacketSender interface {
	Send(buffer gopacket.SerializeBuffer) error
}

type Interceptor interface {
	Intercept(deviceName string, packet gopacket.Packet, packetSender PacketSender) error
}