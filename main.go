package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
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
	files := os.Args[1:]

	if len(files) == 0 {
		log.Fatal("Must specify filename")
	}

	counts = make(map[string]int64)
	for _, filename := range files {
		if handle, err := pcap.OpenOffline(filename); err != nil {
			log.Fatal(err)
		} else {
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			for packet := range packetSource.Packets() {
				handlePacket(packet)
			}
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
