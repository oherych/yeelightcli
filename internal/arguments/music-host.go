package arguments

import (
	"strconv"
	"strings"
)

type MusicHost struct{}

func (a MusicHost) Name() string {
	return "music-host"
}

func (a MusicHost) Example() string {
	// TODO:
	return ""
}

func (a MusicHost) Description() string {
	// TODO:
	return ""
}

func (a MusicHost) Read(in string) (string, int, error) {
	parts := strings.SplitN(in, ":", 2)
	if len(parts) != 2 {
		return "", 0, standardError(a)
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, standardError(a)
	}

	return parts[0], port, nil
}
