package commands

import (
	"context"
	"errors"
	"github.com/oherych/yeelightcli/internal/helper"
	"strings"
	"time"

	"github.com/oherych/yeelight"
	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
)

const timeoutFlag = "timeout"

type DiscoveryCommand struct {
	discovery helper.DiscoveryFn
}

func (c DiscoveryCommand) Use() string {
	return "discovery"
}

func (c DiscoveryCommand) Short(cmd *cobra.Command) string {
	return "Find devices in local network"
}

func (c DiscoveryCommand) Long(cmd *cobra.Command) string {
	return "Application use SSDP protocol to search devices in the local network"
}

func (c DiscoveryCommand) Flags(cmd *cobra.Command) {
	cmd.Flags().Duration(timeoutFlag, 5*time.Second, "timeout")
}

func (c DiscoveryCommand) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (r DiscoveryCommand) Args() []helper.Arg {
	return nil
}

func (c DiscoveryCommand) Run(cmd *cobra.Command, args []string) error {
	timeout, err := cmd.Flags().GetDuration(timeoutFlag)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
	defer cancel()

	items, err := c.discovery(ctx)
	if err != nil && !(errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled)) {
		return err
	}

	return c.displayTable(cmd, items)
}

func (c DiscoveryCommand) displayTable(cmd *cobra.Command, items []yeelight.DiscoveryResultItem) error {
	table := tablewriter.NewWriter(cmd.OutOrStdout())
	table.SetHeader([]string{"Location", "Name", "ID", "Power", "Model/Version", "Support"})

	for _, device := range items {
		table.Append([]string{
			device.Location,
			device.Name,
			device.ID,
			c.powerStr(device.Power),
			device.Model + " [V:" + device.FirmwareVersion + "]",
			strings.Join(device.Support, ", "),
		})

	}

	table.Render()

	return nil
}

func (DiscoveryCommand) powerStr(in bool) string {
	if in {
		return "ON"
	}

	return "OFF"
}
