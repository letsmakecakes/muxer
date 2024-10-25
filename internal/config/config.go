package config

type Destination struct {
	Protocol     string
	Host         string
	Port         int
	TopicOrQueue string
}

type Config struct {
	SourceProtocol string
	SourceHost     string
	SourcePort     int
	SourceTopic    string
	Destinations   []Destination
}
