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
	session_helper "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"
)

func main() {
	commands.Run(session_helper.New())
}
