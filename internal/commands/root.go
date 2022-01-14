package commands

import (
	"context"
	"github.com/oherych/yeelight"
	"github.com/oherych/yeelightcli/internal"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const (
	exampleDomain = "192.168.1.79:55443"
)

type clientBuilder func(command *cobra.Command, host string) internal.Interface
type discoveryFn func(ctx context.Context) (items []yeelight.DiscoveryResultItem, err error)

func Root(applicationName string, build clientBuilder, discovery discoveryFn) *cobra.Command {
	cobra.EnableCommandSorting = false

	command := &cobra.Command{
		Use:               applicationName,
		Short:             "CLI for manipulation Yeelight devises",
		Long:              "",
		SilenceUsage:      true,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	flags.InjectVerboseFlag(command)

	buildDocs(command)
	buildDiscovery(command, discovery)
	//buildGet(command, build)
	//buildSetName(command, build)
	//buildCron(command, build)
	//buildPowerOnOff(command, build, "on", "Turn on device", true)
	//buildPowerOnOff(command, build, "off", "Turn off device", false)
	//buildPowerToggle(command, build)
	//buildDevToggle(command, build)
	//buildSetRGB(command, build)
	//buildSetBright(command, build)
	//buildSetColorTemperature(command, build)
	//buildSetHSV(command, build)
	//buildDefault(command, build)

	return command
}

func buildDocs(command *cobra.Command) {
	const docFolder = "doc"

	command.AddCommand(&cobra.Command{
		Use:    "_doc",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doc.GenMarkdownTree(command, docFolder)
		},
	})
}
