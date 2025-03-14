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

package fft

import (
	"runtime"

	"github.com/Overclock-Validator/gnark-crypto/ecc/bls24-315/fr"
)

// Option defines option for altering the behavior of FFT methods.
// See the descriptions of functions returning instances of this type for
// particular options.
type Option func(*fftConfig)

type fftConfig struct {
	coset   bool
	nbTasks int
}

// OnCoset if provided, FFT(a) returns the evaluation of a on a coset.
func OnCoset() Option {
	return func(opt *fftConfig) {
		opt.coset = true
	}
}

// WithNbTasks sets the max number of task (go routine) to spawn. Must be between 1 and 512.
func WithNbTasks(nbTasks int) Option {
	if nbTasks < 1 {
		nbTasks = 1
	} else if nbTasks > 512 {
		nbTasks = 512
	}
	return func(opt *fftConfig) {
		opt.nbTasks = nbTasks
	}
}

// default options
func fftOptions(opts ...Option) fftConfig {
	// apply options
	opt := fftConfig{
		coset:   false,
		nbTasks: runtime.NumCPU(),
	}
	for _, option := range opts {
		option(&opt)
	}
	return opt
}

// DomainOption defines option for altering the definition of the FFT domain
// See the descriptions of functions returning instances of this type for
// particular options.
type DomainOption func(*domainConfig)

type domainConfig struct {
	shift          *fr.Element
	withPrecompute bool
}

// WithShift sets the FrMultiplicativeGen of the domain.
// Default is generator of the largest 2-adic subgroup.
func WithShift(shift fr.Element) DomainOption {
	return func(opt *domainConfig) {
		opt.shift = new(fr.Element).Set(&shift)
	}
}

// WithoutPrecompute disables precomputation of twiddles in the domain.
// When this option is set, FFTs will be slower, but will use less memory.
func WithoutPrecompute() DomainOption {
	return func(opt *domainConfig) {
		opt.withPrecompute = false
	}
}

// default options
func domainOptions(opts ...DomainOption) domainConfig {
	// apply options
	opt := domainConfig{
		withPrecompute: true,
	}
	for _, option := range opts {
		option(&opt)
	}
	return opt
}
