package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmosquad-labs/squad/x/liquidity/types"
)

func TestMsgCreatePair(t *testing.T) {
	for _, tc := range []struct {
		malleate    func(msg *types.MsgCreatePair)
		expectedErr string
	}{
		{
			func(msg *types.MsgCreatePair) {},
			"",
		},
		{
			func(msg *types.MsgCreatePair) {
				msg.Creator = "invalidaddr"
			},
			"invalid creator address: decoding bech32 failed: invalid separator index -1: invalid address",
		},
		{
			func(msg *types.MsgCreatePair) {
				msg.BaseCoinDenom = "invaliddenom!"
			},
			"invalid denom: invaliddenom!: invalid request",
		},
		{
			func(msg *types.MsgCreatePair) {
				msg.QuoteCoinDenom = "invaliddenom!"
			},
			"invalid denom: invaliddenom!: invalid request",
		},
	} {
		t.Run("", func(t *testing.T) {
			msg := types.NewMsgCreatePair(sdk.AccAddress(crypto.AddressHash([]byte("creator"))), "denom1", "denom2")
			tc.malleate(msg)
			require.Equal(t, types.TypeMsgCreatePair, msg.Type())
			require.Equal(t, types.RouterKey, msg.Route())
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
				signers := msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, msg.GetCreator(), signers[0])
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestMsgCreatePool(t *testing.T) {
	testCases := []struct {
		expErr string
		msg    *types.MsgCreatePool
	}{
		{
			"", // empty means no error expected
			types.NewMsgCreatePool(
				sdk.AccAddress(crypto.AddressHash([]byte("Creator"))),
				1,
				sdk.NewCoins(sdk.NewInt64Coin("denom1", 1000000), sdk.NewInt64Coin("denom2", 1000000)),
			),
		},
		{
			"invalid creator address: empty address string is not allowed: invalid address",
			types.NewMsgCreatePool(
				sdk.AccAddress{},
				1,
				sdk.NewCoins(sdk.NewInt64Coin("denom1", 1000000), sdk.NewInt64Coin("denom2", 1000000)),
			),
		},
		{
			"coin 0denom1 amount is not positive",
			types.NewMsgCreatePool(
				sdk.AccAddress(crypto.AddressHash([]byte("Creator"))),
				1,
				sdk.Coins{sdk.NewInt64Coin("denom1", 0), sdk.NewInt64Coin("denom2", 1000000)},
			),
		},
		{
			"coin denom2 amount is not positive",
			types.NewMsgCreatePool(
				sdk.AccAddress(crypto.AddressHash([]byte("Creator"))),
				1,
				sdk.Coins{sdk.NewInt64Coin("denom1", 1000000), sdk.NewInt64Coin("denom2", 0)},
			),
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, types.TypeMsgCreatePool, tc.msg.Type())
			require.Equal(t, types.RouterKey, tc.msg.Route())

			err := tc.msg.ValidateBasic()
			if tc.expErr == "" {
				require.NoError(t, err)
				signers := tc.msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, tc.msg.GetCreator(), signers[0])
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestMsgDepositBatch(t *testing.T) {
	testCases := []struct {
		expErr string
		msg    *types.MsgDepositBatch
	}{
		{
			"", // empty means no error expected
			types.NewMsgDepositBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Depositor"))),
				1,
				sdk.NewCoins(sdk.NewInt64Coin("denom1", 1000000), sdk.NewInt64Coin("denom2", 1000000)),
			),
		},
		{
			"invalid depositor address: empty address string is not allowed: invalid address",
			types.NewMsgDepositBatch(
				sdk.AccAddress{},
				1,
				sdk.NewCoins(sdk.NewInt64Coin("denom1", 1000000), sdk.NewInt64Coin("denom2", 1000000)),
			),
		},
		{
			"coin 0denom1 amount is not positive",
			types.NewMsgDepositBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Depositor"))),
				1,
				sdk.Coins{sdk.NewInt64Coin("denom1", 0), sdk.NewInt64Coin("denom2", 1000000)},
			),
		},
		{
			"coin denom2 amount is not positive",
			types.NewMsgDepositBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Depositor"))),
				1,
				sdk.Coins{sdk.NewInt64Coin("denom1", 1000000), sdk.NewInt64Coin("denom2", 0)},
			),
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, types.TypeMsgDepositBatch, tc.msg.Type())
			require.Equal(t, types.RouterKey, tc.msg.Route())

			err := tc.msg.ValidateBasic()
			if tc.expErr == "" {
				require.NoError(t, err)
				signers := tc.msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, tc.msg.GetDepositor(), signers[0])
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestMsgWithdrawBatch(t *testing.T) {
	testCases := []struct {
		expErr string
		msg    *types.MsgWithdrawBatch
	}{
		{
			"", // empty means no error expected
			types.NewMsgWithdrawBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Withdrawer"))),
				1,
				sdk.NewInt64Coin("PoolCoinDenom", 500_000),
			),
		},
		{
			"invalid withdrawer address: empty address string is not allowed: invalid address",
			types.NewMsgWithdrawBatch(
				sdk.AccAddress{},
				1,
				sdk.NewInt64Coin("PoolCoinDenom", 500_000),
			),
		},
		{
			"pool coin must be positive: invalid request",
			types.NewMsgWithdrawBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Withdrawer"))),
				1,
				sdk.NewInt64Coin("PoolCoinDenom", 0),
			),
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, types.TypeMsgWithdrawBatch, tc.msg.Type())
			require.Equal(t, types.RouterKey, tc.msg.Route())

			err := tc.msg.ValidateBasic()
			if tc.expErr == "" {
				require.NoError(t, err)
				signers := tc.msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, tc.msg.GetWithdrawer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestMsgMarketOrderBatch(t *testing.T) {
	orderLifespan := 20 * time.Second

	testCases := []struct {
		expErr string
		msg    *types.MsgMarketOrderBatch
	}{
		{
			"", // empty means no error expected
			types.NewMsgMarketOrderBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Orderer"))),
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 100_000_000),
				"denom2",
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
		{
			"invalid orderer address: empty address string is not allowed: invalid address",
			types.NewMsgMarketOrderBatch(
				sdk.AccAddress{},
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 100_000_000),
				"denom2",
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
		{
			"offer coin must be positive: invalid request",
			types.NewMsgMarketOrderBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Orderer"))),
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 0),
				"denom2",
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
		{
			"offer coin denom and demand coin denom must not be same: invalid request",
			types.NewMsgMarketOrderBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Orderer"))),
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 100_000_000),
				"denom1",
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, types.TypeMsgMarketOrderBatch, tc.msg.Type())
			require.Equal(t, types.RouterKey, tc.msg.Route())

			err := tc.msg.ValidateBasic()
			if tc.expErr == "" {
				require.NoError(t, err)
				signers := tc.msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, tc.msg.GetOrderer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestMsgLimitOrderBatch(t *testing.T) {
	orderLifespan := 20 * time.Second

	testCases := []struct {
		expErr string
		msg    *types.MsgLimitOrderBatch
	}{
		{
			"", // empty means no error expected
			types.NewMsgLimitOrderBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Orderer"))),
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 100_000_000),
				"denom2",
				parseDec("1.0"),
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
		{
			"invalid orderer address: empty address string is not allowed: invalid address",
			types.NewMsgLimitOrderBatch(
				sdk.AccAddress{},
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 100_000_000),
				"denom2",
				parseDec("1.0"),
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
		{
			"offer coin must be positive: invalid request",
			types.NewMsgLimitOrderBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Orderer"))),
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 0),
				"denom2",
				parseDec("1.0"),
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
		{
			"offer coin denom and demand coin denom must not be same: invalid request",
			types.NewMsgLimitOrderBatch(
				sdk.AccAddress(crypto.AddressHash([]byte("Orderer"))),
				1,
				types.SwapDirectionBuy,
				sdk.NewInt64Coin("denom1", 100_000_000),
				"denom1",
				parseDec("1.0"),
				sdk.NewInt(100_000_000),
				orderLifespan,
			),
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, types.TypeMsgLimitOrderBatch, tc.msg.Type())
			require.Equal(t, types.RouterKey, tc.msg.Route())

			err := tc.msg.ValidateBasic()
			if tc.expErr == "" {
				require.NoError(t, err)
				signers := tc.msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, tc.msg.GetOrderer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestMsgCancelOrder(t *testing.T) {
	testCases := []struct {
		expErr string
		msg    *types.MsgCancelOrder
	}{
		{
			"", // empty means no error expected
			types.NewMsgCancelOrder(
				sdk.AccAddress(crypto.AddressHash([]byte("Orderer"))),
				1,
				1,
			),
		},
		{
			"invalid orderer address: empty address string is not allowed: invalid address",
			types.NewMsgCancelOrder(
				sdk.AccAddress{},
				1,
				1,
			),
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, types.TypeMsgCancelOrder, tc.msg.Type())
			require.Equal(t, types.RouterKey, tc.msg.Route())

			err := tc.msg.ValidateBasic()
			if tc.expErr == "" {
				require.NoError(t, err)
				signers := tc.msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, tc.msg.GetOrderer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestMsgCancelAllOrders(t *testing.T) {
	for _, tc := range []struct {
		malleate    func(msg *types.MsgCancelAllOrders)
		expectedErr string
	}{
		{
			func(msg *types.MsgCancelAllOrders) {},
			"",
		},
		{
			func(msg *types.MsgCancelAllOrders) {
				msg.Orderer = "invalidaddr"
			},
			"invalid orderer address: decoding bech32 failed: invalid separator index -1: invalid address",
		},
		{
			func(msg *types.MsgCancelAllOrders) {
				msg.PairIds = []uint64{0}
			},
			"pair id must not be 0: invalid request",
		},
		{
			func(msg *types.MsgCancelAllOrders) {
				msg.PairIds = []uint64{1, 1}
			},
			"duplicate pair id presents in the pair id list",
		},
	} {
		t.Run("", func(t *testing.T) {
			msg := types.NewMsgCancelAllOrders(sdk.AccAddress(crypto.AddressHash([]byte("orderer"))), []uint64{1, 2, 3})
			tc.malleate(msg)
			require.Equal(t, types.TypeMsgCancelAllOrders, msg.Type())
			require.Equal(t, types.RouterKey, msg.Route())
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
				signers := msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, msg.GetOrderer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}