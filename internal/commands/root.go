package commands

import (
	"context"

	"github.com/oherych/yeelight"
	"github.com/oherych/yeelightcli/internal"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/spf13/cobra"
)

const (
	exampleDomain = "192.168.1.79:55443"
)

type clientBuilder func(command *cobra.Command, host string) internal.Interface
type discoveryFn func(ctx context.Context) (items []yeelight.DiscoveryResultItem, err error)

func Root(applicationName string, build clientBuilder, discovery discoveryFn) *cobra.Command {
	command := &cobra.Command{
		Use:               applicationName,
		Short:             "CLI for manipulation Yeelight devises",
		SilenceUsage:      true,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	flags.InjectVerboseFlag(command)

	buildGet(command, build)
	buildDiscovery(command, discovery)
	buildSetColor(command, build)
	buildSetBright(command, build)
	buildSetName(command, build)
	buildPower(command, build)
	buildCron(command, build)

	command.AddCommand(&cobra.Command{Use: "set_default"})
	command.AddCommand(&cobra.Command{Use: "set_color_temperature"}) // set_ct_abx
	command.AddCommand(&cobra.Command{Use: "set_hsv"})               // set_hsv
	command.AddCommand(&cobra.Command{Use: "set_adjust"})            // set_adjust
	command.AddCommand(&cobra.Command{Use: "set_music"})             // set_music

	return command
}
