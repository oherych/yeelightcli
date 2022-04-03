package arguments

type NameArg struct{}

func (a NameArg) Name() string {
	return "name"
}

func (a NameArg) Description() string {
	return "The new name of a device."
}

func (a NameArg) Example() string {
	return "new_name"
}

func (a NameArg) Read(in string) (string, error) {
	return in, nil
}
