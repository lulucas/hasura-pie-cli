package module

const Module = `package {{.Module}}

import (
	"github.com/lulucas/hasura-pie/v1"
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
