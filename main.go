package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var counts map[string]int64

func handlePacket(p gopacket.Packet) {

	if appLayer := p.ApplicationLayer(); appLayer != nil {
		payload := appLayer.Payload()
		split := bytes.IndexByte(payload, byte(':'))
		key := string(payload[0:split])
		counts[key]++
	}
	return
}

type statistic struct {
	key   string
	count int64
}

func main() {
	counts = make(map[string]int64)
	if handle, err := pcap.OpenOffline("/tmp/monitor.pcap"); err != nil {
		log.Fatal(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet) // Do something with a packet here.
		}
	}
	stats := make([]statistic, len(counts))
	i := 0
	for k, v := range counts {
		stats[i].key = k
		stats[i].count = v
		i++
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].count < stats[j].count })
	for _, s := range stats {
		fmt.Printf("%d: %s\n", s.count, s.key)
	}
}
