package app

const dockerfileTpl = `FROM alpine

RUN apk update && apk add ca-certificates
RUN apk add tzdata

RUN mkdir -p /app
WORKDIR /app

COPY main /app/main

CMD ["./main"]`

const rsyncTpl = `+ main
+ .env.prod
+ Dockerfile
+ docker-compose.prod.yml
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
API_HOST=api.local
HTTP_HOST=business.local
`

const envProdTpl = `APP_PRODUCTION=true
API_HOST=api.prod
HTTP_HOST=business.prod
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
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD:-postgres}

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
    image: hasura/graphql-engine:v1.2.0-beta.3
    ports:
      - 8080:8080
    depends_on:
      - postgres
    restart: always
    environment:
      HASURA_GRAPHQL_DATABASE_URL: postgres://${POSTGRES_USER:-postgres}:${POSTGRES_PASSWORD:-postgres}@${POSTGRES_HOST:-127.0.0.1}:${POSTGRES_PORT:-5432}/${POSTGRES_DATABASE:-postgres}
      HASURA_GRAPHQL_ENABLE_CONSOLE: "true"
      HASURA_GRAPHQL_ENABLED_LOG_TYPES: startup, http-log, webhook-log, websocket-log, query-log
      HASURA_GRAPHQL_UNAUTHORIZED_ROLE: anonymous
      HASURA_GRAPHQL_ADMIN_SECRET: ${HASURA_GRAPHQL_ADMIN_SECRET:-123456}
      HASURA_GRAPHQL_JWT_SECRET: |
        {
          "type": "HS256",
          "key": "${APP_JWT_KEY}"
        }
      EVENT_ENDPOINT: http://business:${INTERNAL_PORT:-3000}/events
      ACTION_ENDPOINT: http://business:${INTERNAL_PORT:-3000}/actions

volumes:
  db_data:
  redis_data:
`

const dockerComposeProdTpl = `version: '3.6'
services:
  graphql-engine:
    image: hasura/graphql-engine:v1.2.0-beta.3
    ports:
      - 8080:8080
    depends_on:
      - postgres
    restart: always
    environment:
      HASURA_GRAPHQL_DATABASE_URL: postgres://${POSTGRES_USER:-postgres}:${POSTGRES_PASSWORD:-postgres}@${POSTGRES_HOST:-127.0.0.1}:${POSTGRES_PORT:-5432}/${POSTGRES_DATABASE:-postgres}
      HASURA_GRAPHQL_ENABLE_CONSOLE: "true"
      HASURA_GRAPHQL_ENABLED_LOG_TYPES: startup, http-log, webhook-log, websocket-log, query-log
      HASURA_GRAPHQL_UNAUTHORIZED_ROLE: anonymous
      HASURA_GRAPHQL_ADMIN_SECRET: ${HASURA_GRAPHQL_ADMIN_SECRET:-123456}
      HASURA_GRAPHQL_JWT_SECRET: |
        {
          "type": "HS256",
          "key": "${APP_JWT_KEY}"
        }
      EVENT_ENDPOINT: http://business:${INTERNAL_PORT:-3000}/events
      ACTION_ENDPOINT: http://business:${INTERNAL_PORT:-3000}/actions
`
