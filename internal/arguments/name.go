package arguments

type NameArg struct{}

func (a NameArg) Name() string {
	return "name"
}

func (a NameArg) Example() string {
	return "new_name"
}
