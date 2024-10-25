package transport

import (
	"fmt"
	"net"
)

// TCPSourceHandler listens for incoming TCP connections and forwards messages to destinations.
type TCPSourceHandler struct {
	address  string // TCP address to listen on (e.g., "localhost:8080")
	listener net.Listener
}

// NewTCPSourceHandler initializes a new TCP source handler.
func NewTCPSourceHandler(host string, port int) (*TCPSourceHandler, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &TCPSourceHandler{address: address, listener: listener}, nil
}
