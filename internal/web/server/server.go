package server

import (
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func CreateServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (server *Server) ConfigureRoutes() {
	accountHandler := handlers.CreateAccountHandler(server.accountService)
	server.router.Post("/accounts", accountHandler.CreateAccount)
	server.router.Get("/accounts", accountHandler.GetAccounts)
}

func (server *Server) Start() error {
	server.server = &http.Server{
		Addr:    ":" + server.port,
		Handler: server.router,
	}
	return server.server.ListenAndServe()
}
