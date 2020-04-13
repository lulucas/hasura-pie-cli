# Hasura-pie-cli

A CLI for https://github.com/lulucas/hasura-pie

# Quick start

Install CLI

```
go get -u github.com/lulucas/hasura-pie-cli/pie
```

# Document

## Command

Init a project

```
pie init
```

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


Sync module code
```yaml
# config.yml

sync:
  - remote: infra/db
  - remote: infra/redis
  - remote: infra/pay
  - remote: infra/captcha
  - remote: infra/sms
  - remote: account
  - remote: finance
    local: finance
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
