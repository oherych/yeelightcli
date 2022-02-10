package helper

import (
	"context"
	"github.com/oherych/yeelight"
	"github.com/oherych/yeelightcli/internal"
	"github.com/spf13/cobra"
)

func Client(command *cobra.Command, host string) internal.Interface {
	client := yeelight.New(host)

	return client
}

type ClientBuilder func(command *cobra.Command, host string) internal.Interface
type DiscoveryFn func(ctx context.Context) (items []yeelight.DiscoveryResultItem, err error)
