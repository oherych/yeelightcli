package commands

import (
	"github.com/oherych/yeelightcli/internal/commands/cron"
	"github.com/oherych/yeelightcli/internal/commands/power"
	"github.com/oherych/yeelightcli/internal/commands/set"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	applicationName string
	build           helper.ClientBuilder
	discovery       helper.DiscoveryFn
}

func (r RootCommand) Args() []helper.Arg {
	return nil
}

func (r RootCommand) Use() string {
	return r.applicationName
}

func (r RootCommand) Short(_ *cobra.Command) string {
	return "CLI for manipulation Yeelight devises"
}

func (r RootCommand) Long(_ *cobra.Command) string {
	return ""
}

func (r RootCommand) Flags(cmd *cobra.Command) {
	flags.InjectVerboseFlag(cmd)
}

func (r RootCommand) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{
		DocsCommand{},
		DiscoveryCommand{discovery: r.discovery},
		power.Root{Build: r.build},
		set.Root{Build: r.build},
		DefaultCommand{build: r.build},
		cron.Root{Build: r.build},
		GetCommand{build: r.build},
	}
}
