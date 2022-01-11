package main

import (
	"fmt"
	"github.com/oherych/yeelightcli/internal/helper"
	"os"
	"path/filepath"

	"github.com/oherych/yeelight"
	"github.com/oherych/yeelightcli/internal/commands"
)

func main() {
	if err := commands.Root(filepath.Base(os.Args[0]), helper.Client, yeelight.Discovery).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
