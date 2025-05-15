package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Response struct {
	Type       string      `json:"type,omitempty"`
	Host       string      `json:"host,omitempty"`
	PodName    string      `json:"podName,omitempty"`
	ServerPort string      `json:"serverPort,omitempty"`
	Path       string      `json:"path,omitempty"`
	Method     string      `json:"method,omitempty"`
	Headers    http.Header `json:"headers,omitempty"`
	Body       string      `json:"body,omitempty"`
}

type HttpServerHandler struct {
	port string
}

func (h HttpServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := &Response{
		Type:       "http",
		PodName:    os.Getenv("POD_NAME"),
		Host:       r.Host,
		ServerPort: h.port,
		Path:       r.URL.Path,
		Method:     r.Method,
		Headers:    r.Header,
	}
	fmt.Println("Request on url", r.URL.Path)
	json.NewEncoder(w).Encode(resp)
}

func RunHTTPServerOnPort(port string) {
	fmt.Println("http server running")
	http.ListenAndServe(port, HttpServerHandler{port})
}

type TCPServerHandler struct {
	port string
}

func (h TCPServerHandler) ServeTCP(con net.Conn) {
	fmt.Println("request on", con.LocalAddr().String())
	resp := &Response{
		Type:       "tcp",
		Host:       con.LocalAddr().String(),
		PodName:    os.Getenv("POD_NAME"),
		ServerPort: h.port,
	}
	json.NewEncoder(con).Encode(resp)
}

func RunTCPServerOnPort(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tcp server listening")
	for {
		con, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		h := TCPServerHandler{port}
		go h.ServeTCP(con)
	}
}

type UDPServerHandler struct {
	port string
}

func (h UDPServerHandler) ServeUDP(con *net.UDPConn, addr *net.UDPAddr) {
	fmt.Printf("request on %s from %s\n", con.LocalAddr().String(), addr.String())
	resp := &Response{
		Type:       "udp",
		Host:       con.LocalAddr().String(),
		PodName:    os.Getenv("POD_NAME"),
		ServerPort: h.port,
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error marshaling response:", err)
		return
	}

	_, err = con.WriteToUDP(respBytes, addr)
	if err != nil {
		fmt.Println("Error sending UDP response:", err)
	}
}

func RunUDPServerOnPort(port string) {
	// Convert port string to int
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println("Failed to resolve UDP addr:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("udp server listening on port %s\n", port)
	for {
		buffer := make([]byte, 1024)
		_, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		h := UDPServerHandler{port}
		go h.ServeUDP(conn, addr)
	}
}

func main() {
	go RunHTTPServerOnPort(":8080")
	go RunHTTPServerOnPort(":8989")
	go RunHTTPServerOnPort(":9090")

	go RunTCPServerOnPort(":4343")
	go RunTCPServerOnPort(":4545")
	go RunTCPServerOnPort(":5656")

	go RunUDPServerOnPort(":7070")
	go RunUDPServerOnPort(":7272")
	go RunUDPServerOnPort(":7474")

	hold()
}

func hold() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)
}
