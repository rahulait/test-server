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

To build docker image and push to dockerhub, use:
```bash
DOCKER_USER=<dockerhub-username> make release
```

### Usage

```bash
./test-server
```

One can also use pre-built docker image.
```bash
docker run -it rahulait/test-server:latest
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

### Connecting to TCP Port
```
echo 1 | nc -w1 <server-ip> <tcp-port>

# Example:
echo 1 | nc -w1 172.232.174.244 4343
{"type":"tcp","host":"172.232.174.244:4343","serverPort":":4343"}
```

### Connecting to HTTP Port
```
curl <server-ip>:<http-port>

# Example:
curl 172.232.174.244:80
{"type":"http","host":"172.232.174.244","serverPort":":80","path":"/","method":"GET","headers":{"Accept":["*/*"],"User-Agent":["curl/8.5.0"]}}
```

### Connecting to UDP Port
```
echo 1 | nc -u -w1 <server-ip> <udp-port>

# Example:
echo 1 | nc -u -w1 172.232.174.244 7070
{"type":"udp","host":"[::]:7070","serverPort":":7070"}
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests.

## License

This project is licensed under the Apache License 2.0
