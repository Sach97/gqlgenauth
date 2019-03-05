# Gqlgenauth and hasura integrations example
Run Hasura on docker
```
docker-compose up
source .env
go run server/server.go
hasura migrate apply --admin-secret myadminsecretkey
```