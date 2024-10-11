package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddress string
	ln            net.Listener
	quitch        chan struct{}
	msgch         chan Message // new channel only for msg purpose
	// create peerMap of map[net.Addr] to maintain, identify and track message from connections^_
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddress: listenAddr,
		quitch:        make(chan struct{}),
		// byte is always better than any data type or literal
		msgch: make(chan Message),
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
	//clean up
	close(s.msgch) // whenever we stop the server, people can still read from this channel

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
		// instead of print the buffer
		// write buffer into the channel, so that everyone can even if the buffer is full
		s.msgch <- Message{
			payload: buf[:endBuffer],
			from:    conn.RemoteAddr().String(),
		}
		// write back to the connection
		conn.Write([]byte("We got your message, Thank you."))
	}
}

func main() {
	server := NewServer(":3000")

	// go routing for readLoop
	go func() {
		for msg := range server.msgch {
			fmt.Printf("message recieved from (%s): %s\n", msg.from, string(msg.payload))
		}
	}()

	// use log fatal for error logs
	log.Fatal(server.Start())
}
