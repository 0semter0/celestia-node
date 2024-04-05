package state

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	libfraud "github.com/celestiaorg/go-fraud"
	"github.com/celestiaorg/go-header/sync"

	"github.com/celestiaorg/celestia-node/header"
	"github.com/celestiaorg/celestia-node/nodebuilder/core"
	modfraud "github.com/celestiaorg/celestia-node/nodebuilder/fraud"
	"github.com/celestiaorg/celestia-node/share/eds/byzantine"
	"github.com/celestiaorg/celestia-node/state"
)

// coreAccessor constructs a new instance of state.Module over
// a celestia-core connection.
func coreAccessor(
	corecfg core.Config,
	keyring keyring.Keyring,
	keyname string,
	sync *sync.Syncer[*header.ExtendedHeader],
	fraudServ libfraud.Service[*header.ExtendedHeader],
	opts []state.Option,
) (*state.CoreAccessor, Module, *modfraud.ServiceBreaker[*state.CoreAccessor, *header.ExtendedHeader]) {
	ca := state.NewCoreAccessor(keyring, keyname, sync, corecfg.IP, corecfg.RPCPort, corecfg.GRPCPort, opts...)
	sBreaker := &modfraud.ServiceBreaker[*state.CoreAccessor, *header.ExtendedHeader]{
		Service:   ca,
		FraudType: byzantine.BadEncoding,
		FraudServ: fraudServ,
	}
	return ca, ca, sBreaker
}
