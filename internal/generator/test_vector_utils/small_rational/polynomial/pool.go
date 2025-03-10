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

package polynomial

import (
	"github.com/Overclock-Validator/gnark-crypto/internal/generator/test_vector_utils/small_rational"
)

// Do as little as possible to instantiate the interface
type Pool struct {
}

func NewPool(...int) (pool Pool) {
	return Pool{}
}

func (p *Pool) Make(n int) []small_rational.SmallRational {
	return make([]small_rational.SmallRational, n)
}

func (p *Pool) Dump(...[]small_rational.SmallRational) {
}

func (p *Pool) Clone(slice []small_rational.SmallRational) []small_rational.SmallRational {
	res := p.Make(len(slice))
	copy(res, slice)
	return res
}
