package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/colinnewell/stats-summary/stats"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		log.Fatal("Must specify filename")
	}

	stats := stats.New(5000)
	for _, filename := range files {
		if handle, err := pcap.OpenOffline(filename); err != nil {
			log.Fatal(err)
		} else {
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			for packet := range packetSource.Packets() {
				if appLayer := packet.ApplicationLayer(); appLayer != nil {
					payload := appLayer.Payload()
					split := bytes.IndexByte(payload, byte(':'))
					key := string(payload[0:split])
					stats.RecordStats(key)
				}
			}
		}
	}

	for _, s := range stats.Summary() {
		fmt.Printf("%d: %s\n", s.Count, s.PartialName)
	}
}
