package arguments

type ColorTemperatureArg struct{}

func (a ColorTemperatureArg) Name() string {
	return "color temperature"
}

func (a ColorTemperatureArg) Example() string {
	return ""
}
