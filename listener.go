package main

import "fmt"
import "net"
import "os"

func handleConnection(conn net.Conn) {

    for {

    }
}

func main() {
	port := ":" + os.Args[1]

	fmt.Printf("Listening on port %v...\n", port)
    ln, err := net.Listen("tcp",  port)
    if err != nil {
        // handle error
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
        }
		fmt.Printf("Established connection.\n")
        go handleConnection(conn)
    }
}
