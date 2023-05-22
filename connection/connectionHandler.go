package connection

import "bufio"
import "fmt"
import "net"
import "os"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha256"

type Message struct {
	text []byte
}

func (mess Message) EncryptMessage() ([]byte, error) {
	label := []byte("mes")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &test2048Key.PublicKey, mess, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Did not send message, got error from encryption: %s\n", err)
		return nil, err
	}

	return ciphertext, nil
}

func (mess Message) DecryptMessage() ([]byte, error) {
	label := []byte("mes")
	plaintext, err := rsa.DecryptOAEP(sha256.New(), nil, test2048Key, mess, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Did not receive message, got error from decryption: %s\n", err)
		return nil, err
	}

	return plaintext, nil
}

func HandleConnectionListen(conn net.Conn) {
    for {
		buffer := make([]byte, 2048)
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
