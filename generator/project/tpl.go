package project

const modTpl = `module api

go 1.14
`

const ignoreTpl = `.idea
.vscode

*.exe
.modules
`

const configYamlTpl = `version: 2
admin_secret: 
endpoint: http://localhost:8080
metadata_directory: metadata
actions:
  kind: synchronous
  handler_webhook_baseurl: '{{ACTION_ENDPOINT}}'
`
