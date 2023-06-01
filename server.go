package main

func NewServer() *Server {
	return &Server{}
}

func (*Server) Start() error {
	return nil
}

func (*Server) Stop() error {
	return nil
}

type Server struct {
	HostName string
	Port     int
}
