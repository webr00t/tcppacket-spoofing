package http

import (
	"github.com/google/gopacket"
	"github.com/webr00t/tcppacket-spoofing/sender"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/layers"
	"fmt"
	"net"
	"time"
	"strconv"
)

type HttpInterceptor struct {
	Payload string
	Sender  PacketSender
}

type PacketSender struct {
	Handler *pcap.Handle
}

func (packetSender PacketSender) Send(buffer gopacket.SerializeBuffer) error {
	if packetSender.Handler == nil {
		handle, _ := pcap.OpenLive("en1",
			int32(65535),
			true,
			-1 * time.Second)
		packetSender.Handler = handle
	}
	//defer handle.Close()
	//fmt.Println(err)
	return packetSender.Handler.WritePacketData(buffer.Bytes())
}

func (httpInterceptor HttpInterceptor) Intercept(deviceName string, packet gopacket.Packet, packetSender sender.PacketSender) error {

	buffer := gopacket.NewSerializeBuffer()

	//options := gopacket.SerializeOptions{}
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}


	fmt.Println(packet.NetworkLayer().NetworkFlow())
	fmt.Println(packet.TransportLayer().TransportFlow())

	srcIp := packet.NetworkLayer().NetworkFlow().Src().Raw()
	destIp := packet.NetworkLayer().NetworkFlow().Dst().Raw()

	srcPort, _ := strconv.Atoi(packet.TransportLayer().TransportFlow().Src().String())
	destPort, _ := strconv.Atoi(packet.TransportLayer().TransportFlow().Dst().String())

	//srcPort := binary.LittleEndian.Uint16(packet.TransportLayer().TransportFlow().Src().Raw())
	//destPort := binary.LittleEndian.Uint16(packet.TransportLayer().TransportFlow().Dst().Raw())

	var seq uint32
	var ack uint32
	var sourceMacAddress net.HardwareAddr
	var destMacAddress net.HardwareAddr
	var tcpSegmengLength int

	if ethernetLayer := packet.Layer(layers.LayerTypeEthernet); ethernetLayer != nil {
		ether, _ := ethernetLayer.(*layers.Ethernet)
		sourceMacAddress = ether.SrcMAC
		destMacAddress = ether.SrcMAC
	}

	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		// Get actual TCP data from this layer
		tcp, _ := tcpLayer.(*layers.TCP)
		seq = tcp.Seq
		ack = tcp.Ack
		tcpSegmengLength = len(tcp.Payload)
	}

	/*
	fmt.Println("-=======")
	fmt.Println("sequence ", seq)
	fmt.Println("ack ", ack + 1)
	fmt.Println(sourceMacAddress)
	fmt.Println(destMacAddress)

	fmt.Println("source port string", srcPort)
	fmt.Println("source port", srcPort)
	fmt.Println("dest port string", packet.TransportLayer().TransportFlow().Dst().String())
	fmt.Println("dest port", layers.TCPPort(destPort))

	fmt.Println("source ip", srcIp)
	fmt.Println("dest ip", destIp)
	fmt.Println("-=======")
	*/

	var httpResponseString string = "HTTP/1.1 200 OK\r\n" +
	"Content-Type: text/plain; charset=utf-8\r\n" +
	"Date: Fri, 28 Oct 2016 02:37:46 GMT\r\n" +
	"Content-Length: 4\r\n" +
	"Connection: keep-alive\r\n" +
	"\r\n" +
	"ABC\n"


	ipLayer := &layers.IPv4{
		Version: 4,
		IHL: 5,
		TTL: 44,
		Flags: layers.IPv4DontFragment,
		Protocol: layers.IPProtocolTCP,
		SrcIP: net.IP(destIp),
		DstIP: net.IP(srcIp),
	}

	tcpLayer := &layers.TCP{
		//DataOffset: 25,
		DataOffset: 21,
		Window: 74,
		Seq: ack,
		Ack: seq + uint32(tcpSegmengLength),
		SYN: false,
		ACK: true,
		PSH: true,
		SrcPort: layers.TCPPort(destPort),
		DstPort: layers.TCPPort(srcPort),
		Options: []layers.TCPOption{
			layers.TCPOption{
				OptionType: layers.TCPOptionKindNop,
				OptionLength: 1,
			},
			layers.TCPOption{
				OptionType: layers.TCPOptionKindNop,
				OptionLength: 1,
			},
			//layers.TCPOption{
			//	OptionType: layers.TCPOptionKindTimestamps,
			//	OptionLength: 10,
			//	OptionData: []byte{0x0, 0x0, 0x0, 0xf, 0x8, 0xe, 0xe, 0x3, 0xb, 0x9},
			//},
		},
	}

	tcpLayer.SetNetworkLayerForChecksum(ipLayer)

	gopacket.SerializeLayers(buffer, options,
		&layers.Ethernet{EthernetType: layers.EthernetTypeIPv4, DstMAC: sourceMacAddress, SrcMAC: destMacAddress},
		ipLayer,
		tcpLayer,
		//gopacket.Payload([]byte{65, 66, 67}),
		gopacket.Payload([]byte(httpResponseString)),
	)

	err := httpInterceptor.Sender.Send(buffer)
	fmt.Println(err)

	return nil
}

/*
0000   7c d1 c3 73 e8 44 a4 2b 8c 01 c5 85 08 00 45 20  |..s.D.+......E
0010   00 34 33 7a 40 00 39 06 16 a0 5d b8 d8 22 c0 a8  .43z@.9...].."..
0020   01 07 00 50 fa 62 af 46 2a c1 b1 6c d8 fd 80 10  ...P.b.F*..l....
0030   01 1b 57 b1 00 00 01 01 08 0a c1 a0 19 bc 2b c9  ..W...........+.
0040   c0 1b                                            ..
*/
