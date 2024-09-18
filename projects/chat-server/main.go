package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	clients   = make(map[net.Conn]string) // map to store client connection and name
	broadcast = make(chan string)
	mutex     = sync.Mutex{}
)

func main() {
	println("Simple Chat Server")

	// Start TCP server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listener", listener)
	fmt.Println("Chat server started on port 8080")

	// Goroutine to handle broadcasting messages to all clients
	go handleBroadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each client connection concurrently
		go handleConnection(conn)
	}

}

// Function to handle incoming connections
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Ask the user for their name
	conn.Write([]byte("Enter your name: "))
	name, _ := bufio.NewReader(conn).ReadString('\n')
	name = strings.TrimSpace(name)

	mutex.Lock()
	clients[conn] = name
	mutex.Unlock()

	fmt.Printf("%s joined the chat\n", name)
	broadcast <- fmt.Sprintf("%s has joined the chat", name)

	conn.Write([]byte("Welcome to the chat, type /list to see connected users.\n"))

	// Read client input in a loop
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// Handle disconnect
			fmt.Printf("%s left the chat\n", name)
			broadcast <- fmt.Sprintf("%s has left the chat", name)

			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()

			return
		}

		message = strings.TrimSpace(message)

		// Handle commands
		if message == "/list" {
			listConnectedUsers(conn)
		} else {
			broadcast <- fmt.Sprintf("%s: %s", name, message)
		}
	}

}

// Function to broadcast messages to all clients
func handleBroadcast() {
	for {
		// Receive message from broadcast channel
		message := <-broadcast

		// Send the message to every connected client
		mutex.Lock()
		for conn := range clients {
			conn.Write([]byte(message + "\n"))
		}
		mutex.Unlock()
	}
}

// Function to list all connected users
func listConnectedUsers(conn net.Conn) {
	mutex.Lock()
	conn.Write([]byte("Connected users:\n"))
	for _, name := range clients {
		conn.Write([]byte(name + "\n"))
	}
	mutex.Unlock()
}
