package utils

import (
	"fmt"
	"muxer/internal/config"
	"strings"
)

func ParseDestinations(destinations string) ([]config.Destination, error) {
	var destConfigs []config.Destination
	destList := strings.Split(destinations, ",")
	for _, dest := range destList {
		parts := strings.Split(dest, ":")
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid destination format, expected <protocol>:<host>:<port>")
		}
		destConfigs = append(destConfigs, config.Destination{
			Protocol: parts[0],
			Host:     parts[1],
			Port:     parts[2],
		})
	}
	return destConfigs, nil
}
