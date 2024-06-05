package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var clients []net.Conn
var mutex = &sync.Mutex{}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	mutex.Lock()
	clients = append(clients, conn)
	mutex.Unlock()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		message = strings.TrimSpace(message)
		fmt.Printf("Received message from %s: %s\n", conn.RemoteAddr().String(), message)

		broadcastMessage(conn, message)
	}
}

func broadcastMessage(sender net.Conn, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, client := range clients {
		if client != sender {
			_, err := client.Write([]byte("Client " + sender.RemoteAddr().String() + ": " + message + "\n"))
			if err != nil {
				fmt.Println("Error writing message to client:", err)
			}
		}
	}
}
