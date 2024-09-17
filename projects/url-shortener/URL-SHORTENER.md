# URL Shortener

Goal: Build a basic URL shortener service.

Features:
Accept long URLs and return a shortened version.
Redirect shortened URLs to the original.
Store URL mappings in a database or a simple in-memory map.

Learning: HTTP routing, URL encoding, basic web server handling.

First, initialize a Go module and install any necessary packages.
```bash
go mod init url-shortener
go get github.com/gorilla/mux
```

Start the server by running the Go program:
```bash
go run main.go
```

You can shorten a URL by sending a POST request using curl or Postman:
```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://www.google.com"}'
```

You can then test the redirect functionality by visiting the shortened URL:
```bash
curl http://localhost:8080/<shortenedURL>
```