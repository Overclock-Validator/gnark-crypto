// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fptower

import (
	"math/big"

	"github.com/Overclock-Validator/gnark-crypto/ecc"
	"github.com/Overclock-Validator/gnark-crypto/ecc/bls24-315/fr"
)

// generator of the curve
var xGen big.Int

var glvBasis ecc.Lattice

func init() {
	xGen.SetString("-3218079743", 10)
	_r := fr.Modulus()
	ecc.PrecomputeLattice(_r, &xGen, &glvBasis)
}
