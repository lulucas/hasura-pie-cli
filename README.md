# Hasura-pie-cli

A CLI for https://github.com/lulucas/hasura-pie

# Quick start

Install Pie CLI

```
go get -u github.com/lulucas/hasura-pie-cli/pie
```

Install Hasura CLI

https://hasura.io/docs/1.0/graphql/manual/hasura-cli/install-hasura-cli.html

# Document

## Command

Init a project

```
pie init myproject
```

Generate module

```
pie g m account
```

Generate action TODO

```
pie g a login
```

Generate event TODO

```
pie g e user_created
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

sync:
  module:
    - remote: infra/redis
    - remote: infra/pay
    - remote: infra/captcha
    - remote: infra/sms
    - remote: account
    - remote: finance
    - remote: analysis

```

# Related project

* https://github.com/ekhabarov/go-pg-generator
