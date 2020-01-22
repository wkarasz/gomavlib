// +build ignore

package main

import (
	"fmt"

	"github.com/wkarasz/gomavlib"
	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

func main() {
	// create a node which
	// - communicates with a serial port
	// - understands ardupilotmega dialect
	// - writes messages with given system id
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointSerial{"/dev/ttyUSB0:57600"},
		},
		Dialect:     ardupilotmega.Dialect,
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	// gomavlib provides different kinds of event
	for evt := range node.Events() {
		switch ee := evt.(type) {
		case *gomavlib.EventFrame:
			fmt.Printf("frame received: %v\n", ee)

		case *gomavlib.EventParseError:
			fmt.Printf("parse error: %v\n", ee)

		case *gomavlib.EventChannelOpen:
			fmt.Printf("channel opened: %v\n", ee)

		case *gomavlib.EventChannelClose:
			fmt.Printf("channel closed: %v\n", ee)
		}
	}
}
