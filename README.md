# GoComm
To gain experience with Go, I wanted to create a server client modelled chat. There are two versions: unencrypted chat and encrypted chat. Unencrypted uses the net go package, while the encrypted version uses crypte/tls.

# Usage
To create the needed certificates for the encrypted version, first run:
```
./makecerts.sh <email>
```

After the certificates are created, you can run the client and server as follows:
```
go run server.go <addres:port>
go run client.go <addres:port>
```

That's all. If you want to run the unencrypted version, you can skip the first step of course.
