package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	cmdnode "github.com/celestiaorg/celestia-node/cmd"
	"github.com/celestiaorg/celestia-node/nodebuilder/core"
	"github.com/celestiaorg/celestia-node/nodebuilder/gateway"
	"github.com/celestiaorg/celestia-node/nodebuilder/header"
	"github.com/celestiaorg/celestia-node/nodebuilder/node"
	"github.com/celestiaorg/celestia-node/nodebuilder/p2p"
	"github.com/celestiaorg/celestia-node/nodebuilder/rpc"
	"github.com/celestiaorg/celestia-node/nodebuilder/state"
)

// NOTE: We should always ensure that the added Flags below are parsed somewhere, like in the
// PersistentPreRun func on parent command.

func init() {
	flags := []*pflag.FlagSet{
		cmdnode.NodeFlags(),
		p2p.Flags(),
		header.Flags(),
		cmdnode.MiscFlags(),
		// NOTE: for now, state-related queries can only be accessed
		// over an RPC connection with a celestia-core node.
		core.Flags(),
		rpc.Flags(),
		gateway.Flags(),
		state.Flags(),
	}

	fullCmd.AddCommand(
		cmdnode.Init(flags...),
		cmdnode.Remove(node.Full, cmdnode.NodeFlags(), p2p.Flags()),
		cmdnode.Reinit(node.Full, cmdnode.NodeFlags(), p2p.Flags()),
		cmdnode.Start(flags...),
		cmdnode.AuthCmd(flags...),
	)
}

var fullCmd = &cobra.Command{
	Use:   "full [subcommand]",
	Args:  cobra.NoArgs,
	Short: "Manage your Full node",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRunEnv(cmd, node.Full, args)
	},
}
