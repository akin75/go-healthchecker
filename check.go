package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	adress := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", adress, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v ", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
