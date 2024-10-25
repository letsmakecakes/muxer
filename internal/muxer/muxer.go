package muxer

import (
	"context"
	"fmt"
	"muxer/internal/config"
	"sync"
)

// Muxer struct represents the core multiplexing service.
type Muxer struct {
	config        config.Config        // Config contains source and destinations configuration.
	sourceHandler SourceHandler        // Source handler based on the specified protocol.
	destHandlers  []DestinationHandler // Destination handlers for each protocol.
	ctx           context.Context      // Context to manage cancellation and shutdown.
	cancel        context.CancelFunc   // Cancel function to gracefully stop the muxer.
	wg            sync.WaitGroup       // WaitGroup to manage goroutines.
}

// SourceHandler defines methods for receiving messages from a source.
type SourceHandler interface {
	Listen(ctx context.Context, forwardFunc func([]byte)) error
	Close() error
}

// DestinationHandler defines methods for sending messages to a destination.
type DestinationHandler interface {
	Send(message []byte) error
	Close() error
}

// NewMuxer initializes and returns a new Muzer instance.
func NewMuxer(cfg config.Config) (*Muxer, error) {
	ctx, cancel := context.WithCancel(context.Background())

	// Initialize source handler based on protocol
	sourceHandler, err := NewSourceHandler(cfg.)
}

func NewSourceHandler(protocol, host string, port int, topic string) (SourceHandler, error) {
	switch protocol {
	}

	return nil, fmt.Errorf("unsupported source protocol: %s", protocol)
}