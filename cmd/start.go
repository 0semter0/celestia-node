package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Start constructs a CLI command to start Celestia Node daemon of any type with the given flags.
func Start(options ...func(*cobra.Command)) *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Short: `Starts Node daemon. First stopping signal gracefully stops the Node and second terminates it.
Options passed on start override configuration options only on start and are not persisted in config.`,
		Aliases:      []string{"daemon"},
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			ctx := cmd.Context()
			node, err := NewRunner(ctx)
			if err != nil {
				return err
			}

			go func() {
				ctx := node.Start(ctx)
				<-ctx.Done()
				// <-node.Stop(ctx)
				fmt.Println("Node stopped")
			}()

			err = <-node.Errors()
			return err
		},
	}
	// Apply each passed option to the command
	for _, option := range options {
		option(cmd)
	}
	return cmd
}
