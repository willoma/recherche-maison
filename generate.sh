#!/bin/sh

go tool sqlc generate --file db/sqlc.yaml
go tool templ generate -path web
go mod tidy
