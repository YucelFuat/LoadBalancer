package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "sync"
)

var (
    counter     int
    counterLock sync.Mutex // Ensures thread-safe access to the counter

    listenAddr = "localhost:8080"

    servers = []string{
        "localhost:5001",
        "localhost:5002",
        "localhost:5003",
    }
)

func main() {
    // Start listening for incoming connections
    listener, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatalf("Failed to listen: %s", err)
    }
    defer listener.Close()

    // Accept and handle incoming connections
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %s", err)
            continue
        }

        // Choose backend server to proxy to
        backend := chooseBackend()
        fmt.Printf("Selected backend: %s\n", backend)

        // Start a new goroutine to handle the proxying
        go func() {
            err := proxy(backend, conn)
            if err != nil {
                log.Printf("WARNING: proxying failed: %v", err)
            }
        }()
    }
}

// Proxy function to handle data transfer between client and backend server
func proxy(backend string, c net.Conn) error {
    // Connect to the backend server
    bc, err := net.Dial("tcp", backend)
    if err != nil {
        return fmt.Errorf("failed to connect to backend %s: %v", backend, err)
    }

    // Transfer data from client to backend server
    go func() {
        io.Copy(bc, c)
    }()

    // Transfer data from backend server to client
    go func() {
        io.Copy(c, bc)
    }()

    return nil
}

// ChooseBackend selects the backend server in a round-robin fashion
func chooseBackend() string {
    // Lock to ensure counter is updated atomically
    counterLock.Lock()
    defer counterLock.Unlock()

    // Calculate index of backend server to use
    backendIndex := counter % len(servers)
    // Increment counter for next selection
    counter++
    
    return servers[backendIndex]
}


