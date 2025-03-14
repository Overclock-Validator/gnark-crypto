// Copyright 2020 Consensys Software Inc.
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

// Code generated by consensys/gnark-crypto DO NOT EDIT

package twistededwards

import (
	"math/big"
	"sync"

	"github.com/Overclock-Validator/gnark-crypto/ecc/bls12-377/fr"
)

// CurveParams curve parameters: ax^2 + y^2 = 1 + d*x^2*y^2
type CurveParams struct {
	A, D     fr.Element
	Cofactor fr.Element
	Order    big.Int
	Base     PointAffine
}

// GetEdwardsCurve returns the twisted Edwards curve on bls12-377/Fr
func GetEdwardsCurve() CurveParams {
	initOnce.Do(initCurveParams)
	// copy to keep Order private
	var res CurveParams

	res.A.Set(&curveParams.A)
	res.D.Set(&curveParams.D)
	res.Cofactor.Set(&curveParams.Cofactor)
	res.Order.Set(&curveParams.Order)
	res.Base.Set(&curveParams.Base)

	return res
}

var (
	initOnce    sync.Once
	curveParams CurveParams
)

func initCurveParams() {
	curveParams.A.SetString("-1")
	curveParams.D.SetString("3021")
	curveParams.Cofactor.SetString("4")
	curveParams.Order.SetString("2111115437357092606062206234695386632838870926408408195193685246394721360383", 10)

	curveParams.Base.X.SetString("717051916204163000937139483451426116831771857428389560441264442629694842243")
	curveParams.Base.Y.SetString("882565546457454111605105352482086902132191855952243170543452705048019814192")
}

// mulByA multiplies fr.Element by curveParams.A
func mulByA(x *fr.Element) {
	x.Neg(x)
}
