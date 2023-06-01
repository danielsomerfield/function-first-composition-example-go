package server

import (
	"function-first-composition-example-go/review-server/ratings"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"strconv"
)

func NewServer(host string, port int) *Server {
	return &Server{
		Engine:      gin.Default(),
		Shutdown:    make(chan int),
		Startup:     make(chan int),
		Address:     host + ":" + strconv.Itoa(port),
		ServiceName: "review-server",
	}
}

func (server *Server) Start() error {
	if server.HTTPServer != nil {
		log.Panicf("server already started")
	}

	server.HTTPServer = &http.Server{
		Addr:    server.Address,
		Handler: server.Engine,
	}

	ratings.Initialize(server.Engine)

	go func() {
		log.Printf("%v about to start running at %v", server.ServiceName, server.HTTPServer.Addr)
		server.HTTPServer.RegisterOnShutdown(func() {
			server.Shutdown <- 0
		})

		listener, err := net.Listen("tcp", server.HTTPServer.Addr)
		if err != nil {
			log.Fatalf("listen failed: %s\n", err)
		}

		server.Startup <- 0

		log.Printf("%v starting listening at %v", server.ServiceName, server.HTTPServer.Addr)

		if err := server.HTTPServer.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

		log.Printf("%v finished running at %v", server.ServiceName, server.HTTPServer.Addr)
	}()

	<-server.Startup
	log.Printf("%v running at %v", server.ServiceName, server.HTTPServer.Addr)

	return nil
}

func (server *Server) Stop() error {
	return nil
}

type Server struct {
	Engine      *gin.Engine
	Shutdown    chan int
	Startup     chan int
	Address     string
	HTTPServer  *http.Server
	ServiceName string
}
