package model

import (
	"fmt"
	"sort"
	"strings"
)

const (
	importTpl = `
import (
%s
)

`
)

var (
	packageImportByTypes = map[string]string{
		"numeric":     "github.com/shopspring/decimal",
		"uuid":        "github.com/satori/go.uuid",
		"timestamp":   "time",
		"timestamptz": "time",
		"time":        "time",
		"timetz":      "time",
		"date":        "time",
		"interval":    "time",
		"jsonb":       "encoding/json",
	}
)

func getImports(cols []*column) string {
	imports := map[string]bool{}
	for _, c := range cols {
		if i, ok := packageImportByTypes[c.Type]; ok {
			imports[fmt.Sprintf("\t\"%s\"", i)] = true
		}
	}

	var keys []string
	for k, _ := range imports {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return fmt.Sprintf(importTpl, strings.Join(keys, "\n"))
}
