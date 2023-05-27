package connection

import "bufio"
import "fmt"
import "net"
import "os"


type Message struct {
	Text []byte
}

func HandleConnectionListen(conn net.Conn) {
    for {
		buffer := Message{Text: make([]byte, 2048)}
		_, err := conn.Read(buffer.Text)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("\nPeer closed the connection")
				os.Exit(0)
			}
			fmt.Print("\n")
			fmt.Println(err)
			os.Exit(1)
		}
		// Need to decrypt
		respText := string(buffer.Text)
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
		textBytes := Message {Text: []byte(text)}
		// Need to encrypt
		conn.Write(textBytes.Text)
    }
}
