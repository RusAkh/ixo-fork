package bonddoc

import (
	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey

	DefaultCodeSpace = types.DefaultCodeSpace
)

type (
	Keeper              = keeper.Keeper
	MsgCreateBond       = types.MsgCreateBond
	MsgUpdateBondStatus = types.MsgUpdateBondStatus
	StoredBondDoc       = types.StoredBondDoc
)

var (
	NewKeeper = keeper.NewKeeper
	ModuleCdc = types.ModuleCdc
)
