package queue

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
	"solcli/pkg/client"
	"solcli/pkg/models"

	"solcli/pkg/api"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func NewCreateCmd(c *client.Client) *cobra.Command {

	/*
		"no-access" - Disallows all access.
		"read-only" - Read-only access to the messages.
		"consume" - Consume (read and remove) messages.
		"modify-topic" - Consume messages or modify the topic/selector.
		"delete" - Consume messages, modify the topic/selector or delete the Client created endpoint altogether.
	*/

	opts := models.QueueData{
		EgressEnabled:  true,
		IngressEnabled: true,
		MaxBindCount:   1,
		Permission:     "modify-topic",
	}

	cmd := &cobra.Command{
		Use:   "create [queueName]",
		Short: "Create a queue",
		Long: `Create a solace Queue within a messageVPN.
Unless specified with -v the queue  will be created in the 'default' vpn
	- Egress is enabled
	- Ingress is enabled
	- Max Bind Count is "1"
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("Creating Queue...")

			qPath := path.Join(QueueAPIPath, VPNName, "queues")

			opts.QueueName = args[0]

			ah := api.NewAPIHandler(c, qPath)

			opts.QueueName = args[0]

			jsonOpts, err := json.Marshal(opts)
			if err != nil {
				slog.Error("Failed to marshal JSON!")
			}

			slog.Debug("Attempting to create Queue with the following parameters:",
				slog.String("jsonOpts", string(jsonOpts)),
			)

			_, err = ah.QueueAPIRequest("POST", bytes.NewBuffer(jsonOpts))
			if err != nil {
				slog.Error("Failed to create Queue!",
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			slog.Info("Queue created successfully!")
			return nil
		},
	}
	return cmd
}
