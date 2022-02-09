package helper

import (
	"github.com/spf13/cobra"
)

type Command interface {
	Use() string
	Short(cmd *cobra.Command) string
	Long(cmd *cobra.Command) string
	Flags(cmd *cobra.Command)
	Args() []Arg
}

type CommandExample interface {
	Example(cmd *cobra.Command) string
}

type CommandRun interface {
	Run(cmd *cobra.Command, args []string) error
}

type CommandSubCommand interface {
	SubCommand(cmd *cobra.Command) []Command
}
