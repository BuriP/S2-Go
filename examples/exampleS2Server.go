//go:build ignore

// + build ignore
package main

import (
	"github.com/BuriP/S2-Go/websockets/s2server"
)

func main() {
	server := s2server.NewWebSocketServer()
	server.Run(":8080")
}
