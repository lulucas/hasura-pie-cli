# Hasura-pie-cli

A CLI for https://github.com/lulucas/hasura-pie

# Quick start

Install CLI

```
go get -u github.com/lulucas/hasura-pie-cli
```

# Document

## Command

Generate module

```
pie g m account
```

Sync model from postgres tables

```
# sync specific table
pie s m users

# sync all tabls
pie s m
```

## Config

Create config.yml in cli work directory

```
postgres:
  host: 127.0.0.1
  user: postgres
  password: postgres
hasura:
  endpoint: http://localhost:8080
  admin_key: 123
```

# Related project

* https://github.com/ekhabarov/go-pg-generator
