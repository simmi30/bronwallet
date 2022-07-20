// Copyright (c) 2013-2015 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netparams

import "github.com/brsuite/brond/chaincfg"

// Params is used to group parameters for various networks such as the main
// network and test networks.
type Params struct {
	*chaincfg.Params
	RPCClientPort string
	RPCServerPort string
}

// MainNetParams contains parameters specific running bronwallet and
// brond on the main network (wire.MainNet).
var MainNetParams = Params{
	Params:        &chaincfg.MainNetParams,
	RPCClientPort: "8360",
	RPCServerPort: "8360",
}

// TestNet3Params contains parameters specific running bronwallet and
// brond on the test network (version 3) (wire.TestNet3).
var TestNet3Params = Params{
	Params:        &chaincfg.TestNet3Params,
	RPCClientPort: "18740",
	RPCServerPort: "18740",
}

// SimNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var SimNetParams = Params{
	Params:        &chaincfg.SimNetParams,
	RPCClientPort: "18556",
	RPCServerPort: "18554",
}
