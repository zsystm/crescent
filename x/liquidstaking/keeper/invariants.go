package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmosquad-labs/squad/x/liquidstaking/types"
)

// RegisterInvariants registers all liquidstaking invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "net-amount",
		NetAmountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "total-liquid-tokens",
		TotalLiquidTokensInvariant(k))
	ir.RegisterRoute(types.ModuleName, "liquid-delegation",
		LiquidDelegationInvariant(k))
}

// AllInvariants runs all invariants of the liquidstaking module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		for _, inv := range []func(Keeper) sdk.Invariant{
			NetAmountInvariant,
			TotalLiquidTokensInvariant,
			LiquidDelegationInvariant,
		} {
			res, stop := inv(k)(ctx)
			if stop {
				return res, stop
			}
		}
		return "", false
	}
}

// NetAmountInvariant checks that the amount of btoken supply with NetAmount.
func NetAmountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		msg := ""
		broken := false
		lvs := k.GetAllLiquidValidators(ctx)
		if lvs.Len() == 0 {
			return msg, broken
		}
		NetAmount := k.NetAmount(ctx)
		balance := k.bankKeeper.GetBalance(ctx, types.LiquidStakingProxyAcc, k.stakingKeeper.BondDenom(ctx)).Amount
		NetAmountExceptBalance := NetAmount.Sub(balance.ToDec())
		bondedBondDenom := k.BondedBondDenom(ctx)
		bTokenTotalSupply := k.bankKeeper.GetSupply(ctx, bondedBondDenom)
		if bTokenTotalSupply.IsPositive() && !NetAmountExceptBalance.IsPositive() {
			msg = "found positive btoken supply with non-positive net amount"
			broken = true
		}
		if !bTokenTotalSupply.IsPositive() && NetAmountExceptBalance.IsPositive() {
			msg = "found positive net amount with non-positive btoken supply"
			broken = true
		}
		return sdk.FormatInvariant(
			types.ModuleName, "btoken supply with net amount invariant broken",
			msg,
		), broken
	}
}

// TotalLiquidTokensInvariant checks equal total liquid tokens of proxy account with total liquid tokens of liquid validators.
func TotalLiquidTokensInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		lvs := k.GetAllLiquidValidators(ctx)
		_, _, totalLiquidTokensOfProxyAcc := k.CheckTotalRewards(ctx, types.LiquidStakingProxyAcc)
		totalLiquidTokensOfLiquidValidators := lvs.TotalLiquidTokens(ctx, k.stakingKeeper)

		// TODO: check delShares to tokens convert decimal error, consider to check TotalDelShares
		broken := !totalLiquidTokensOfProxyAcc.Equal(totalLiquidTokensOfLiquidValidators)
		return sdk.FormatInvariant(
			types.ModuleName, "total liquid tokens invariant broken",
			fmt.Sprintf("found unmatched total liquid tokens of proxy account %s with total liquid tokens of liquid validators %s\n",
				totalLiquidTokensOfProxyAcc.String(), totalLiquidTokensOfLiquidValidators.String()),
		), broken
	}
}

// LiquidDelegationInvariant checks all delegation of proxy account involved liquid validators.
func LiquidDelegationInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		msg := ""
		count := 0
		liquidValidatorMap := k.GetAllLiquidValidators(ctx).Map()

		// remove delegation condition -> Unbond(slash, undelegate, redelegate)
		// remove validator condition -> Unbond(slash, undelegate, redelegate), UnbondAllMatureValidators(BlockValidatorUpdates on staking endblock)
		k.stakingKeeper.IterateDelegations(
			ctx, types.LiquidStakingProxyAcc,
			func(_ int64, del stakingtypes.DelegationI) (stop bool) {
				delAddr := del.GetValidatorAddr().String()
				if _, ok := liquidValidatorMap[delAddr]; !ok {
					msg += fmt.Sprintf("\t%s has delegation but not liquid validator\n", delAddr)
					count++
				}
				return false
			},
		)

		broken := count != 0
		return sdk.FormatInvariant(
			types.ModuleName, "total liquid tokens invariant broken",
			fmt.Sprintf("found %d delegation of proxy account for not liquid validators\n%s", count, msg),
		), broken
	}
}