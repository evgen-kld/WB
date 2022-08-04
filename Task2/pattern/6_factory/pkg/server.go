package pkg

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintInfo() {
	fmt.Println(s.Type, s.Core, s.Memory)
}
