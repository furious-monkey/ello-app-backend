/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package main

import (
	"github.com/teamgram/marmota/pkg/commands"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/code/internal/server"
)

func main() {
	commands.Run(server.New())
}
