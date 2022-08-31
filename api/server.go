package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/junimslage10/gofinance-backend/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Rotas
	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)

	router.POST("/category", server.createCategory)
	router.GET("/category", server.getCategories)
	router.GET("/category/:id", server.getCategory)
	router.DELETE("/category/:id", server.deleteCategory)
	router.PUT("/category/:id", server.updateCategory)

	router.POST("/account", server.createAccount)
	router.GET("/account", server.getAccounts)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account/graph/:user_id/:type", server.getAccountGraph)
	router.GET("/account/reports/:user_id/:type", server.getAccountReports)
	router.DELETE("/account/:id", server.deleteAccount)
	router.PUT("/account/:id", server.updateAccount)

	router.POST("/login", server.login)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Api as error": err.Error()}
}
