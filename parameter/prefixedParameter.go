package parameter

type PrefixedParameter struct {
	name     string
	alias    string
	typeName string
}

func NewPrefixed(name string, alias string, typeName string) *PrefixedParameter {
	return &PrefixedParameter{name: name, alias: alias, typeName: typeName}
}

func (h *PrefixedParameter) GetString() string {
	return h.name + ": " + h.getAlias()
}

func (h *PrefixedParameter) getAlias() string {
	if h.alias == "" {
		return "$" + h.name
	}
	return "$" + h.alias
}

func (h *PrefixedParameter) GetPrefix() string {
	return h.getAlias() + ": " + h.typeName
}

func (h *PrefixedParameter) getPostfix() string {
	return h.name + ": " + h.getAlias()
}
