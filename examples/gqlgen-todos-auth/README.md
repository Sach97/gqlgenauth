# Gqlgenauth and hasura integrations example
Run Hasura on docker
```
docker run -d --net=host \
  -e HASURA_GRAPHQL_DATABASE_URL=postgres://postgres:@postgres:5432/postgres \
  -e HASURA_GRAPHQL_ENABLE_CONSOLE=true \
  hasura/graphql-engine:latest
```