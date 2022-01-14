package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/oherych/yeelight"
	"github.com/oherych/yeelightcli/internal/commands"
	"github.com/oherych/yeelightcli/internal/helper"
)

func main() {
	root := commands.Root(filepath.Base(os.Args[0]), helper.Client, yeelight.Discovery)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
