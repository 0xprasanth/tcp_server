package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	listenAddress string
	ln            net.Listener
	quitch        chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddress: listenAddr,
		quitch:        make(chan struct{}),
	}

}

func (s *Server) Start() error {
	// ln, err => listen network, error
	ln, err := net.Listen("tcp", s.listenAddress)

	if err != nil {
		return err
	}
	// close the listner
	defer ln.Close() // cleans up after quitch channel connected

	s.ln = ln
	// start accept loop
	go s.acceptLoop()

	// listen for sys singal to stop server (like ctrl+c);
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// wait for signal or quit signal
	select {
	case <-signalChan:
		fmt.Println("Recevied interupt, stopping server gracefully...")
	case <-s.quitch:
		fmt.Println("Server stopping...")
	}

	// wait for the quitch channel
	// <-s.quitch

	return nil
}

func (s *Server) acceptLoop() {
	// inifinte loop
	for {
		// * underscore ( _ ) => acctual connection literal
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		// print if new connection and has no error
		fmt.Println("new connection to server from", conn.RemoteAddr(), " of type: ", conn.RemoteAddr().Network())

		// spin up new `go` routine for non-blocking operations
		go s.readLoop(conn)

	}
}

// each if someone connects,  we need to read from connection and somewrite it
func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)
	for {
		//
		endBuffer, err := conn.Read((buf))
		if err != nil {
			fmt.Println("read error: ", err)
			continue
		}
		msg := buf[:endBuffer]
		fmt.Println(string(msg))
	}
}

func main() {
	server := NewServer(":3000")
	// server.Start()
	// use log fatal for error logs
	log.Fatal(server.Start())
}
