package main

import (
	"fmt"
	"os"
	"time"

	"github.com/webr00t/tcppacket-spoofing/listener/http"
	sender_http "github.com/webr00t/tcppacket-spoofing/sender/http"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

func main() {
	httpListener := http.NewHttpListener("en1")
	httpPackets, err := httpListener.Listen(http.HttpPacketReader{})

	if err != nil {
		fmt.Errorf("error listening", err)
		return
	}

	handle, _ := pcap.OpenLive("en1",
		int32(65535),
		true,
		-1*time.Second)

	httpInterceptor := sender_http.HttpInterceptor{Sender: sender_http.PacketSender{Handler: handle}}

	for packet := range httpPackets {
		httpInterceptor.Intercept("en1", packet, sender_http.PacketSender{})
	}

}

func WritePacketsToFile() *pcapgo.Writer {
	dumpFile, _ := os.Create("dump.pcap")
	defer dumpFile.Close()

	packetWrite := pcapgo.NewWriter(dumpFile)
	packetWrite.WriteFileHeader(
		65535,
		layers.LinkTypeEthernet,
	)

	return packetWrite
}
