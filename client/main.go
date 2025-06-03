package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Client represents the chat client
type Client struct {
	conn     net.Conn
	username string
}

// NewClient creates a new client instance
func NewClient(host string) (*Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}

	return &Client{
		conn: conn,
	}, nil
}

// Start begins the client session
func (c *Client) Start() {
	defer c.conn.Close()

	// Handle server messages in a separate goroutine
	go c.receiveMessages()

	// Read username input
	reader := bufio.NewReader(os.Stdin)

	// Read the server's username prompt
	serverPrompt, _ := bufio.NewReader(c.conn).ReadString('\n')
	fmt.Print(serverPrompt)

	// Get username from user
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading username: %v\n", err)
		return
	}

	c.username = strings.TrimSpace(username)

	// Send username to server
	c.conn.Write([]byte(c.username + "\n"))

	// Read welcome message
	welcomeMsg, _ := bufio.NewReader(c.conn).ReadString('\n')
	fmt.Print(welcomeMsg)

	// Show help
	c.showHelp()

	// Main input loop
	fmt.Printf("\n💬 Chat as %s (type '/help' for commands):\n", c.username)
	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("\nError reading input: %v\n", err)
			break
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		// Handle client commands
		if strings.HasPrefix(message, "/") {
			if c.handleCommand(message) {
				break // Exit if command returns true
			}
			continue
		}

		// Send message to server
		_, err = c.conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Printf("Failed to send message: %v\n", err)
			break
		}
	}
}

// receiveMessages handles incoming messages from the server
func (c *Client) receiveMessages() {
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n❌ Connection to server lost.")
			os.Exit(1)
		}

		message = strings.TrimSpace(message)
		if message != "" {
			// Clear current line and print message
			fmt.Printf("\r%s\n> ", message)
		}
	}
}

// handleCommand processes client-side commands
func (c *Client) handleCommand(command string) bool {
	switch command {
	case "/exit", "/quit":
		fmt.Println("👋 Goodbye!")
		c.conn.Write([]byte("/exit\n"))
		return true
	case "/help":
		c.showHelp()
		return false
	default:
		fmt.Printf("❓ Unknown command: %s (type '/help' for available commands)\n", command)
		return false
	}
}

// showHelp displays available commands
func (c *Client) showHelp() {
	fmt.Println("\n📖 Available commands:")
	fmt.Println("  /help  - Show this help message")
	fmt.Println("  /exit  - Leave the chat")
	fmt.Println("  /quit  - Leave the chat")
	fmt.Println("\n💡 Simply type your message and press Enter to chat!")
}

func main() {
	// Default server address
	host := "localhost:9000"

	// Allow custom host via command line argument
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	fmt.Println("🔗 Terminal Chat Client")
	fmt.Printf("Connecting to server at %s...\n", host)

	client, err := NewClient(host)
	if err != nil {
		fmt.Printf("❌ Failed to connect: %v\n", err)
		fmt.Println("\n💡 Make sure the server is running on", host)
		fmt.Println("Usage: go run client/main.go [host:port]")
		os.Exit(1)
	}

	fmt.Println("✅ Connected to chat server!")
	client.Start()
}
