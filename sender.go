package main

import "bufio"
import "fmt"
import "net"
import "os"

func main() {
	port := ":" + os.Args[1]

	fmt.Printf("Sending connection request on port %v...\n", port)

    // Dial
    _, err := net.Dial("tcp", port)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	fmt.Printf("Established connection.\n")

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
		// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    	// status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
        	os.Exit(1)
		}
		// fmt.Println(status)
	}

}
