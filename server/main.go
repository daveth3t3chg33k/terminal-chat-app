package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// Client represents a connected client
type Client struct {
	conn     net.Conn
	username string
	send     chan string
}

// Server manages the chat server
type Server struct {
	clients    map[*Client]bool
	broadcast  chan string
	register   chan *Client
	unregister chan *Client
	mutex      sync.RWMutex
}

// NewServer creates a new server instance
func NewServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the server hub
func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.mutex.Lock()
			s.clients[client] = true
			s.mutex.Unlock()
			log.Printf("User '%s' joined the chat", client.username)
			
			// Notify all clients about the new user
			joinMessage := fmt.Sprintf("[%s] *** %s joined the chat ***", 
				time.Now().Format("15:04:05"), client.username)
			s.broadcast <- joinMessage

		case client := <-s.unregister:
			s.mutex.Lock()
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
				s.mutex.Unlock()
				log.Printf("User '%s' left the chat", client.username)
				
				// Notify all clients about the user leaving
				leaveMessage := fmt.Sprintf("[%s] *** %s left the chat ***", 
					time.Now().Format("15:04:05"), client.username)
				s.broadcast <- leaveMessage
			} else {
				s.mutex.Unlock()
			}

		case message := <-s.broadcast:
			s.mutex.RLock()
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					delete(s.clients, client)
					close(client.send)
				}
			}
			s.mutex.RUnlock()
		}
	}
}

// handleClient manages individual client connections
func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()

	log.Printf("Starting handshake with client %s", conn.RemoteAddr())

	// Request username
	conn.Write([]byte("Enter your username: \n"))
	log.Printf("Sent username prompt to client %s", conn.RemoteAddr())
	
	reader := bufio.NewReader(conn)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading username from %s: %v", conn.RemoteAddr(), err)
		return
	}
	username = strings.TrimSpace(username)
	
	log.Printf("Received username '%s' from client %s", username, conn.RemoteAddr())
	
	if username == "" {
		log.Printf("Empty username from client %s, disconnecting", conn.RemoteAddr())
		conn.Write([]byte("Username cannot be empty. Disconnecting.\n"))
		return
	}

	// Create client
	client := &Client{
		conn:     conn,
		username: username,
		send:     make(chan string, 256),
	}

	// Register client
	s.register <- client

	// Send welcome message
	welcomeMsg := fmt.Sprintf("Welcome to the chat, %s! Type '/exit' to quit.\n", username)
	conn.Write([]byte(welcomeMsg))

	// Start goroutines for handling client
	go s.writeMessages(client)
	go s.readMessages(client)

	// Wait for the read goroutine to finish
	select {}
}

// readMessages reads messages from the client
func (s *Server) readMessages(client *Client) {
	defer func() {
		s.unregister <- client
		client.conn.Close()
	}()

	reader := bufio.NewReader(client.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		// Handle exit command
		if message == "/exit" {
			break
		}

		// Broadcast the message with timestamp and username
		timestamp := time.Now().Format("15:04:05")
		formattedMessage := fmt.Sprintf("[%s] %s: %s", timestamp, client.username, message)
		s.broadcast <- formattedMessage
	}
}

// writeMessages sends messages to the client
func (s *Server) writeMessages(client *Client) {
	defer client.conn.Close()

	for message := range client.send {
		if _, err := client.conn.Write([]byte(message + "\n")); err != nil {
			break
		}
	}
}

func main() {
	const port = ":9000"
	
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Printf("🚀 Terminal Chat Server started on port %s", port)
	log.Println("Waiting for clients to connect...")

	server := NewServer()
	go server.Run()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		log.Printf("New connection from %s", conn.RemoteAddr())
		go server.handleClient(conn)
	}
} 