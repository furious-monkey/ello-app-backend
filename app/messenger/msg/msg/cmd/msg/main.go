package main

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/commands"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/internal/server"
)

func main() {
	commands.Run(server.New())
}
