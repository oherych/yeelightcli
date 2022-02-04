package arguments

type HsvArg struct{}

func (a HsvArg) Name() string {
	return "hsv"
}

func (a HsvArg) Example() string {
	return ""
}
