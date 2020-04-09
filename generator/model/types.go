package model

import (
	"fmt"
	"github.com/lulucas/hasura-pie-cli/utils"
	"log"
	"strings"
)

const (
	structTpl = `type %s struct {
%s}

`
)

var (
	//key - pg type, value - go type
	nonNullableTypes = map[string]string{
		"int2":        "int16",
		"int4":        "int32",
		"int8":        "int64",
		"float4":      "float32",
		"float8":      "float64",
		"numeric":     "decimal.Decimal",
		"money":       "float64",
		"bpchar":      "string",
		"varchar":     "string",
		"text":        "string",
		"bytea":       "[]byte",
		"uuid":        "uuid.UUID",
		"timestamp":   "time.Time",
		"timestamptz": "time.Time",
		"time":        "time.Time",
		"timetz":      "time.Time",
		"date":        "time.Time",
		"interval":    "time.Time",
		"bool":        "bool",
		"bit":         "uint32",
		"varbit":      "uint32",
		"jsonb":       "json.RawMessage",
		"xml":         "json.RawMessage",
	}

	nullableTypes = map[string]string{
		"int2":        "*int",
		"int4":        "*int",
		"int8":        "*int",
		"float4":      "*float64",
		"float8":      "*float64",
		"numeric":     "*decimal.Decimal",
		"money":       "*float64",
		"bpchar":      "null.String",
		"varchar":     "null.String",
		"text":        "*string",
		"bytea":       "*[]byte",
		"uuid":        "*uuid.UUID",
		"timestamp":   "*time.Time",
		"timestamptz": "*time.Time",
		"time":        "*time.Time",
		"timetz":      "*time.Time",
		"date":        "*time.Time",
		"interval":    "*time.Time",
		"bool":        "*bool",
		"bit":         "*uint32",
		"varbit":      "*uint32",
		"jsonb":       "*json.RawMessage",
		"xml":         "*xml.RawMessage",
	}
)

func getStruct(tab string, cols []*column) string {
	var body string
	for _, c := range cols {
		d := ""
		if c.Default != nil {
			d = *c.Default
		}
		body += fmt.Sprintf("\t%s %s // default: %s\n",
			utils.Snake2Camel(c.Name), convertType(c), d)
	}

	return fmt.Sprintf(structTpl, strings.TrimSuffix(utils.Snake2Camel(tab), "s"), body)
}

func getMap(n bool) map[string]string {
	if n {
		return nullableTypes
	}
	return nonNullableTypes
}

func convertType(c *column) string {
	t, ok := getMap(c.IsNull)[c.Type]
	if !ok {
		log.Fatalln("unknown type: ", t, c.Type)
	}

	if c.IsArray {
		return "[]" + t
	}
	return t
}
