package queue

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
		Use:   "delete [queueName]",
		Short: "delete a queue",
		Long:  `Delete a message queue from a solace vpn`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("Deleting queue...")

			delPath := path.Join(QueueAPIPath, VPNName, "queues", args[0])
			slog.Debug("Attempting to delete queue", slog.Any("vpn", VPNName), slog.Any("url", delPath))

			ah := api.NewAPIHandler(c, delPath)

			_, err := ah.MsgVpnAPIRequest("DELETE", nil)
			if err != nil {
				slog.Error("Failed to delete Queue!",
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			slog.Info("Queue delete successfully!")
			return nil
		},
	}
	return cmd
}
