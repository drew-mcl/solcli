package vpn

import (
	"bytes"
	"encoding/json"
	"os"
	"solcli/pkg/client"

	"solcli/pkg/api"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func NewCreateCmd(c *client.Client) *cobra.Command {

	opts := &api.MsgVpnObject{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a VPN",
		Long: `Create a VPN.
The VPN will be created with the following defaults:
	- Message VPN is enabled
	- Client username is "default"
	- Client password is "default"
	- Client profile is "default"	
	- Client ACL profile is "default"
	- Client authorization group is "default"
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("Creating VPN...")

			ah := api.NewAPIHandler(c, api.VPNAPIPath)

			opts.MsgVpnName = args[0]

			jsonOpts, err := json.Marshal(opts)
			if err != nil {
				slog.Error("Failed to marshal JSON!")
			}

			slog.Debug("Attempting to create VPN with the following parameters:",
				slog.String("jsonOpts", string(jsonOpts)),
			)

			vpn, err := ah.MsgVpnAPIRequest("POST", bytes.NewBuffer(jsonOpts))
			if err != nil {
				slog.Error("Failed to create VPN!",
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			slog.Info("VPN created successfully!",
				slog.Any("vpn", vpn),
			)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Enabled, "enabled", "e", true, "Enable the VPN")
	cmd.Flags().BoolVarP(&opts.SempOverMsgBusShowEnabled, "semp-over-msg-bus-show-enabled", "s", true, "Enable SEMP over message bus show commands")

	return cmd
}
