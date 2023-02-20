migrate:
	go run cmd/migrate-db/main.go

generate: generate-tools generate-ent generate-gqlgen

generate-tools:
	go generate ./tools.go

generate-ent: generate-tools
	./bin/ent generate ./ent/schema

generate-gqlgen: generate-tools
	./bin/gqlgen
