package arguments

type OffOn struct {
}

func (a OffOn) Name() string {
	return "off/on"
}

func (a OffOn) Example() string {
	return "off"
}

func (a OffOn) Read(in string) (bool, error) {
	if in == "off" {
		return false, nil
	}

	if in == "on" {
		return false, nil
	}

	return false, standardError(a)
}
