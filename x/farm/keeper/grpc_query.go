package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/v3/x/farm/types"
)

// Querier is used as Keeper will have duplicate methods if used directly,
// and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the lending module.
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.Keeper.paramSpace.GetParamSet(ctx, &params)
	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Querier) Plans(c context.Context, req *types.QueryPlansRequest) (*types.QueryPlansResponse, error) {
	// TODO: implement
	return &types.QueryPlansResponse{}, nil
}