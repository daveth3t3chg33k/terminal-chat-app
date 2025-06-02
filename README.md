# 💬 Terminal Chat App in Go

A fully functional terminal-based chat application built in Go, featuring a TCP server and multiple concurrent clients. Perfect for learning Go networking, concurrency, and building real-time applications.

## ✨ Features

### 🖥️ Server
- **TCP Server** listening on configurable port (default: `:9000`)
- **Concurrent Connections** - supports multiple clients simultaneously
- **Real-time Broadcasting** - messages sent to all connected clients
- **User Management** - tracks usernames and connection states
- **Thread-safe Operations** - uses mutexes for concurrent access
- **Graceful Handling** - manages client joins/leaves with system notifications
- **Timestamped Messages** - all messages include timestamps

### 🔌 Client
- **Easy Connection** - connects to server with simple TCP
- **Interactive Interface** - clean terminal-based chat experience
- **Concurrent I/O** - send and receive messages simultaneously
- **Command System** - built-in commands like `/help` and `/exit`
- **Graceful Exit** - handles disconnections smoothly
- **Real-time Updates** - instant message delivery

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or later installed
- Terminal/Command prompt access

### 1. Clone or Download
```bash
# If you have this code, navigate to the project directory
cd terminal-chat
```

### 2. Start the Server
Open a terminal and run:
```bash
go run server/main.go
```

You should see:
```
🚀 Terminal Chat Server started on port :9000
Waiting for clients to connect...
```

### 3. Connect Clients
Open **additional terminal windows** (one for each client) and run:
```bash
go run client/main.go
```

Or connect to a specific host:
```bash
go run client/main.go localhost:9000
```

### 4. Start Chatting!
1. Enter your username when prompted
2. Type messages and press Enter
3. See messages from other users in real-time
4. Use `/help` to see available commands
5. Use `/exit` to leave the chat

## 📖 Usage Examples

### Starting Multiple Clients
```bash
# Terminal 1 - Server
$ go run server/main.go
🚀 Terminal Chat Server started on port :9000
Waiting for clients to connect...

# Terminal 2 - Client 1
$ go run client/main.go
🔗 Terminal Chat Client
Connecting to server at localhost:9000...
✅ Connected to chat server!
Enter your username: Alice
Welcome to the chat, Alice! Type '/exit' to quit.

# Terminal 3 - Client 2
$ go run client/main.go
🔗 Terminal Chat Client
Connecting to server at localhost:9000...
✅ Connected to chat server!
Enter your username: Bob
Welcome to the chat, Bob! Type '/exit' to quit.
```

### Sample Chat Session
```
[14:30:15] *** Alice joined the chat ***
[14:30:23] *** Bob joined the chat ***
[14:30:30] Alice: Hello everyone!
[14:30:35] Bob: Hi Alice! How are you?
[14:30:42] Alice: I'm great! This chat app is awesome 🎉
[14:30:50] *** Charlie joined the chat ***
[14:30:55] Charlie: Hey folks!
```

## 🎯 Available Commands

| Command | Description |
|---------|-------------|
| `/help` | Show available commands |
| `/exit` | Leave the chat gracefully |
| `/quit` | Same as `/exit` |

## 🏗️ Project Structure

```
terminal-chat/
├── go.mod              # Go module definition
├── README.md           # This file
├── server/
│   └── main.go        # TCP server implementation
└── client/
    └── main.go        # TCP client implementation
```

## 🔧 Configuration

### Server Port
To change the server port, modify the `port` constant in `server/main.go`:
```go
const port = ":8080"  // Change from default :9000
```

### Client Connection
Connect to a different server:
```bash
go run client/main.go your-server.com:9000
```

## 🛠️ Technical Details

### Architecture
- **Server**: Multi-threaded TCP server using goroutines
- **Client**: Concurrent TCP client with separate read/write routines
- **Communication**: Plain text over TCP sockets
- **Concurrency**: Channels and mutexes for thread-safe operations

### Key Components
- **Connection Management**: Maps for tracking active clients
- **Message Broadcasting**: Channel-based message distribution
- **Error Handling**: Graceful handling of network errors
- **Resource Cleanup**: Proper connection and goroutine cleanup

## 🧪 Testing

### Test with Multiple Clients
1. Start the server
2. Open 3-5 terminal windows
3. Run the client in each window with different usernames
4. Send messages and observe real-time delivery
5. Test joining/leaving notifications

### Test Error Scenarios
- Start client without server running
- Kill server while clients are connected
- Send empty messages
- Try very long usernames

## 🚀 Building for Production

### Build Binaries
```bash
# Build server
go build -o server ./server

# Build client  
go build -o client ./client

# Run built binaries
./server
./client
```

### Cross-Platform Builds
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o server.exe ./server
GOOS=windows GOARCH=amd64 go build -o client.exe ./client

# Linux
GOOS=linux GOARCH=amd64 go build -o server ./server
GOOS=linux GOARCH=amd64 go build -o client ./client

# macOS
GOOS=darwin GOARCH=amd64 go build -o server ./server
GOOS=darwin GOARCH=amd64 go build -o client ./client
```

## 🎨 Future Enhancements

- [ ] **Private Messages** - Direct messaging between users
- [ ] **Chat Rooms** - Multiple chat channels
- [ ] **Message History** - Persistent message storage
- [ ] **User Authentication** - Login system
- [ ] **File Sharing** - Send files through chat
- [ ] **Emoji Support** - Rich text and emoji rendering
- [ ] **Config Files** - YAML/JSON configuration
- [ ] **Logging** - Advanced logging with levels
- [ ] **Web Interface** - Browser-based client

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 📞 Support

If you encounter any issues:
1. Check the server is running on the correct port
2. Verify firewall settings allow TCP connections
3. Ensure Go is properly installed
4. Check the console for error messages

---

**Happy Chatting! 🎉**

Built with ❤️ using Go's powerful concurrency features. 