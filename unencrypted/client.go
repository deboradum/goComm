package main

import C "gocomm/connection"
import "fmt"
import "net"
import "os"
import "runtime"

func main() {
	address := os.Args[1]

	// Sends handshake request.
	fmt.Printf("Sending connection request on address %v...\n", address)
    conn, err := net.Dial("tcp", address)
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
