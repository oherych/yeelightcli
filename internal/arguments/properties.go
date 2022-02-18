package arguments

import (
	"github.com/oherych/yeelight"
	"strings"
)

type Properties struct{}

func (a Properties) Name() string {
	return "properties"
}

func (a Properties) Example() string {
	return yeelight.PropertyPower + "," + yeelight.PropertyRGB
}

func (a Properties) Description() string {
	return "Comma-separated list of properties. Can be used the next properties: " +
		strings.Join(yeelight.Properties(), ", ")
}

func (a Properties) Read(in string) ([]string, error) {
	properties := strings.Split(in, ",")
	for i := range properties {
		properties[i] = strings.TrimSpace(properties[i])
	}

	return properties, nil
}
