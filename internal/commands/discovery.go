package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/oherych/yeelightcli/internal/helper"

	"github.com/spf13/cobra"
)

func buildDiscovery(parent *cobra.Command, discovery discoveryFn) {
	helper.BuildCommand(parent, "discovery", func(cmd *cobra.Command) {
		const timeoutFlag = "timeout"

		cmd.Short = "Find devices in local network"
		cmd.Long = "Application use SSDP protocol to search devices in the local network"

		cmd.Flags().Duration(timeoutFlag, 5*time.Second, "timeout")

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			timeout, err := cmd.Flags().GetDuration(timeoutFlag)
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
			defer cancel()

			items, err := discovery(ctx)
			if err != nil && !(errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled)) {
				return err
			}

			fmt.Println(items)

			return nil
		}
	})
}
