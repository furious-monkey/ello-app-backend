package main

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/channels/internal/server"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/commands"
)

func main() {
	commands.Run(server.New())
}
