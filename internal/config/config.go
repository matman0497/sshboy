package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Store struct{}
type Config struct {
	Servers []Server `yaml:"servers"`
}

var (
	cfg   *Config
	store Store
)

func Init() {
	store.Init()
}

func List() []Server {
	return store.List()
}

func Save() error {
	store.Save()
	return nil
}

func Add(name, host, user string) error {
	return store.Add(name, host, user)
}

func GetServer(name string) *Server {
	return store.Get(name)
}

func Delete(name string) error {
	return store.Delete(name)
}

func (Store) Init() {

	data, err := os.ReadFile(getPath())
	if err != nil {

		fmt.Println(err)

	}

	c := &Config{}
	if err := yaml.Unmarshal(data, c); err != nil {

		return
	}

	cfg = c
}

func (Store) List() []Server {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.Servers
}

func (Store) Save() error {
	bytes, _ := yaml.Marshal(cfg)

	os.WriteFile(getPath(), bytes, 0644)

	return nil
}

func (s Store) Add(name string, host string, user string) error {
	if s.Get(name) != nil {
		return fmt.Errorf("server %s already exists", name)
	}

	cfg.Servers = append(cfg.Servers, Server{name, host, user})

	return nil
}

func (Store) Get(name string) *Server {
	for i := range cfg.Servers {
		if cfg.Servers[i].Name == name {
			return &cfg.Servers[i]
		}
	}

	return nil
}

func (Store) Delete(name string) error {
	var index = -1
	for i := range cfg.Servers {
		if cfg.Servers[i].Name == name {
			index = i
		}
	}

	if index == -1 {
		return nil
	}

	cfg.Servers = append(cfg.Servers[:index], cfg.Servers[index+1:]...)

	return nil
}

func getPath() string {
	home, _ := os.UserHomeDir()

	if os.Getenv("TESTING") == "true" {
		return fmt.Sprintf(
			"%s/.sshboy/inventory-testing.yaml",
			home,
		)
	}

	return fmt.Sprintf(
		"%s/.sshboy/inventory.yaml",
		home,
	)
}
