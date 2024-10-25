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
	sourceHandler, err := NewSourceHandler(cfg.SourceProtocol, cfg.SourceHost, cfg.SourcePort, cfg.SourceTopic)
	if err != nil {
		return nil, err
	}

	// Initialize destination handlers
	var destHandlers []DestinationHandler
	for _, dest := range cfg.Destinations {
		handler, err := NewDestinationHandler(dest.Protocol, dest.Host, dest.Port, dest.TopicOrQueue)
		if err != nil {
			return nil, err
		}
		destHandlers = append(destHandlers, handler)
	}

	return &Muxer{
		config:        cfg,
		sourceHandler: sourceHandler,
		destHandlers:  destHandlers,
		ctx:           ctx,
		cancel:        cancel,
	}, nil
}

func NewSourceHandler(protocol, host string, port int, topic string) (SourceHandler, error) {
	switch protocol {
	}

	return nil, fmt.Errorf("unsupported source protocol: %s", protocol)
}

func NewDestinationHandler(protocol, host string, port int, topicOrQueue string) (DestinationHandler, error) {
	switch protocol {
	}
	return nil, fmt.Errorf("unsupported destination protocol: %s", protocol)
}

// Start begins the multiplexing operation, forwarding messages from the source to each destination.
func (m *Muxer) Start() error {
	m.wg.Add(1)
	go m.sourceHandler.Listen(m.ctx, m.forwardMessage)

	// Wait until stop is called
	m.wg.Wait()
	return nil
}

// forwardMessage takes a message from the source and forwards it to each destination.
func (m *Muxer) forwardMessage(message []byte) {
	for _, destHandler := range m.destHandlers {
		destHandler.Send(message)
	}
}

// Stop gracefully shuts down the muxer and cleans up all handlers.
func (m *Muxer) Stop() error {
	m.cancel()  // Signal all goroutines to stop
	m.wg.Wait() // Wait for all goroutines to complete

	// Close source handler
	if err := m.sourceHandler.Close(); err != nil {
		return err
	}

	// Close each destination handler
	for _, handler := range m.destHandlers {
		if err := handler.Close(); err != nil {
			return err
		}
	}

	return nil
}
