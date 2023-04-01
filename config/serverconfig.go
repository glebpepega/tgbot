package config

import "github.com/glebpepega/goodvibesbot/update"

type ServerConfig struct {
	StructResponse *update.UpdateResponse
	Offset         int
	UpdateChan     chan update.Update
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		StructResponse: update.NewResponse(),
		Offset:         0,
		UpdateChan:     make(chan update.Update),
	}
}
