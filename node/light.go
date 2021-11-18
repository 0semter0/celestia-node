package node

import (
	"context"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"go.uber.org/fx"

	"github.com/celestiaorg/celestia-node/node/p2p"
	"github.com/celestiaorg/celestia-node/node/services"
)

// lightComponents keeps all the components as DI options required to built a Light Node.
func lightComponents(cfg *Config, repo Repository) fx.Option {
	return fx.Options(
		// manual providing
		fx.Provide(context.Background),
		fx.Provide(func() *Config {
			return cfg
		}),
		fx.Provide(func() ConfigLoader {
			return repo.Config
		}),
		fx.Provide(repo.Datastore),
		fx.Provide(repo.Keystore),
		// components
		p2p.Components(cfg.P2P),
		fx.Provide(services.Share),
		fx.Provide(func() tmbytes.HexBytes {
			return cfg.HeadHash
		}),
		fx.Provide(services.HeaderExchange),
		fx.Provide(services.HeaderStore),
		fx.Provide(services.Header),
		fx.Provide(services.LightAvailability), // TODO(@Wondertan): For FULL node this should be full availability
		fx.Provide(services.DASer),
	)
}
