package http

import "github.com/google/gopacket"

type HttpListener struct {
	Interface string
}

func (listener HttpListener) Listen() chan gopacket.Packet {
	return nil
}

