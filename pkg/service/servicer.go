package service

import (
	"github.com/jorgefg4/todolist/pkg/server"
)

// Defines the interface for the service layer
type Servicer interface {
	// Creates a new server
	NewServer() server.Server
}
