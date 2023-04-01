package main

import (
	"github.com/glebpepega/goodvibesbot/config"
	"github.com/glebpepega/goodvibesbot/fetcher"
	"github.com/glebpepega/goodvibesbot/sender"
)

func main() {
	c := config.NewServerConfig()
	sender.Sender(c.UpdateChan)
	fetcher.GetUpdates(c.StructResponse, &c.Offset, c.UpdateChan)
}
