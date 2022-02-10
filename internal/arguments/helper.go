package arguments

import "github.com/oherych/yeelightcli/internal/helper"

func standardError(a helper.Arg) error {
	return helper.Error{Reason: "wrong value of " + a.Name() + " argument"}
}
