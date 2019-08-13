package selector

type Interface struct {
	name         string
	subselection SelectorSlice
}

func NewInterface(name string) *Interface {
	return &Interface{name: name}
}

func (f *Interface) WithSubselection(selection ...Selector) *Interface {
	f.subselection = selection
	return f
}

func (f *Interface) Stringify() string {
	var result string
	result = "... on " + f.name
	result += f.subselection.Stringify()
	return result
}

func (f *Interface) HasSubselection() bool {
	return len(f.subselection) > 0
}

func (f *Interface) GetPrefixes() []string {
	return f.subselection.GetPrefixes()
}

func (f *Interface) GetFragments() map[string]string {
	return f.subselection.GetFragments()
}
