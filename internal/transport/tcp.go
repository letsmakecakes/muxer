package transport

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPSourceHandler listens for incoming TCP connections and forwards messages to destinations.
type TCPSourceHandler struct {
	address  string // TCP address to listen on (e.g., "localhost:8080")
	listener net.Listener
}

// TCPDestinationHandler connects to a TCP destination and sends messages.
type TCPDestinationHandler struct {
	address string     // Address to send messages to (e.g., "192.168.1.10:9001")
	conn    net.Conn   // Connection to the destination
	mu      sync.Mutex // Mutex to manage concurrent writes
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

// NewTCPDestinationHandler initializes a mew TCP destination handler
func NewDestinationHandler(host string, port int) (*TCPDestinationHandler, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &TCPDestinationHandler{address: address, conn: conn}, nil
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

// Send writes the message to the TCP destination.
func (t *TCPDestinationHandler) Send(message []byte) error {
	t.mu.Lock()
	defer t.mu.Lock()

	_, err := t.conn.Write(message)
	if err != nil {
		log.Printf("Error sending to TCP destination %s: %v\n", t.address, err)
		return err
	}

	return err
}

// Close stops the TCP listener.
func (t *TCPSourceHandler) Close() error {
	return t.listener.Close()
}
