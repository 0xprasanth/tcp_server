# TCP Server :gear: <!-- omit in toc -->

This project demonstrates a simple custom **TCP server** built in Go, using only the standard library. The server can handle incoming connections, read data sent by clients, and respond appropriately.

**Table of Contents**
- [Features :sparkles:](#features-sparkles)
- [How it Works :hammer\_and\_wrench:](#how-it-works-hammer_and_wrench)
- [TODO's :memo:](#todos-memo)
- [Getting Started :rocket:](#getting-started-rocket)
  - [Prerequisites :memo:](#prerequisites-memo)
  - [Running the Server :arrow\_forward:](#running-the-server-arrow_forward)
  - [Connecting to the Server 💻](#connecting-to-the-server-)
  - [Stopping the Server ⏹️](#stopping-the-server-️)
- [Code Overview :man\_technologist:](#code-overview-man_technologist)
- [Example Output :scroll:](#example-output-scroll)
- [License :page\_facing\_up:](#license-page_facing_up)


## Features :sparkles:

- **Concurrency**: Handles multiple clients using goroutines.
- **Custom message handling**: Reads data from clients and logs the messages to the server console.
- **Graceful shutdown**: The server can be stopped gracefully using system signals (e.g., `Ctrl+C`).
- **Minimal dependencies**: Built using the Go standard library with no third-party packages.

## How it Works :hammer_and_wrench:

1. The server listens on a specified TCP port.
2. It accepts incoming connections and spins up a new goroutine to handle each connection.
3. For each connection, the server reads incoming data from the client and logs it.
4. The server continues to accept connections and can be gracefully terminated using a system signal.

## TODO's :memo:
Here are some planned features and improvements for the TCP server:

- [ ] Add peer tracking: Implement a peerMap of type map[net.Addr] to store and track information about connected peers.
    - This will allow the server to:
      - Identify peer connections.
      - Track messages from specific peers.
      - Log and display peer-specific information.

 - [ ] Message broadcasting: Enable the server to broadcast a message from one client to all connected peers.
 - [ ] Connection status management: Keep track of active and inactive connections.
 - [ ] Add unit tests: Write unit tests to ensure the server behaves correctly under various scenarios (e.g., connection errors, message parsing).

## Getting Started :rocket:

### Prerequisites :memo:

- [Go](https://golang.org/dl/) (version 1.16 or later)

### Running the Server :arrow_forward:

1. Clone this repository:

   ```bash
   git clone https://github.com/0xprasanth/tcp_server.git
   ```

2. Navigate to the project directory:

   ```bash
   cd tcp_server
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

   The server will start listening on port `:3000`.

### Connecting to the Server 💻

To test the server, you can use tools like `telnet` or `nc` (Netcat):

```bash
nc localhost 3000
```

Send messages through the client, and the server will log them to the console.

### Stopping the Server ⏹️

The server can be stopped gracefully by pressing `Ctrl+C` in the terminal where it's running. The server will handle the signal and stop accepting connections before exiting.

## Code Overview :man_technologist:

- `Server`: A struct that encapsulates the server logic, including the listener and quit channel.
- `Message`: this struct contains information about send and the payload from the client.
- `Start()`: Starts the TCP server and listens for incoming connections. It also waits for a quit signal to stop the server.
- `acceptLoop()`: Accepts incoming connections in a loop and hands off each connection to a new goroutine.
- `readLoop()`: Reads data from each client connection, logs the messages, and closes the connection when done.

## Example Output :scroll:

When a client connects and sends a message, the server logs the message:

```bash
Received message: Hello, server!
```

## License :page_facing_up:

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
