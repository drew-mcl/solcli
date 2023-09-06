package vpn

import (
	"os"
	"path"
	"solcli/pkg/client"

	"solcli/pkg/api"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func NewDeleteCmd(c *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [vpnName]",
		Short: "delete a VPN",
		Long:  `Delete a message VPN from solace`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("Deleting VPN...")

			delPath := path.Join(VPNAPIPath, args[0])
			slog.Debug("Attempting to delete vpn at", slog.Any("url", delPath))

			ah := api.NewAPIHandler(c, delPath)

			_, err := ah.MsgVpnAPIRequest("DELETE", nil)
			if err != nil {
				slog.Error("Failed to delete VPN!",
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			slog.Info("VPN delete successfully!")
			return nil
		},
	}
	return cmd
}
