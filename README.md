# Example GraphQL Server by gqlgen and ent
Setup `.env`:
```zsh
copy .env.example .env
# and Fill `DATABASE_URL` in .env
```

Install tools to `./bin`:
```zsh
go generate ./tools.go
```

DB Migration:
```zsh
go run cmd/migrate-db/main.go
```

Create Dummy Data:
```zsh
go run cmd/create-dummies/main.go
```

Run Server:
```zsh
go run server.gp
```

Access GraphiQL Client: http://localhost:8080

### Others
regenerate ent:
```zsh
./bin/ent generate ./ent/schema
```

regenerate gqlgen:
```zsh
./bin/gqlgen
```
