package QueryBuilder

import (
	"gitlab.shop.by/beseller-go/gqlbuilder/selector"
	"strings"
)

type Builder struct {
	operationName string
	isMutation    bool
	selected      selector.SelectorSlice
}

func (builder *Builder) Mutation() *Builder {
	builder.isMutation = true
	return builder
}

func (builder *Builder) Query() *Builder {
	builder.isMutation = false
	return builder
}

func (builder *Builder) WithName(operationName string) *Builder {
	builder.operationName = operationName
	return builder
}

func (builder *Builder) Select(selected ...selector.Selector) *Builder {
	builder.selected = selected
	return builder
}

func (builder *Builder) Stringify() (string, error) {
	var query string

	if builder.isMutation {
		query = "mutation "
	} else {
		query = "query "
	}
	query += builder.operationName
	prefixes := builder.selected.GetPrefixes()
	if len(prefixes) > 0 {
		query += "(" + strings.Join(prefixes, ", ") + ")"
	}

	query += builder.selected.Stringify()

	fragments := builder.selected.GetFragments()
	slice := make([]string, 0, len(fragments))
	for _, value := range fragments {
		slice = append(slice, value)
	}
	query += " " + strings.Join(slice, " ")
	return query, nil
}
