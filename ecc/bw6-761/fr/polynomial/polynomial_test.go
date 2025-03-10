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
	"github.com/Overclock-Validator/gnark-crypto/ecc/bw6-761/fr"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPolynomialEval(t *testing.T) {

	// build polynomial
	f := make(Polynomial, 20)
	for i := 0; i < 20; i++ {
		f[i].SetOne()
	}

	// random value
	var point fr.Element
	point.SetRandom()

	// compute manually f(val)
	var expectedEval, one, den fr.Element
	var expo big.Int
	one.SetOne()
	expo.SetUint64(20)
	expectedEval.Exp(point, &expo).
		Sub(&expectedEval, &one)
	den.Sub(&point, &one)
	expectedEval.Div(&expectedEval, &den)

	// compute purported evaluation
	purportedEval := f.Eval(&point)

	// check
	if !purportedEval.Equal(&expectedEval) {
		t.Fatal("polynomial evaluation failed")
	}
}

func TestPolynomialAddConstantInPlace(t *testing.T) {

	// build polynomial
	f := make(Polynomial, 20)
	for i := 0; i < 20; i++ {
		f[i].SetOne()
	}

	// constant to add
	var c fr.Element
	c.SetRandom()

	// add constant
	f.AddConstantInPlace(&c)

	// check
	var expectedCoeffs, one fr.Element
	one.SetOne()
	expectedCoeffs.Add(&one, &c)
	for i := 0; i < 20; i++ {
		if !f[i].Equal(&expectedCoeffs) {
			t.Fatal("AddConstantInPlace failed")
		}
	}
}

func TestPolynomialSubConstantInPlace(t *testing.T) {

	// build polynomial
	f := make(Polynomial, 20)
	for i := 0; i < 20; i++ {
		f[i].SetOne()
	}

	// constant to sub
	var c fr.Element
	c.SetRandom()

	// sub constant
	f.SubConstantInPlace(&c)

	// check
	var expectedCoeffs, one fr.Element
	one.SetOne()
	expectedCoeffs.Sub(&one, &c)
	for i := 0; i < 20; i++ {
		if !f[i].Equal(&expectedCoeffs) {
			t.Fatal("SubConstantInPlace failed")
		}
	}
}

func TestPolynomialScaleInPlace(t *testing.T) {

	// build polynomial
	f := make(Polynomial, 20)
	for i := 0; i < 20; i++ {
		f[i].SetOne()
	}

	// constant to scale by
	var c fr.Element
	c.SetRandom()

	// scale by constant
	f.ScaleInPlace(&c)

	// check
	for i := 0; i < 20; i++ {
		if !f[i].Equal(&c) {
			t.Fatal("ScaleInPlace failed")
		}
	}

}

func TestPolynomialAdd(t *testing.T) {

	// build unbalanced polynomials
	f1 := make(Polynomial, 20)
	f1Backup := make(Polynomial, 20)
	for i := 0; i < 20; i++ {
		f1[i].SetOne()
		f1Backup[i].SetOne()
	}
	f2 := make(Polynomial, 10)
	f2Backup := make(Polynomial, 10)
	for i := 0; i < 10; i++ {
		f2[i].SetOne()
		f2Backup[i].SetOne()
	}

	// expected result
	var one, two fr.Element
	one.SetOne()
	two.Double(&one)
	expectedSum := make(Polynomial, 20)
	for i := 0; i < 10; i++ {
		expectedSum[i].Set(&two)
	}
	for i := 10; i < 20; i++ {
		expectedSum[i].Set(&one)
	}

	// caller is empty
	var g Polynomial
	g.Add(f1, f2)
	if !g.Equal(expectedSum) {
		t.Fatal("add polynomials fails")
	}
	if !f1.Equal(f1Backup) {
		t.Fatal("side effect, f1 should not have been modified")
	}
	if !f2.Equal(f2Backup) {
		t.Fatal("side effect, f2 should not have been modified")
	}

	// all operands are distinct
	_f1 := f1.Clone()
	_f1.Add(f1, f2)
	if !_f1.Equal(expectedSum) {
		t.Fatal("add polynomials fails")
	}
	if !f1.Equal(f1Backup) {
		t.Fatal("side effect, f1 should not have been modified")
	}
	if !f2.Equal(f2Backup) {
		t.Fatal("side effect, f2 should not have been modified")
	}

	// first operand = caller
	_f1 = f1.Clone()
	_f2 := f2.Clone()
	_f1.Add(_f1, _f2)
	if !_f1.Equal(expectedSum) {
		t.Fatal("add polynomials fails")
	}
	if !_f2.Equal(f2Backup) {
		t.Fatal("side effect, _f2 should not have been modified")
	}

	// second operand = caller
	_f1 = f1.Clone()
	_f2 = f2.Clone()
	_f1.Add(_f2, _f1)
	if !_f1.Equal(expectedSum) {
		t.Fatal("add polynomials fails")
	}
	if !_f2.Equal(f2Backup) {
		t.Fatal("side effect, _f2 should not have been modified")
	}
}

func TestPolynomialText(t *testing.T) {
	var one, negTwo fr.Element
	one.SetOne()
	negTwo.SetInt64(-2)

	p := Polynomial{one, negTwo, one}

	assert.Equal(t, "X² - 2X + 1", p.Text(10))
}
