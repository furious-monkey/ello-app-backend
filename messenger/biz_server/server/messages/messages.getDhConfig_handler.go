// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package messages

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/crypto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

var (
	// TODO(@benqi): 直接指定了dh2048_p和dh2048_g!!!
	dh2048_p = []byte{
		0xF8, 0x87, 0xA5, 0x15, 0x98, 0x35, 0x20, 0x1E, 0xF5, 0x81, 0xE5, 0x95,
		0x1B, 0xE4, 0x54, 0xEA, 0x53, 0xF5, 0xE7, 0x26, 0x30, 0x03, 0x06, 0x79,
		0x3C, 0xC1, 0x0B, 0xAD, 0x3B, 0x59, 0x3C, 0x61, 0x13, 0x03, 0x7B, 0x02,
		0x70, 0xDE, 0xC1, 0x20, 0x11, 0x9E, 0x94, 0x13, 0x50, 0xF7, 0x62, 0xFC,
		0x99, 0x0D, 0xC1, 0x12, 0x6E, 0x03, 0x95, 0xA3, 0x57, 0xC7, 0x3C, 0xB8,
		0x6B, 0x40, 0x56, 0x65, 0x70, 0xFB, 0x7A, 0xE9, 0x02, 0xEC, 0xD2, 0xB6,
		0x54, 0xD7, 0x34, 0xAD, 0x3D, 0x9E, 0x11, 0x61, 0x53, 0xBE, 0xEA, 0xB8,
		0x17, 0x48, 0xA8, 0xDC, 0x70, 0xAE, 0x65, 0x99, 0x3F, 0x82, 0x4C, 0xFF,
		0x6A, 0xC9, 0xFA, 0xB1, 0xFA, 0xE4, 0x4F, 0x5D, 0xA4, 0x05, 0xC2, 0x8E,
		0x55, 0xC0, 0xB1, 0x1D, 0xCC, 0x17, 0xF3, 0xFA, 0x65, 0xD8, 0x6B, 0x09,
		0x13, 0x01, 0x2A, 0x39, 0xF1, 0x86, 0x73, 0xE3, 0x7A, 0xC8, 0xDB, 0x7D,
		0xDA, 0x1C, 0xA1, 0x2D, 0xBA, 0x2C, 0x00, 0x6B, 0x2C, 0x55, 0x28, 0x2B,
		0xD5, 0xF5, 0x3C, 0x9F, 0x50, 0xA7, 0xB7, 0x28, 0x9F, 0x22, 0xD5, 0x3A,
		0xC4, 0x53, 0x01, 0xC9, 0xF3, 0x69, 0xB1, 0x8D, 0x01, 0x36, 0xF8, 0xA8,
		0x89, 0xCA, 0x2E, 0x72, 0xBC, 0x36, 0x3A, 0x42, 0xC1, 0x06, 0xD6, 0x0E,
		0xCB, 0x4D, 0x5C, 0x1F, 0xE4, 0xA1, 0x17, 0xBF, 0x55, 0x64, 0x1B, 0xB4,
		0x52, 0xEC, 0x15, 0xED, 0x32, 0xB1, 0x81, 0x07, 0xC9, 0x71, 0x25, 0xF9,
		0x4D, 0x48, 0x3D, 0x18, 0xF4, 0x12, 0x09, 0x32, 0xC4, 0x0B, 0x7A, 0x4E,
		0x83, 0xC3, 0x10, 0x90, 0x51, 0x2E, 0xBE, 0x87, 0xF9, 0xDE, 0xB4, 0xE6,
		0x3C, 0x29, 0xB5, 0x32, 0x01, 0x9D, 0x95, 0x04, 0xBD, 0x42, 0x89, 0xFD,
		0x21, 0xEB, 0xE9, 0x88, 0x5A, 0x27, 0xBB, 0x31, 0xC4, 0x26, 0x99, 0xAB,
		0x8C, 0xA1, 0x76, 0xDB,
	}

	dh2048_g = []byte{0x03}
)

/*
	request:
	{ messages_getDhConfig
	  version: 0 [INT],
	  random_length: 256 [INT],
	}

	reply:
	{ messages_dhConfig
		g: 3 [INT],
		p: C7 1C AE B9 C6 B1 C9 04 8E 6C 52 2F 70 F1 3F 73... [256 BYTES],
		version: 3 [INT],
		random: F1 2A FB 6B 97 B6 0A 17 B9 3E 2F 65 28 33 4D 03... [256 BYTES],
	}
*/
// messages.getDhConfig#26cf8950 version:int random_length:int = messages.DhConfig;
func (s *MessagesServiceImpl) MessagesGetDhConfig(ctx context.Context, request *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetDhConfig - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): 直接设定P和G
	dhConfig := mtproto.NewTLMessagesDhConfig()
	dhConfig.SetG(2)
	dhConfig.SetP(dh2048_p)
	dhConfig.SetVersion(3)
	dhConfig.SetRandom(crypto.GenerateNonce(int(request.GetRandomLength())))

	glog.Infof("RpcMetadataFromIncoming - reply {%v}", dhConfig)
	return dhConfig.To_Messages_DhConfig(), nil
}
