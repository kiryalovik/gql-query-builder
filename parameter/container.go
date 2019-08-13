package parameter

import "strings"

type Container []Parameter

func (c *Container) GetString() string {
	if len(*c) == 0 {
		return ""
	}
	var slice []string
	for _, param := range *c {
		slice = append(slice, param.GetString())
	}
	return "(" + strings.Join(slice, ", ") + ")"
}

func (c *Container) GetPrefixes() []string {
	var slice []string
	for _, param := range *c {
		prefix := param.GetPrefix()
		if prefix != "" {
			slice = append(slice, prefix)
		}
	}
	return slice
}
