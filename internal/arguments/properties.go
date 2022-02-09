package arguments

import "strings"

type Properties struct{}

func (a Properties) Name() string {
	return "properties"
}

func (a Properties) Example() string {
	return "power,rgb"
}

func (a Properties) Read(in string) ([]string, error) {
	properties := strings.Split(in, ",")
	for i := range properties {
		properties[i] = strings.TrimSpace(properties[i])
	}

	return properties, nil
}
