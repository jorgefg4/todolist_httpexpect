package service

import (
	"github.com/jorgefg4/todolist/pkg/server"
)

type Servicer interface {
	// Creates a new server
	NewServer() server.Server
}
