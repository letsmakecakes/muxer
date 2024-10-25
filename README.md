# Muxer

**Muxer** is a command-line tool built in Golang for receiving requests and multiplexing them to multiple predefined destinations. It supports a variety of protocols, including **TCP, UDP, Kafka, RabbitMQ, and Redis**, with plans for future protocol support.

## Features
- Accepts requests over different protocols and forwards them to multiple destinations.
- Configurable via command-line arguments, supporting different source and destination protocols.
- Supports graceful shutdown to manage open connections securely.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Supported Protocols](#supported-protocols)
- [Command-Line Arguments](#command-line-arguments)
- [Project Structure](#project-structure)
- [Contributing](#contributing)

## Installation

### Prerequisites
- [Go 1.20+](https://golang.org/dl/)

### Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/muxer.git
    cd muxer
    ```

2. Build the project:
    ```bash
    go build -o muxer ./cmd/muxer
    ```

3. Run the `muxer` executable:
    ```bash
    ./muxer --help
    ```

## Usage

To start the `Muxer` tool, specify the source protocol, source host and port, and the list of destinations.

### Basic Example
Forward TCP traffic from port 8080 to two destinations, one over UDP and another over TCP.

```bash
./muxer --source-protocol tcp --source-host 0.0.0.0 --source-port 8080 --destinations udp:192.168.1.10:9001,tcp:192.168.1.11:9002
```

### Example with Kafka as Source
Forward messages from a Kafka topic to multiple TCP and UDP destinations:

```bash
./muxer --source-protocol kafka --source-host 192.168.1.40 --source-port 9092 --source-topic source_topic --destinations tcp:192.168.1.12:9003,udp:192.168.1.13:9004
```

## Supported Protocols

The following protocols are currently supported for sources and destinations:
- **TCP**
- **UDP**
- **Kafka**
- **RabbitMQ**
- **Redis**

Each protocol has specific configurations. For example, Kafka requires a topic, and RabbitMQ requires a queue name.

## Command-Line Arguments

| Argument           | Description                                                                                |
|--------------------|--------------------------------------------------------------------------------------------|
| `--source-protocol` | Protocol for the source (e.g., `tcp`, `udp`, `kafka`, `rabbitmq`, `redis`).               |
| `--source-host`     | Source host address (default: `0.0.0.0`).                                                 |
| `--source-port`     | Port to listen on for the source protocol.                                                |
| `--source-topic`    | Topic name (for Kafka) or queue name (for RabbitMQ).                                      |
| `--destinations`    | Comma-separated list of destinations in `<protocol>:<host>:<port>` format.                |

### Destination Format
Each destination should follow the `<protocol>:<host>:<port>` format. For protocols like Kafka or RabbitMQ, add the topic or queue name as an additional field.

### Example
Forward from a Redis source to two Kafka topics and one UDP endpoint:

```bash
./muxer --source-protocol redis --source-host 192.168.1.50 --source-port 6379 --source-key source_key --destinations kafka:192.168.1.40:9092:topic1,udp:192.168.1.20:9003
```

## Contributing

Contributions are welcome! To get started:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Test thoroughly.
5. Submit a pull request.

Please make sure your code adheres to the Go coding standards.

## License
This project is licensed under the MIT License.
