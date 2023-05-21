package main

import "bufio"
import "fmt"
import "net"
import "os"

func handleConnection(conn net.Conn) {
	// Create input reader
	reader := bufio.NewReader(os.Stdin)

    for {
		fmt.Printf("You: ")
		text, _ := reader.ReadString('\n')
		// Skip empty text.
		if text == "\n" {
			continue
		}
		fmt.Printf(text)
    }
}

func main() {
	port := ":" + os.Args[1]

	fmt.Printf("Listening on port %v...\n", port)
    ln, err := net.Listen("tcp",  port)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
        	os.Exit(1)
        }
		fmt.Printf("Established connection.\n")
        go handleConnection(conn)
    }
}
