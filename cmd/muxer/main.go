package muxer

import (
	"flag"
	"fmt"
	"log"
	"muxer/internal/config"
	"muxer/pkg/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sourceProtocol := flag.String("source-protocol", "", "Source protocol: tcp, udp, kafka, rabbitmq, redis")
	sourceHost := flag.String("source-host", "0.0.0.0", "Source host")
	sourcePort := flag.Int("source-port", 0, "Source port")
	sourceTopicOrQueue := flag.String("source-topic", "", "Source topic/queue (for Kafka/RabbitMQ only)")
	destinations := flag.String("destinations", "", "Comma-separated list of destinations in <protocol>:<host>:<port> format")

	flag.Parse()

	if *sourceProtocol == "" || *sourcePort == 0 || *destinations == "" {
		log.Fatalf("source-protocol, source-port, and destinations are required fields")
	}

	destConfigs, err := utils.ParseDestinations(*destinations)
	if err != nil {
		log.Fatal("Error parsing destinations: %v", err)
	}

	config := config.Config{
		SourceProtocol: *sourceProtocol,
		SourceHost:     *sourceHost,
		SourcePort:     *sourcePort,
		SourceTopic:    *sourceTopicOrQueue,
		Destinations:   destConfigs,
	}

	// TODO: Initialize muxer with source and destination configurations

	// TODO: Start multiplexing

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("\nShutting down muxer...")

	// TODO: Graceful shutdowm muxer

	fmt.Println("Muxer stopped gracefully")
}
