package commands

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
)

type GetCommand struct {
	build helper.ClientBuilder
}

func (g GetCommand) Use() string {
	return "get [host] [parameter1,parameter2..]"
}

func (g GetCommand) Short(cmd *cobra.Command) string {
	return "Read device values"
}

func (g GetCommand) Long(cmd *cobra.Command) string {
	return ""
}

func (g GetCommand) Flags(cmd *cobra.Command) {

}

func (r GetCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.Properties{}}
}

func (g GetCommand) Run(cmd *cobra.Command, args []string) error {
	properties, err := arguments.Properties{}.Read(args[1])
	if err != nil {
		return err
	}

	result, err := g.build(cmd, args[0]).GetProperties(cmd.Context(), properties)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(cmd.OutOrStdout())
	table.SetHeader([]string{"Name", "Value"})

	for name, value := range result {
		table.Append([]string{name, value})
	}

	table.Render()

	return nil
}
