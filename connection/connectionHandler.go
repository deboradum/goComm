package connection

import "bufio"
import "fmt"
import "net"
import "os"
import "time"

func HandleConnectionListen(conn net.Conn) {
    for {
		buffer := make([]byte, 1024)
		size, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if size == 0 {
			continue
		}
		respText := string(buffer)
		time := time.Now().Format("15:04:05")
		fmt.Print("\n")
		fmt.Printf(time + ": " +  respText)
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