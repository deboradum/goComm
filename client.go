package main

import "fmt"
import "net"
import "os"
import C "gocomm/connection"

func main() {
	port := ":" + os.Args[1]

	fmt.Printf("Sending connection request on port %v...\n", port)

    // Dial
    conn, err := net.Dial("tcp", port)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	fmt.Printf("Established connection.\n")
	for {
		go C.HandleConnectionListen(conn)
		go C.HandleConnectionWrite(conn)
	}

}
