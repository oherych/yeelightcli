package commands

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

type DocsCommand struct{}

func (c DocsCommand) Use() string {
	return "_doc"
}

func (c DocsCommand) Short() string {
	return ""
}

func (c DocsCommand) Flags(cmd *cobra.Command) {
	cmd.Hidden = true
}

func (c DocsCommand) Args() []helper.Arg {
	return nil
}

func (c DocsCommand) Run(cmd *cobra.Command, args []string) error {
	const docFolder = "doc"

	return doc.GenMarkdownTree(cmd.Root(), docFolder)
}
