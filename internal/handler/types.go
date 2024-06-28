package handler

import "github.com/dormoron/mist"

type Handler interface {
	RegisterRoutes(server *mist.HTTPServer)
}
