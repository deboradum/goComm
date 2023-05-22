package main

import "fmt"
import "net"
import "os"
import C "gocomm/connection"
import "runtime"

func main() {
	port := ":" + os.Args[1]

	// Sends handshake request.
	fmt.Printf("Sending connection request on port %v...\n", port)
    conn, err := net.Dial("tcp", port)
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
