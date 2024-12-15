package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
)

// Number of bytes read from the TCP connection.
var TcpBufferLength = 2048 // ~2KB

type TcpServerOps struct {
	Port int64
}

type TcpServer struct {
	opts TcpServerOps
	rp   RequestProcessor
}

func NewTcpServer(opts TcpServerOps, rp RequestProcessor) *TcpServer {
	return &TcpServer{
		opts: opts,
		rp:   rp,
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

	buf := make([]byte, TcpBufferLength)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("@%s Closing connection due to server shutdown.\n", conn.RemoteAddr())
			return
		default:
			n, err := conn.Read(buf) // Will never exceed the size of the input buffer
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Printf("@%s TCP connection closed.\n", conn.RemoteAddr())
				} else {
					fmt.Printf("@%s TCP listener error: %s\n", conn.RemoteAddr(), err)
				}
				return
			}

			// n will never be greater than TcpBufferLength, but it could be less than
			// TcpBufferLength if fewer bytes are available. We use the actual number of
			// bytes read (n) to process the data.
			data := string(buf[:n])
			fmt.Printf("@%s Received '%d' bytes: %s\n", conn.RemoteAddr(), n, data)

			// The last character is a newline, so we remove it
			if len(data) > 0 && data[len(data)-1] == '\n' {
				data = data[:len(data)-1]
			}

			req, err := ParseProtocol(data)
			if err != nil {
				fmt.Printf("@%s Error parsing protocol: %s\n", conn.RemoteAddr(), err)
				s.writeError(err, conn)
				continue
			}

			res, err := s.rp.Process(ctx, req)
			if err != nil {
				fmt.Printf("@%s Error processing request: %s\n", conn.RemoteAddr(), err)
				s.writeError(err, conn)
				continue
			}

			s.writeSuccess(res.Value, conn)
		}
	}
}

func (s *TcpServer) writeError(err error, conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("ERROR: %s\n", err)))
}

func (s *TcpServer) writeSuccess(value []byte, conn net.Conn) {
	if len(value) == 0 {
		conn.Write([]byte("OK\n"))
		return
	}
	conn.Write([]byte(fmt.Sprintf("OK: %s\n", value)))
}
