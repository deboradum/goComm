package main

import C "gocomm/connection"

import "crypto/tls"
import "fmt"
import "os"
import "runtime"


func main() {
	address := os.Args[1]

	// Read cert and key.
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		fmt.Println(err)
        os.Exit(1)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	// Sends handshake request.
	fmt.Printf("Sending connection request on address %v...\n", address)
    conn, err := tls.Dial("tcp", address, cfg)
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
