package selector

import (
	"gitlab.shop.by/beseller-go/gqlbuilder/parameter"
)

type Field struct {
	name         string
	alias        string
	params       parameter.Container
	subselection SelectorSlice
}

func NewField(name string) *Field {
	return &Field{name: name}
}

func (f *Field) Alias(alias string) *Field {
	f.alias = alias
	return f
}

func (f *Field) WithParams(params ...parameter.Parameter) *Field {
	f.params = params
	return f
}

func (f *Field) WithSubselection(selection ...Selector) *Field {
	f.subselection = selection
	return f
}

func (f *Field) Stringify() string {
	var result string
	if f.alias != "" {
		result += f.alias + ": "
	}
	result += f.name
	if f.params != nil {
		result += f.params.GetString()
	}
	result += f.subselection.Stringify()
	return result
}

func (f *Field) HasSubselection() bool {
	return len(f.subselection) > 0
}

func (f *Field) GetPrefixes() []string {
	result := f.params.GetPrefixes()
	subPrefixes := f.subselection.GetPrefixes()
	if len(subPrefixes) > 0 {
		result = append(result, subPrefixes...)
	}
	return result
}

func (f *Field) GetFragments() map[string]string {
	return f.subselection.GetFragments()
}
