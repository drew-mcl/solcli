package queue

import (
	"os"
	"path"
	"solcli/pkg/api"
	"solcli/pkg/client"
	"solcli/pkg/printer"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func NewListCmd(c *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List queues",
		Long: `List queues.
This command will list all queues on the Solace PubSub+ Event Broker.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("Listing Queues", slog.Any("vpn", VPNName))

			qPath := path.Join(QueueAPIPath, VPNName, "queues")

			ah := api.NewAPIHandler(c, qPath)

			apiResp, err := ah.QueueAPIRequest("GET", nil)

			if err != nil {
				slog.Error("Failed to list Queues!",
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			slog.Info("Queues listed successfully")
			opts := printer.PrintOptions{
				OmitEmpty: true,
				AllFields: true,
			}
			printer.Print(apiResp, opts, 0)

			return nil
		},
	}
	return cmd
}
