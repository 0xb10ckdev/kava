package v0_11

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/app"
	v0_9bep3 "github.com/kava-labs/kava/x/bep3/legacy/v0_9"
	v0_9cdp "github.com/kava-labs/kava/x/cdp/legacy/v0_9"
	v0_9committee "github.com/kava-labs/kava/x/committee/legacy/v0_9"
	v0_9pricefeed "github.com/kava-labs/kava/x/pricefeed/legacy/v0_9"
)

func TestMain(m *testing.M) {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)
	app.SetBip44CoinType(config)

	os.Exit(m.Run())
}

func TestMigrateBep3(t *testing.T) {
	bz, err := ioutil.ReadFile(filepath.Join("testdata", "bep3-v09.json"))
	require.NoError(t, err)
	var oldGenState v0_9bep3.GenesisState
	cdc := app.MakeCodec()
	require.NotPanics(t, func() {
		cdc.MustUnmarshalJSON(bz, &oldGenState)
	})

	newGenState := MigrateBep3(oldGenState)
	err = newGenState.Validate()
	require.NoError(t, err)
}

func TestMigrateCdp(t *testing.T) {
	bz, err := ioutil.ReadFile(filepath.Join("testdata", "cdp-v09.json"))
	require.NoError(t, err)
	var oldGenState v0_9cdp.GenesisState
	cdc := app.MakeCodec()
	require.NotPanics(t, func() {
		cdc.MustUnmarshalJSON(bz, &oldGenState)
	})

	newGenState := MigrateCDP(oldGenState)
	err = newGenState.Validate()
	require.NoError(t, err)
}

func TestMigratePricefeed(t *testing.T) {
	bz, err := ioutil.ReadFile(filepath.Join("testdata", "pricefeed-v09.json"))
	require.NoError(t, err)
	var oldGenState v0_9pricefeed.GenesisState
	cdc := app.MakeCodec()
	require.NotPanics(t, func() {
		cdc.MustUnmarshalJSON(bz, &oldGenState)
	})
	newGenState := MigratePricefeed(oldGenState)
	err = newGenState.Validate()
	require.NoError(t, err)
}

func TestMigrateCommittee(t *testing.T) {
	bz, err := ioutil.ReadFile(filepath.Join("testdata", "committee-v09.json"))
	require.NoError(t, err)
	var oldGenState v0_9committee.GenesisState
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	v0_9committee.RegisterCodec(cdc)

	require.NotPanics(t, func() {
		cdc.MustUnmarshalJSON(bz, &oldGenState)
	})

	newGenState := MigrateCommittee(oldGenState)
	err = newGenState.Validate()
	require.NoError(t, err)
}
