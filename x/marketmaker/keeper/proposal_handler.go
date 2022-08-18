package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/crescent-network/crescent/v2/x/marketmaker/types"
)

// HandleMarketMakerProposal is a handler for executing a market maker proposal.
func HandleMarketMakerProposal(ctx sdk.Context, k Keeper, proposal *types.MarketMakerProposal) error {
	if proposal.Inclusions != nil {
		if err := k.InclusionMarketMakers(ctx, proposal.Inclusions); err != nil {
			return err
		}
	}

	if proposal.Distributions != nil {
		if err := k.DistributionMarketMakerIncentives(ctx, proposal.Distributions); err != nil {
			return err
		}
	}

	if proposal.Exclusions != nil {
		if err := k.ExclusionMarketMakers(ctx, proposal.Exclusions); err != nil {
			return err
		}
	}

	if proposal.Rejections != nil {
		if err := k.RejectionMarketMakers(ctx, proposal.Rejections); err != nil {
			return err
		}
	}

	return nil
}

// InclusionMarketMakers is a handler for include applied and not eligible market makers.
func (k Keeper) InclusionMarketMakers(ctx sdk.Context, proposals []types.MarketMakerHandle) error {
	for _, p := range proposals {
		mmAddr, err := sdk.AccAddressFromBech32(p.Address)
		if err != nil {
			return err
		}
		mm, found := k.GetMarketMaker(ctx, mmAddr, p.PairId)
		if !found {
			return sdkerrors.Wrapf(types.ErrNotExistMarketMaker, "%s is not a applied market maker", p.Address)
		}
		// pass already eligible market maker
		if mm.Eligible {
			return sdkerrors.Wrapf(types.ErrInvalidInclusion, "%s is already eligible market maker", p.Address)
		}
		mm.Eligible = true
		k.SetMarketMaker(ctx, mm)

		// refund Deposit amount
		deposit, found := k.GetDeposit(ctx, mmAddr, p.PairId)
		if !found {
			return types.ErrInvalidDeposit
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, mmAddr, deposit.Amount)
		if err != nil {
			return err
		}
		k.DeleteDeposit(ctx, mmAddr, p.PairId)

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeIncludeMarketMaker,
				sdk.NewAttribute(types.AttributeKeyAddress, p.Address),
				sdk.NewAttribute(types.AttributeKeyPairId, fmt.Sprintf("%d", p.PairId)),
			),
		})
	}
	return nil
}

// ExclusionMarketMakers is a handler for exclude eligible market makers.
func (k Keeper) ExclusionMarketMakers(ctx sdk.Context, proposals []types.MarketMakerHandle) error {
	for _, p := range proposals {
		mmAddr, err := sdk.AccAddressFromBech32(p.Address)
		if err != nil {
			return err
		}
		mm, found := k.GetMarketMaker(ctx, mmAddr, p.PairId)
		if !found {
			return sdkerrors.Wrapf(types.ErrNotExistMarketMaker, "%s is not market maker", p.Address)
		}

		if !mm.Eligible {
			return sdkerrors.Wrapf(types.ErrInvalidExclusion, "%s is not eligible market maker", p.Address)
		}

		k.DeleteMarketMaker(ctx, mmAddr, p.PairId)

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeExcludeMarketMaker,
				sdk.NewAttribute(types.AttributeKeyAddress, p.Address),
				sdk.NewAttribute(types.AttributeKeyPairId, fmt.Sprintf("%d", p.PairId)),
			),
		})
	}
	return nil
}

// RejectionMarketMakers is a handler for reject applied and not eligible market makers.
func (k Keeper) RejectionMarketMakers(ctx sdk.Context, proposals []types.MarketMakerHandle) error {
	for _, p := range proposals {
		mmAddr, err := sdk.AccAddressFromBech32(p.Address)
		if err != nil {
			return err
		}

		mm, found := k.GetMarketMaker(ctx, mmAddr, p.PairId)
		if !found {
			return sdkerrors.Wrapf(types.ErrNotExistMarketMaker, "%s is not market maker", p.Address)
		}

		if mm.Eligible {
			return sdkerrors.Wrapf(types.ErrInvalidRejection, "%s is eligible market maker", p.Address)
		}

		k.DeleteMarketMaker(ctx, mmAddr, p.PairId)

		// refund DepositAmount
		deposit, found := k.GetDeposit(ctx, mmAddr, p.PairId)
		if !found {
			return types.ErrInvalidDeposit
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, mmAddr, deposit.Amount)
		if err != nil {
			return err
		}
		k.DeleteDeposit(ctx, mmAddr, p.PairId)

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeRejectMarketMaker,
				sdk.NewAttribute(types.AttributeKeyAddress, p.Address),
				sdk.NewAttribute(types.AttributeKeyPairId, fmt.Sprintf("%d", p.PairId)),
			),
		})
	}
	return nil
}

// DistributionMarketMakerIncentives is a handler for distribute incentives to eligible market makers.
func (k Keeper) DistributionMarketMakerIncentives(ctx sdk.Context, proposals []types.IncentiveDistribution) error {
	params := k.GetParams(ctx)
	totalIncentives := sdk.Coins{}
	for _, p := range proposals {
		totalIncentives = totalIncentives.Add(p.Amount...)

		mm, found := k.GetMarketMaker(ctx, p.GetAccAddress(), p.PairId)
		if !found {
			return types.ErrNotExistMarketMaker
		}
		if !mm.Eligible {
			return types.ErrNotEligibleMarketMaker
		}

	}

	budgetAcc := params.IncentiveBudgetAcc()
	err := k.bankKeeper.SendCoins(ctx, budgetAcc, types.ClaimableIncentiveReserveAcc, totalIncentives)
	if err != nil {
		return err
	}

	for _, p := range proposals {
		var amount sdk.Coins
		// add and set if already claimable incentive exist
		incentive, found := k.GetIncentive(ctx, p.GetAccAddress())
		if found {
			amount = incentive.Claimable.Add(p.Amount...)
		} else {
			amount = p.Amount
		}
		k.SetIncentive(ctx, types.Incentive{
			Address:   p.Address,
			Claimable: amount,
		})
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDistributeIncentives,
			sdk.NewAttribute(types.AttributeKeyBudgetAddress, budgetAcc.String()),
			sdk.NewAttribute(types.AttributeKeyTotalIncentives, totalIncentives.String()),
		),
	})
	return nil
}