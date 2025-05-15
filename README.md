# Test Server

A simple server project for testing connections using different protocols. It runs a server listening on different ports using HTTP, TCP and UDP protocols.

## Features

- Lightweight and easy to set up
- Available in docker image
- Suitable for local development and testing
- Examples provided for running it in kubernetes cluster
- When running within k8s and proper env vars set, it returns name of pod which answered the request

## Getting Started

### Installation

```bash
git clone https://github.com/rahulait/test-server.git
cd test-server
make build
```

### Usage

```bash
./test-server
```

The server will start listening on following ports:
| Port | Protocol |
|------|----------|
| 80   | HTTP     |
| 8080 | HTTP     |
| 8989 | HTTP     |
| 9090 | HTTP     |
| 4343 | TCP      |
| 4545 | TCP      |
| 5656 | TCP      |
| 7070 | UDP      |
| 7272 | UDP      |
| 7474 | UDP      |

## Contributing

Contributions are welcome! Please open issues or submit pull requests.

## License

This project is licensed under the Apache License 2.0
