package bonddoc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type InternalAccountID = string

func NewHandler(k Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateBond:
			return handleMsgCreateBond(ctx, k, msg)
		case MsgUpdateBondStatus:
			return handleMsgUpdateBondStatus(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleMsgCreateBond(ctx sdk.Context, k Keeper, msg MsgCreateBond) sdk.Result {

	err := k.SetBondDoc(ctx, &msg)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func handleMsgUpdateBondStatus(ctx sdk.Context, k Keeper, msg MsgUpdateBondStatus) sdk.Result {

	ExistingBondDoc, err := getBondDoc(ctx, k, msg.GetBondDid())
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Bond").Result()
	}

	newStatus := msg.GetStatus()
	if !newStatus.IsValidProgressionFrom(ExistingBondDoc.GetStatus()) {
		return sdk.ErrUnknownRequest("Invalid Status Progression requested").Result()
	}

	// TODO: actions depending on new status (refer to how projects module does this)

	ExistingBondDoc.SetStatus(newStatus)
	_, _ = k.UpdateBondDoc(ctx, ExistingBondDoc)

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func getBondDoc(ctx sdk.Context, k Keeper, bondDid ixo.Did) (StoredBondDoc, sdk.Error) {
	ixoBondDoc, err := k.GetBondDoc(ctx, bondDid)

	return ixoBondDoc.(StoredBondDoc), err
}
