package main

import "fmt"

type Server struct {
	hostname  string
	ipAddress string
	arch      string
	cpu       int
	ram       int
}

func (s *Server) getHostname() string {
	return s.hostname
}
func (s *Server) setHostname(name string) {
	s.hostname = name
}

func main() {
	instance := new(Server)
	instance.setHostname("example.local")
	instance.cpu = 48
	instance.ram = 256
	fmt.Println(instance.getHostname(), &instance, *instance)
}