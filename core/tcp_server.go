package core

import (
	"errors"
	"fmt"
	"io"
	"net"
)

// Number of bytes read from the TCP connection.
var TcpBufferLength = 2048

type TcpServerOps struct {
	Port int64
}

type TcpServer struct {
	opts TcpServerOps
}

func NewTcpServer(opts TcpServerOps) *TcpServer {
	return &TcpServer{
		opts: opts,
	}
}

// Start spins up the TCP server and starts listening for messages.
func (s *TcpServer) Start() {
	fmt.Println("Server starting on port:", s.opts.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.opts.Port))
	if err != nil {
		fmt.Printf("Unable to start tcp server: %s", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("TCP connection closed.")
			} else {
				fmt.Printf("TCP listener error: %s\n", err)
			}
			continue
		}

		// New connections are handled in a separate Go Routine to avoid blocking
		// the main thread. The main thread should be in charge of listening for
		// all "new" incoming connections.
		go s.handleTcpConnection(conn)
	}
}

func (s *TcpServer) handleTcpConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New TCP connection opened!")

	for {
		buf := make([]byte, TcpBufferLength)
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("TCP connection closed.")
			} else {
				fmt.Printf("TCP listener error: %s\n", err)
			}
			return
		}

		if n < TcpBufferLength {
			fmt.Printf("Received '%d' bytes: %s\n", n, buf[:n])
		} else {
			fmt.Printf("Received '%d' bytes which exceed the limit of '%d' bytes. Extra bytes will be ignored: %s", n, TcpBufferLength, buf[:TcpBufferLength])
		}
	}
}
