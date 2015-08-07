package tescat

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"strconv"
)

type Capture struct {
	uniques map[[20]byte]bool
}

func NewCapture() (c *Capture, err error) {
	c = new(Capture)

	c.uniques = make(map[[20]byte]bool)
	return
}

func (c *Capture) StartFromUDP(port int) (err error) {
	addr, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return err
	}
	for {
		buf := make([]byte, 1024)

		n, _, err := conn.ReadFromUDP(buf)
		if err == nil {
			m := NewRawMessage(buf[0:n])
			c.processDefault(m)
		} else {
			fmt.Println("Error: ", err)
		}
	}
}

func (c *Capture) StartFromPCAP(file string, port int) (err error) {

	handle, err := pcap.OpenOffline(file)
	if err != nil {
		log.Fatal("PCAP OpenOffline error:", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		udp, ok := udpLayer.(*layers.UDP)
		if ok && udp != nil {
			if udp.DstPort != layers.UDPPort(port) {
				m := NewRawMessage(udp.Payload)
				c.processDefault(m)
			}
		}

	}
	return
}

func (c *Capture) processDefault(m *Message) {
	var key [20]byte
	copy(key[0:20], m.Bytes())
	exists := c.uniques[key]
	if !exists {
		c.uniques[key] = true
		fmt.Println("UNIQ " + m.String())
	}
}
