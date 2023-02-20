# Example GraphQL Server by gqlgen and ent

Setup `.env`:
```zsh
copy .env.example .env
# and Fill `DATABASE_URL` in .env
```

Install tools to `./bin`:
```zsh
make generate-tools
```

DB Migration:
```zsh
make migrate-db
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
