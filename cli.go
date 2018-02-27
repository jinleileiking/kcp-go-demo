package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	kcp "github.com/xtaci/kcp-go"
)

func main() {
	conn, err := kcp.DialWithOptions("127.0.0.1:10000", nil, 10, 3)

	if err != nil {
		spew.Dump(err)
		os.Exit(-1)
	}

	b := "hello world"
	num, err := conn.Write([]byte(b))

	if err != nil {
		spew.Dump(num, err)
	}
	spew.Dump("Client done", num)

}
