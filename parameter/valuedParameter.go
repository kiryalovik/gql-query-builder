package parameter

import "encoding/json"

type ValuedParameter struct {
	name  string
	value interface{}
}

func NewValued(name string, value interface{}) *ValuedParameter {
	return &ValuedParameter{name: name, value: value}
}

func (p *ValuedParameter) GetString() string {
	var result string
	value, _ := json.Marshal(p.value)
	result = p.name + ": " + string(value)
	return result
}

func (p *ValuedParameter) GetPrefix() string {
	return ""
}
