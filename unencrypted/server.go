package main

import C "gocomm/connection"
import "fmt"
import "net"
import "os"
import "runtime"

func main() {
	address := os.Args[1]

	// Listens for handshake request.
	fmt.Printf("Listening on address %v...\n", address)
    ln, err := net.Listen("tcp",  address)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Accepts handshake request.
    conn, err := ln.Accept()
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
