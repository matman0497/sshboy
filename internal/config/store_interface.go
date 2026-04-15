package config

type ServerStore interface {
	List() []Server
	Get(name string) *Server
	Save() error
	Delete(name string) error
}
