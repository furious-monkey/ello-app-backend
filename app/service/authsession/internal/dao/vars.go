package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

const (
	saltTimeout = 30 * 60 // salt timeout
)

func removeAllNil(vList []*mtproto.Authorization) []*mtproto.Authorization {
	for i := 0; i < len(vList); {
		if vList[i] != nil {
			i++
			continue
		}

		if i < len(vList)-1 {
			copy(vList[i:], vList[i+1:])
		}

		vList[len(vList)-1] = nil
		vList = vList[:len(vList)-1]
	}

	return vList
}
