package helper

import (
	"github.com/oherych/yeelight"
	"github.com/oherych/yeelightcli/internal"
	"github.com/spf13/cobra"
)

func Client(command *cobra.Command, host string) internal.Interface {
	client := yeelight.New(host)

	return client
}
