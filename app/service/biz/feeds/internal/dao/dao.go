// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/feeds/internal/dao/dao/mysql_dao"
)

type Dao struct {
	*mysql_dao.Mysql
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql: mysql_dao.NewMysqlDao(db),
	}
}
