package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Servers []Server `yaml:"servers"`
}

var cfg *Config

func Init() {

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

func Get() *Config {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg
}

func Save() {
	bytes, _ := yaml.Marshal(cfg)

	os.WriteFile(getPath(), bytes, 0644)
}

func Add(name string, host string, user string) error {

	if GetServer(name) != nil {
		return fmt.Errorf("server %s already exists", name)
	}

	cfg.Servers = append(cfg.Servers, Server{name, host, user})

	return nil
}

func GetServer(name string) *Server {

	for i := range cfg.Servers {
		if cfg.Servers[i].Name == name {
			return &cfg.Servers[i]
		}
	}

	return nil

}

func DeleteServer(name string) bool {

	var index = -1
	for i := range cfg.Servers {
		if cfg.Servers[i].Name == name {
			index = i
		}
	}

	if index == -1 {
		return false
	}

	cfg.Servers = append(cfg.Servers[:index], cfg.Servers[index+1:]...)

	return true

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
