package http

import (
	"github.com/google/gopacket"
	"time"
	"github.com/google/gopacket/pcap"
	"fmt"
)

type HttpListener struct {
	Interface string
}

func (listener HttpListener) Listen() chan gopacket.Packet {
	_, err := pcap.OpenLive(listener.Interface,
		int32(65535),
		true,
		-1 * time.Second)

	fmt.Println(err)
	return nil
}

