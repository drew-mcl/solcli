package vpn

import (
	"os"
	"solcli/pkg/api"
	"solcli/pkg/client"
	"solcli/pkg/printer"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func NewListCmd(c *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List VPNs",
		Long: `List VPNs.
This command will list all VPNs on the Solace PubSub+ Event Broker.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("Listing VPNs...")

			ah := api.NewAPIHandler(c, api.VPNAPIPath)

			apiResp, err := ah.MsgVpnAPIRequest("GET", nil)

			if err != nil {
				slog.Error("Failed to list VPNs!",
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			slog.Info("VPNs listed successfully")
			opts := printer.PrintOptions{
				OmitEmpty: false,
				AllFields: true,
			}
			printer.Print(apiResp, opts)

			return nil
		},
	}
	return cmd
}
