#!/bin/bash

dalgen3 --xml=$1 --db=elloapp --go2=gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization/internal/dal/dataobject

gofmt -w ../dao/mysql_dao/*.go
gofmt -w ../dataobject/*.go
