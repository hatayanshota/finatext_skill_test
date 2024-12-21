package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

const fileName = "mysql.pcap"

func main() {
	handle, err := pcap.OpenOffline(fileName)
	if err != nil {
		log.Fatalf("failed to open pcap file: %v", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// ペイロード（Application Layer）を取得
		if appLayer := packet.ApplicationLayer(); appLayer != nil {
			payload := string(appLayer.Payload())

			if len(payload) > 0 {
				log.Printf("payload: %v", payload)
			}
		}
	}
}
