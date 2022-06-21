package service

import (
	"github.com/jorgefg4/todolist/pkg/server"
)

type Servicer interface {
	NewServer() server.Server
}
