package main

import (
	"github.com/spf13/cobra"

	cmdnode "github.com/celestiaorg/celestia-node/cmd"
	"github.com/celestiaorg/celestia-node/nodebuilder/node"
)

// NOTE: We should always ensure that the added Flags below are parsed somewhere, like in the PersistentPreRun func on
// parent command.

func init() {
	bridgeCmd.AddCommand(
		cmdnode.Init(
			cmdnode.NodeFlags(),
			cmdnode.P2PFlags(),
			cmdnode.CoreFlags(),
			cmdnode.MiscFlags(),
			cmdnode.RPCFlags(),
			cmdnode.KeyFlags(),
		),
		cmdnode.Start(
			cmdnode.NodeFlags(),
			cmdnode.P2PFlags(),
			cmdnode.CoreFlags(),
			cmdnode.MiscFlags(),
			cmdnode.RPCFlags(),
			cmdnode.KeyFlags(),
		),
	)
}

var bridgeCmd = &cobra.Command{
	Use:   "bridge [subcommand]",
	Args:  cobra.NoArgs,
	Short: "Manage your Bridge node",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var (
			ctx = cmd.Context()
			err error
		)

		ctx = cmdnode.WithNodeType(ctx, node.Bridge)

		ctx, err = cmdnode.ParseNodeFlags(ctx, cmd)
		if err != nil {
			return err
		}

		cfg := cmdnode.NodeConfig(ctx)

		err = cmdnode.ParseP2PFlags(cmd, &cfg)
		if err != nil {
			return err
		}

		err = cmdnode.ParseCoreFlags(cmd, &cfg)
		if err != nil {
			return err
		}

		ctx, err = cmdnode.ParseMiscFlags(ctx, cmd)
		if err != nil {
			return err
		}

		cmdnode.ParseRPCFlags(cmd, &cfg)
		cmdnode.ParseKeyFlags(cmd, &cfg)

		// set config
		ctx = cmdnode.WithNodeConfig(ctx, &cfg)
		cmd.SetContext(ctx)
		return nil
	},
}
