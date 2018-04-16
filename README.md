# simpletcp
Golang server for simple tcp connection

# Demo
```
go run cmd/main.go -addr ":8080" # default is :8080
# Run in other terminal
go run cmd/client.go -addr ":8080"
```

# Usage
```
import (
  server "github.com/aiqu/simpletcpserver"
 )
 
func main() {
  s, _ := server.New(":8080")
  for {
    time.Sleep(time.Second)
    // Ping all connected clients every seconds
    s.PingAll()
  }
}
```
