package main

import (
	"os"
	"solcli/cmd/queue"
	"solcli/cmd/vpn"
	"solcli/pkg/client"
	"solcli/pkg/loghandler"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var logLevels = map[string]slog.Level{
	"DEBUG":   slog.LevelDebug,
	"INFO":    slog.LevelInfo,
	"WARNING": slog.LevelWarn,
	"ERROR":   slog.LevelError,
	"D":       slog.LevelDebug,
	"I":       slog.LevelInfo,
	"W":       slog.LevelWarn,
	"E":       slog.LevelError,
}

var (
	c            *client.Client
	programLevel = new(slog.LevelVar)
	username     string
	password     string
	rawurl       string
	logLevel     string

	rootCmd = &cobra.Command{
		Use:   "solcli",
		Short: "solcli is a command line interface for the Solace PubSub+ Event Broker",
		Long: `solcli is a command line interface for the Solace PubSub+ Event Broker.
It is designed to be used in conjunction with the Solace PubSub+ Event Broker
Docker image, but can be used with any Solace PubSub+ Event Broker instance.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			h := loghandler.New(os.Stderr, &slog.HandlerOptions{Level: programLevel})
			slog.SetDefault(slog.New(h))

			//Validate solace authentication

			if logLevel != "" {
				levelStr := strings.ToUpper(logLevel)
				level, ok := logLevels[levelStr]
				if !ok {
					slog.Warn("Invalid log level specified, defaulting to Warn")
				} else {
					programLevel.Set(level)
				}
			}

			err := c.InitalizeClient(map[string]string{
				"Content-Type": "application/json",
			}, username, password, rawurl)
			if err != nil {
				slog.Error("Failed to initalize client!",
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			return nil
		},
	}
)

func init() {
	c = client.NewClient()
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username to authenticate with")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password to authenticate with")
	rootCmd.PersistentFlags().StringVarP(&rawurl, "url", "U", "", "URL to connect to")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "", "Log level to use")
	//Set name, password and url as required flags
}

func main() {
	rootCmd.AddCommand(
		vpn.NewVPNCmd(c),
		queue.NewQueueCmd(c),
	)
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Failed to execute command!",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
}
