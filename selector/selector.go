package selector

import "strings"

type Selector interface {
	GetPrefixes() []string
	HasSubselection() bool
	GetFragments() map[string]string
	Stringify() string
}

type SelectorSlice []Selector

func (slice *SelectorSlice) Stringify() string {
	if len(*slice) == 0 {
		return ""
	}
	var selections []string
	for _, selector := range *slice {
		selections = append(selections, selector.Stringify())
	}
	return "{" + strings.Join(selections, " ") + "}"
}

func (slice *SelectorSlice) GetPrefixes() []string {
	if len(*slice) == 0 {
		return nil
	}
	var result []string
	for _, selector := range *slice {
		currentPrefixes := selector.GetPrefixes()
		if len(currentPrefixes) == 0 {
			continue
		}
		result = append(result, currentPrefixes...)
	}
	return result
}

func (slice *SelectorSlice) GetFragments() map[string]string {
	result := make(map[string]string)
	for _, selector := range *slice {
		currentFragments := selector.GetFragments()
		if len(currentFragments) == 0 {
			continue
		}
		for name, fragment := range currentFragments {
			if _, ok := result[name]; !ok {
				result[name] = fragment
			}
		}
	}
	return result
}
