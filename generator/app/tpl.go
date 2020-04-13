package app

const dockerfileTpl = `FROM alpine

RUN apk update && apk add ca-certificates
RUN apk add tzdata

RUN mkdir -p /app
WORKDIR /app

COPY main /app/main

CMD ["./main"]`

const rsyncTpl = `- *.go
+ main
+ .env.prod
+ Dockerfile
+ docker-compose.prod.yml
+ certs
+ certs/**
+ metadata
+ metadata/**
+ migrations
+ migrations/**
- **
`

const mainTpl = `package main

import (
	"github.com/lulucas/hasura-pie"
)

func main() {
	app := pie.NewApp()
	app.AddModule(

	)
	app.Start()
}
`

const envTpl = `APP_PRODUCTION=false
APP_ADMIN_SECRET=
APP_JWT_KEY=
`

const envProdTpl = `APP_PRODUCTION=true
APP_ADMIN_SECRET=
APP_JWT_KEY=
APP_API_HOST=api.prod
APP_REST_HOST=business.prod
APP_TLS_ENABLED=true
`

const dockerComposeTpl = `version: '3.6'
services:
  postgres:
    image: postgres:12
    restart: always
    ports:
    - 5432:5432
    volumes:
      - db_data:/var/lib/postgresql/data
    environment:
      POSTGRES_PASSWORD: ${DB_PASSWORD:-postgres}
      POSTGRES_DB: ${DB_DATABASE:-postgres}

  redis:
    image: bitnami/redis:5.0
    restart: always
    volumes:
      - redis_data:/bitnami/redis/data
    ports:
      - 6379:6379
    environment:
      ALLOW_EMPTY_PASSWORD: 'true'

  graphql-engine:
    image: hasura/graphql-engine:v1.2.0-beta.3.cli-migrations-v2
    ports:
      - 8080:8080
    depends_on:
      - postgres
    restart: always
    volumes:
      - ./migrations:/hasura-migrations
      - ./metadata:/hasura-metadata
    environment:
      HASURA_GRAPHQL_DATABASE_URL: postgres://${DB_USER:-postgres}:${DB_PASSWORD:-postgres}@${DB_HOST:-postgres}:${DB_PORT:-5432}/${DB_DATABASE:-postgres}
      HASURA_GRAPHQL_ENABLE_CONSOLE: "false"
      HASURA_GRAPHQL_ENABLED_LOG_TYPES: startup, http-log, webhook-log, websocket-log, query-log
      HASURA_GRAPHQL_UNAUTHORIZED_ROLE: anonymous
      HASURA_GRAPHQL_ADMIN_SECRET: ${APP_ADMIN_SECRET}
      HASURA_GRAPHQL_JWT_SECRET: |
        {
          "type": "HS256",
          "key": "${APP_JWT_KEY}"
        }
      EVENT_ENDPOINT: http://host.docker.internal:${INTERNAL_PORT:-3000}/events
      ACTION_ENDPOINT: http://host.docker.internal:${INTERNAL_PORT:-3000}/actions

volumes:
  db_data:
  redis_data:
`

const dockerComposeProdTpl = `version: '3.6'
services:
  nginx-proxy:
    image: jwilder/nginx-proxy
    container_name: nginx-proxy
    restart: always
    ports:
      - ${NGINX_PORT:-80}:80
      - ${NGINX_TLS_PORT:-443}:443
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock:ro
      - ./certs:/etc/nginx/certs

  graphql-engine:
    image: hasura/graphql-engine:v1.2.0-beta.3.cli-migrations-v2
    expose:
      - 8080
    restart: always
    volumes:
      - ./migrations:/hasura-migrations
      - ./metadata:/hasura-metadata
    environment:
      HASURA_GRAPHQL_DATABASE_URL: postgres://${DB_USER:-postgres}:${DB_PASSWORD:-postgres}@${DB_HOST:-127.0.0.1}:${DB_PORT:-5432}/${DB_DATABASE:-postgres}
      HASURA_GRAPHQL_ENABLE_CONSOLE: "false"
      HASURA_GRAPHQL_ENABLED_LOG_TYPES: startup, http-log, webhook-log, websocket-log, query-log
      HASURA_GRAPHQL_UNAUTHORIZED_ROLE: anonymous
      HASURA_GRAPHQL_ADMIN_SECRET: ${APP_ADMIN_SECRET}
      HASURA_GRAPHQL_JWT_SECRET: |
        {
          "type": "HS256",
          "key": "${APP_JWT_KEY}"
        }
      EVENT_ENDPOINT: http://business:${INTERNAL_PORT:-3000}/events
      ACTION_ENDPOINT: http://business:${INTERNAL_PORT:-3000}/actions

  business:
    image: api/business
    build: .
    restart: always
    depends_on:
      - graphql-engine
    env_file:
      - .env
    environment:
      VIRTUAL_HOST: ${APP_REST_HOST}
    expose:
      - ${APP_EXTERNAL_PORT:-8000}

`

const configYmlTpl = `version: 2
admin_secret: 
endpoint: http://localhost:8080
metadata_directory: metadata
`
