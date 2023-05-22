package connection

import "bufio"
import "fmt"
import "net"
import "os"

func HandleConnectionListen(conn net.Conn) {
    for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Print("\n")
			fmt.Println(err)
			os.Exit(1)
		}
		respText := string(buffer)
		fmt.Print("\n")
		fmt.Printf("Anon: " +  respText)
		fmt.Printf("You: ")
    }
}

func HandleConnectionWrite(conn net.Conn) {
	// Create input reader
	reader := bufio.NewReader(os.Stdin)
    for {
		fmt.Printf("You: ")
		text, _ := reader.ReadString('\n')
		// Skip empty text.
		if text == "\n" {
			continue
		}
		textBytes := []byte(text)
		conn.Write(textBytes)
    }
}