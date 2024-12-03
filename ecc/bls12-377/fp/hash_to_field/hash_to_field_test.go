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

package hash_to_field

import (
	"testing"

	"github.com/Overclock-Validator/gnark-crypto/ecc/bls12-377/fp"
)

func TestHashInterface(t *testing.T) {
	msg := []byte("test")
	sep := []byte("separator")
	res, err := fp.Hash(msg, sep, 1)
	if err != nil {
		t.Fatal("hash to field", err)
	}

	htfFn := New(sep)
	htfFn.Write(msg)
	bts := htfFn.Sum(nil)
	var res2 fp.Element
	res2.SetBytes(bts[:fp.Bytes])
	if !res[0].Equal(&res2) {
		t.Error("not equal")
	}
}
