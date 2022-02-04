package arguments

type Properties struct{}

func (a Properties) Name() string {
	return "properties"
}

func (a Properties) Example() string {
	return "power,rgb"
}
