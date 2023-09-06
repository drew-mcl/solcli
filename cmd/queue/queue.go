package queue

import (
	"solcli/pkg/client"

	"github.com/spf13/cobra"
)

var VPNName string
var QueueAPIPath = "/SEMP/v2/config/msgVpns"

func NewQueueCmd(c *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queue",
		Short: "Manage queues",
		Long:  `Manage solace queues inside message vpns.`,
		Args:  cobra.NoArgs,
	}
	cmd.AddCommand(NewCreateCmd(c))
	cmd.AddCommand(NewListCmd(c))
	cmd.AddCommand(NewDeleteCmd(c))
	cmd.Flags().StringVarP(&VPNName, "vpn", "v", "default", "The name of the message vpn for the queue operation to take places")
	return cmd
}
