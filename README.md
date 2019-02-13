# gqlgenauth
A package for dearling with user registration and authentification in go 

WIP.


## Instructions 

First 
```shell
source .env
```


## TODOs

- [ ] replace .env logic with config.toml where necessary
- [x] create jwt
- [ ] create roles through jwt and test it on hasura
- [ ] create middleware
- [ ] create a boilerplate in graphql with gqlgen (chi, mux)
- [ ] create a rest boilerplate using different routers (chi, mux)


## Inspirations
This package is inspired by :
- [gotrue](https://github.com/netlify/gotrue)
- [go-graphql-starter](https://github.com/OscarYuen/go-graphql-starter/issues/22)