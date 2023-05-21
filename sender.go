package main

import "bufio"
import "fmt"
import "net"
import "os"

func main() {
	port := ":" + os.Args[1]

	fmt.Printf("Sending connection request on port %v...\n", port)

    // Dial
    conn, err := net.Dial("tcp", port)

    if err != nil {
        // handle error
    }
	fmt.Printf("Established connection.\n")
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    status, err := bufio.NewReader(conn).ReadString('\n')


    fmt.Println(status)
}
