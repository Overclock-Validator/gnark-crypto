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

// Package htf provides hasher based on RFC 9380 Section 5.
//
// The [RFC 9380] defines a method for hashing bytes to elliptic curves. Section
// 5 of the RFC describes a method for uniformly hashing bytes into a field
// using a domain separation. The hashing is implemented in [fp], but this
// package provides a wrapper for the method which implements [hash.Hash] for
// using the method recursively.
//
// [RFC 9380]: https://datatracker.ietf.org/doc/html/rfc9380
package hash_to_field

import (
	_ "hash"

	_ "github.com/Overclock-Validator/gnark-crypto/ecc/bw6-633/fr"
)
