package main

import "fmt"

type Server struct {
	Host string
	Port int
}

type Setting func(*Server)

func WithHost(host string) Setting {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Setting {
	return func(s *Server) {
		s.Port = port
	}
}

func NewServer(setting ...Setting) *Server {
	server := &Server{Host: "127.0.0.1", Port: 3300}
	for _, set := range setting {
		set(server)
	}
	return server
}

type S struct {
	Name string
}

func ss() *S {
	return &S{Name: "Wilson"}
}

func (s *S) printME() {

	fmt.Println(s.Name)
}

func init() {

	server := NewServer(
		WithHost("localhost"),
		WithPort(3000),
	)

	fmt.Printf("Host:\t%s\nPort:\t%d\n", server.Host, server.Port)

	ss().printME()
}
