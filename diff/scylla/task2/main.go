// Throttling bandwidth with Go
// Imagine a service that is serving big log files over raw TCP, you decided to implement simple QoS for an existing server.

// The goal of the task is to create small Go package that would allow for throttling bandwidth for TCP connections.

// Requirements:

// keep it simple
// think of it as open-source software that builds upon existing solutions
// use only stdlib and/or supplementary repositories (golang.org/x/*)
// prefer reusing existing packages
// e.g. https://pkg.go.dev/golang.org/x/time/rate?tab=doc
// the package should allow for:
// setting one cumulative bandwidth limit for all connections* (aka global limit)
// setting single individual bandwidth limit for all connections** (aka connection limit)
// changing limits in runtime (applies to all existing connections)
// for a 30s transfer sample consumed bandwidth should be accurate +/- 5%
// Reasonable deadline for submitting a solution is 8 days. Questions, clarifications - please drop us an e-mail!

// * example: limit x kB/s means all connections can cumulatively consume at most x kB in a given second

// ** example: limit y kB/s means each and every connection can consume at most y kB independently from each other in a given second
package main

import (
	"fmt"
	"net"
	"time"

	"github.com/olexandrPodustov/diff/scylla/task2/bandwidththrottle"
	"golang.org/x/time/rate"
)

func main() {
	// Create a BandwidthThrottle with a global rate of 1024 bytes per second.
	bt := bandwidththrottle.NewBandwidthThrottle(rate.Limit(1024))

	// Create a TCP server.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed starting the listener:", err)
		return
	}
	defer listener.Close()

	// Accept and handle incoming connections.
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed accepting the connection:", err)
			continue
		}

		// Set a connection-specific bandwidth rate.
		bt.SetConnectionRate(conn, rate.Limit(512)) // 512 bytes per second

		// Wrap the connection with bandwidth throttling.
		throttledConn := bt.WrapConnection(conn)

		// Handle the connection in a goroutine.
		go handleConnection(throttledConn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("failed reading the data:", string(data), err)
			return
		}

		// Simulate data processing.
		processData(data[:n])
	}
}

func processData(data []byte) {
	// Simulate data processing.
	fmt.Printf("Received %d bytes\n", len(data))

	// Introduce some delay to see throttling in action.
	time.Sleep(time.Millisecond * 100)
}
