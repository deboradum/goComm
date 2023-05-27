package main

import C "gocomm/connection"

import "crypto/tls"
import "fmt"
import "os"
import "runtime"

func main() {
	address := os.Args[1]

    // Read cert and key.
    cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		fmt.Println(err)
        os.Exit(1)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", address, cfg)
	if err != nil {
		fmt.Println(err)
        os.Exit(1)
	}

	// Listens for handshake request.
	fmt.Printf("Listening on address %v...\n", address)

    // Accepts handshake request.
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("Established connection.\n")

    go C.HandleConnectionListen(conn)
    go C.HandleConnectionWrite(conn)
    // Waits until routines are done.
    runtime.Goexit()
}
