package tpl

const Module = `package {{.Module}}

import (
	"github.com/lulucas/hasura-pie"
)

type {{.Module}} struct {
}

func New() *{{.Module}} {
	return &{{.Module}}{}
}

func (m *{{.Module}}) BeforeCreated(bc pie.BeforeCreatedContext) {
	
}

func (m *{{.Module}}) Created(cc pie.CreatedContext) {
	
}

`
