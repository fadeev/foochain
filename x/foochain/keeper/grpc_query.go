package keeper

import (
	"github.com/cosmonaut/foochain/x/foochain/types"
)

var _ types.QueryServer = Keeper{}
