# Simple Chat Server

Goal: Build a basic TCP chat server where multiple clients can connect and chat.

Features:
Handle multiple clients using goroutines.
Broadcast messages from one client to all other connected clients.
Basic user commands like /list to show all connected users.

Learning: Goroutines, channels, basic TCP connections.

Run the server by executing:
```bash
go run chat_server.go
```

### Test the Server:
Use telnet or nc (netcat) to connect to the server from multiple terminals:
```bash
telnet localhost 8080
```
OR
```bash
nc localhost 8080
```

Commands:\
The `/list` command allows clients to see all connected users.