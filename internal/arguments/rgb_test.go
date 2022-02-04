package arguments

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropertiesRGB_Read(t *testing.T) {
	tests := map[string]struct {
		val int
		err error
	}{
		"FF0000":      {val: 16711680},
		"wrong_value": {err: helper.Error{Reason: "wrong value of rgb argument", Instruction: ""}},
	}

	for in, exp := range tests {
		t.Run(in, func(t *testing.T) {
			got, err := Rgb{}.Read(in)

			assert.Equal(t, exp.val, got)
			assert.Equal(t, exp.err, err)
		})
	}
}
