package core

import (
	"context"
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
func (s *TcpServer) Start(ctx context.Context) {
	fmt.Println("Server starting on port:", s.opts.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.opts.Port))
	if err != nil {
		fmt.Printf("Unable to start tcp server: %s", err)
		return
	}
	defer listener.Close()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped accepting new connections.")
			return
		default:
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
			go s.handleTcpConnection(ctx, conn)
		}
	}
}

func (s *TcpServer) handleTcpConnection(ctx context.Context, conn net.Conn) {
	// TODO: it would probably be a good idea to create a custom `log(...)` function
	// the right metadata about the relevant requests (e.g. remote address, time, etc.)
	fmt.Printf("@%s New TCP connection opened.\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("@%s Closing connection due to server shutdown.\n", conn.RemoteAddr())
			return
		default:
			buf := make([]byte, TcpBufferLength)
			n, err := conn.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Printf("@%s TCP connection closed.\n", conn.RemoteAddr())
				} else {
					fmt.Printf("@%s TCP listener error: %s\n", conn.RemoteAddr(), err)
				}
				return
			}

			if n < TcpBufferLength {
				fmt.Printf("@%s Received '%d' bytes: %s", conn.RemoteAddr(), n, buf[:n])
			} else {
				fmt.Printf("@%s Received '%d' bytes which exceed the limit of '%d' bytes. Extra bytes will be ignored: %s", conn.RemoteAddr(), n, TcpBufferLength, buf[:TcpBufferLength])
			}
		}
	}
}
