USER := postgres
PASS := Test_1tesT
HOST_PORT := 127.0.0.1:5432
DATABASE := notis

generate-rps:
	@protoc --go_out=./pkg --go_opt=paths=source_relative     --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative     api/notes_api.proto

migrate-test:
	@go run github.com/pressly/goose/v3/cmd/goose@latest --dir="db/migrations" postgres "postgres://$(USER):$(PASS)@$(HOST_PORT)/$(DATABASE)?sslmode=disable" reset
	@go run github.com/pressly/goose/v3/cmd/goose@latest --dir="db/migrations" postgres "postgres://$(USER):$(PASS)@$(HOST_PORT)/$(DATABASE)?sslmode=disable" up

migrate-run:
	@go run github.com/pressly/goose/v3/cmd/goose@latest --dir="db/migrations" postgres "postgres://$(USER):$(PASS)@$(HOST_PORT)/$(DATABASE)?sslmode=disable" up

migrate-drop:
	@go run github.com/pressly/goose/v3/cmd/goose@latest -dir db/migrations postgres "postgres://$(USER):$(PASS)@$(HOST_PORT)/$(DATABASE)?sslmode=disable" down-to 20240218144349

install-mfd-generator:
	@go get github.com/vmkteam/mfd-generator

mfd-model:
	@go run github.com/vmkteam/mfd-generator model -m 'docs/model/schema.mfd' -p 'db' -o 'internal/db'
	
mfd-repo:
	@go run github.com/vmkteam/mfd-generator repo -m 'docs/model/schema.mfd' -p 'db' -o 'internal/db'

#   go run github.com/pressly/goose/v3/cmd/goose@latest -dir db/migrations/ create new_kek_table sql
#   export PATH=$PATH:/usr/local/go/bin
#   export PATH=$PATH:$HOME/go/bin