# github.com/Sach97/ninshoo
A package for dearling with user registration and authentification in go 

WIP.


## Instructions 

First 
```shell
source .env
```


## TODOs

- [x] replace .env logic with config.toml where necessary
- [x] create jwt
- [ ] strategy pattern at the root for choosing between Rest or Graphql
- [ ] create roles through jwt and test it on hasura
- [x] create middleware
- [x] create a boilerplate in graphql with gqlgen (chi, mux)
- [ ] create a rest boilerplate using different routers (chi, mux)
- [ ] update examples with .gitlab-ci.yaml for CI/CD integrations with kubernetes
- [ ] tests


## Inspirations
This package is inspired by :
- [gotrue](https://github.com/netlify/gotrue)
- [go-graphql-starter](https://github.com/OscarYuen/go-graphql-starter/issues/22)
