package utils

import (
	"github.com/titanchain/titanchain/eth"
	"github.com/titanchain/titanchain/eth/downloader"
	"github.com/titanchain/titanchain/ethstats"
	"github.com/titanchain/titanchain/les"
	"github.com/titanchain/titanchain/node"
	"github.com/titanchain/titanchain/titanx"
	"github.com/titanchain/titanchain/titanxlending"
	whisper "github.com/titanchain/titanchain/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var titanXServ *titanx.TitanX
			ctx.Service(&titanXServ)
			var lendingServ *titanxlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, titanXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterTitanXService(stack *node.Node, cfg *titanx.Config) {
	titanX := titanx.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return titanX, nil
	}); err != nil {
		Fatalf("Failed to register the TitanX service: %v", err)
	}

	// register titanxlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return titanxlending.New(titanX), nil
	}); err != nil {
		Fatalf("Failed to register the TitanXLending service: %v", err)
	}
}
