package helper

import (
	"github.com/spf13/cobra"
)

type Command interface {
	Use() string
	Short() string
	Flags(cmd *cobra.Command)
	Args() []Arg
}

type CommandLong interface {
	Long() string
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
