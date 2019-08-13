package selector

type Fragment struct {
	name         string
	on           string
	subselection SelectorSlice
}

func NewFragment(name string, on string, subselection SelectorSlice) *Fragment {
	return &Fragment{name: name, on: on, subselection: subselection}
}

func (f *Fragment) WithSelection(selection ...Selector) *Fragment {
	f.subselection = selection
	return f
}

func (f *Fragment) Stringify() string {
	return "..." + f.name
}

func (f *Fragment) HasSubselection() bool {
	return len(f.subselection) > 0
}

func (f *Fragment) GetFragments() map[string]string {
	var fragment string
	fragment = "fragment " + f.name + " on " + f.on
	fragment += (&f.subselection).Stringify()

	result := f.subselection.GetFragments()
	result[f.name] = fragment

	return result
}

func (f *Fragment) GetPrefixes() []string {
	return f.subselection.GetPrefixes()
}
