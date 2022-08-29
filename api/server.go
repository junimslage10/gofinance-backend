package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/junimslage10/gofinance-backend/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func newServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Api as error": err.Error()}
}
