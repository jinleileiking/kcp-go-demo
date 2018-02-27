package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	kcp "github.com/xtaci/kcp-go"
)

func main() {
	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)

	if err != nil {
		spew.Dump(err)
	}

	log.Println("Server started")

	for {
		conn, err := lis.AcceptKCP()

		if err != nil {
			spew.Dump("AcceptKCP failed", err)
		}

		log.Println("remote address:", conn.RemoteAddr())
		conn.SetStreamMode(true)
		conn.SetWriteDelay(true)
		// conn.SetNoDelay(config.NoDelay, config.Interval, config.Resend, config.NoCongestion)
		// conn.SetMtu(config.MTU)
		conn.SetWindowSize(1000, 1000)
		// conn.SetACKNoDelay(config.AckNodelay)

		for {
			b := make([]byte, 11)
			num, err := conn.Read(b)

			if err != nil {
				spew.Dump(num, err)
			}

			if num != 0 {
				spew.Dump(b)
			}
		}
	}

}
