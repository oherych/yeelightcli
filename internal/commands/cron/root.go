package cron

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Root struct {
	Build helper.ClientBuilder
}

func (c Root) Use() string {
	return "cron"
}

func (c Root) Short() string {
	// TODO: check
	return "Schedule light turn on/off"
}

func (c Root) Flags(cmd *cobra.Command) {
}

func (c Root) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{
		Get{build: c.Build},
		Add{build: c.Build},
		Delete{build: c.Build},
	}
}

func (c Root) Args() []helper.Arg {
	return nil
}
