package transport

import (
	"bufio"
	"context"
	"fmt"
	"log"
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

// Listen starts the TCP server, accepting connections and reading messages.
func (t *TCPSourceHandler) Listen(ctx context.Context, forwardFunc func([]byte)) error {
	defer t.listener.Close()

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Printf("TCP accept error: %v\n", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			scanner := bufio.NewScanner(c)
			for scanner.Scan() {
				select {
				case <-ctx.Done():
					return
				default:
					message := scanner.Bytes()
					forwardFunc(message) // Forward the received message to muxer
				}
			}
			if err := scanner.Err(); err != nil {
				log.Printf("Error reading from TCP connection: %v\n", err)
			}
		}(conn)
	}
}
