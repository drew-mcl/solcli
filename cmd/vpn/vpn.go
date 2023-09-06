package vpn

import (
	"solcli/pkg/client"

	"github.com/spf13/cobra"
)

var VPNAPIPath = "/SEMP/v2/config/msgVpns"

func NewVPNCmd(c *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "Manage VPNs",
		Long: `Manage VPNs.
This command will allow you to manage VPNs on the Solace PubSub+ Event Broker.`,
		Args: cobra.NoArgs,
	}
	cmd.AddCommand(NewCreateCmd(c))
	cmd.AddCommand(NewListCmd(c))
	cmd.AddCommand(NewDeleteCmd(c))
	return cmd
}
