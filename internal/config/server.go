package config

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	User string `yaml:"user"`
}

func (server *Server) SetHost(host string) {
	server.Host = host
}

func (server *Server) SetName(name string) {
	server.Name = name
}
