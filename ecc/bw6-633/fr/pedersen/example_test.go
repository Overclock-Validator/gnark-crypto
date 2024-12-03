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

package pedersen

import (
	"crypto/rand"
	"fmt"

	"github.com/Overclock-Validator/gnark-crypto/ecc"
	curve "github.com/Overclock-Validator/gnark-crypto/ecc/bw6-633"
	"github.com/Overclock-Validator/gnark-crypto/ecc/bw6-633/fr"
)

// This example demonstrates how to use the Pedersen commitment scheme
// to commit to a set of values and prove knowledge of the committed values.
//
// Does not perform any batching or multi-proof optimization.
func Example_singleProof() {
	const nbElem = 4
	// create a proving key with independent basis elements
	var buf [32]byte
	basis := make([]curve.G1Affine, nbElem)
	for i := range basis {
		_, err := rand.Read(buf[:])
		if err != nil {
			panic(err)
		}
		// we use hash-to-curve to avoid linear dependencies between basis elements
		basis[i], err = curve.HashToG1(buf[:], []byte(fmt.Sprintf("basis %d", i)))
		if err != nil {
			panic(err)
		}
	}
	// create a proving and verifying key. NB! Must be done using MPC
	pks, vk, err := Setup([][]curve.G1Affine{basis})
	if err != nil {
		panic(err)
	}
	// currently we only have a single proving key
	pk := pks[0]
	toCommit := make([]fr.Element, nbElem)
	for i := range toCommit {
		toCommit[i].SetRandom()
	}
	// commit to the values
	commitment, err := pk.Commit(toCommit)
	if err != nil {
		panic(err)
	}
	// prove knowledge of the committed values
	pok, err := pk.ProveKnowledge(toCommit)
	if err != nil {
		panic(err)
	}
	// verify the proof
	if err := vk.Verify(commitment, pok); err != nil {
		panic(err)
	}

	fmt.Println("verified")
	// output: verified
}

// This example shows how to batch the commitment and proof generation.
func ExampleBatchProve() {
	const nbPks = 3
	const nbElem = 4
	// create a proving key with independent basis elements
	var buf [32]byte
	basis := make([][]curve.G1Affine, nbPks)
	for i := range basis {
		basis[i] = make([]curve.G1Affine, nbElem)
		for j := range basis[i] {
			_, err := rand.Read(buf[:])
			if err != nil {
				panic(err)
			}
			// we use hash-to-curve to avoid linear dependencies between basis elements
			basis[i][j], err = curve.HashToG1(buf[:], []byte(fmt.Sprintf("basis %d", i)))
			if err != nil {
				panic(err)
			}
		}
	}
	// create a proving and verifying key. NB! Must be done using MPC
	pks, vk, err := Setup(basis)
	if err != nil {
		panic(err)
	}
	// generate random values to commit to
	toCommit := make([][]fr.Element, nbPks)
	for i := range toCommit {
		toCommit[i] = make([]fr.Element, nbElem)
		for j := range toCommit[i] {
			toCommit[i][j].SetRandom()
		}
	}
	// commit to the values
	commitments := make([]curve.G1Affine, nbPks)
	for i := range commitments {
		commitments[i], err = pks[i].Commit(toCommit[i])
		if err != nil {
			panic(err)
		}
	}
	// combination coefficient is randomly sampled by the verifier. NB! In non-interactive protocol use Fiat-Shamir!
	var combinationCoeff fr.Element
	combinationCoeff.SetRandom()
	proof, err := BatchProve(pks, toCommit, combinationCoeff)
	if err != nil {
		panic(err)
	}
	// fold the commitments
	foldedCommitment, err := new(curve.G1Affine).Fold(commitments, combinationCoeff, ecc.MultiExpConfig{NbTasks: 1})
	if err != nil {
		panic(err)
	}
	// verify the proof
	if err := vk.Verify(*foldedCommitment, proof); err != nil {
		panic(err)
	}
	fmt.Println("verified")

	// Output: verified
}

// This example shows how to batch verify multiple proofs using multiple
// verifying keys.
func ExampleBatchVerifyMultiVk() {
	const nbPks = 3
	const nbElem = 4
	// create a proving key with independent basis elements
	var buf [32]byte
	basis := make([][]curve.G1Affine, nbPks)
	for i := range basis {
		basis[i] = make([]curve.G1Affine, nbElem)
		for j := range basis[i] {
			_, err := rand.Read(buf[:])
			if err != nil {
				panic(err)
			}
			// we use hash-to-curve to avoid linear dependencies between basis elements
			basis[i][j], err = curve.HashToG1(buf[:], []byte(fmt.Sprintf("basis %d", i)))
			if err != nil {
				panic(err)
			}
		}
	}
	// we create independent proving keys (different sigmas) with same G2
	// g2Point does not have to be generated in a trusted manner
	_, _, _, g2Point := curve.Generators()
	pks := make([]ProvingKey, nbPks)
	vks := make([]VerifyingKey, nbPks)
	for i := range basis {
		pkss, vkss, err := Setup(basis[i:i+1], WithG2Point(g2Point))
		if err != nil {
			panic(err)
		}
		pks[i] = pkss[0]
		vks[i] = vkss
	}
	// generate random values to commit to
	toCommit := make([][]fr.Element, nbPks)
	for i := range toCommit {
		toCommit[i] = make([]fr.Element, nbElem)
		for j := range toCommit[i] {
			toCommit[i][j].SetRandom()
		}
	}
	// commit to the values
	commitments := make([]curve.G1Affine, nbPks)
	for i := range commitments {
		var err error
		commitments[i], err = pks[i].Commit(toCommit[i])
		if err != nil {
			panic(err)
		}
	}
	// prove the commitments
	proofs := make([]curve.G1Affine, nbPks)
	for i := range proofs {
		var err error
		proofs[i], err = pks[i].ProveKnowledge(toCommit[i])
		if err != nil {
			panic(err)
		}
	}
	// combination coefficient is randomly sampled by the verifier. NB! In non-interactive protocol use Fiat-Shamir!
	var combinationCoeff fr.Element
	combinationCoeff.SetRandom()
	// batch verify the proofs
	if err := BatchVerifyMultiVk(vks, commitments, proofs, combinationCoeff); err != nil {
		panic(err)
	}

	// alternatively, we can also provide the folded proof
	foldedProof, err := new(curve.G1Affine).Fold(proofs, combinationCoeff, ecc.MultiExpConfig{NbTasks: 1})
	if err != nil {
		panic(err)
	}
	if err := BatchVerifyMultiVk(vks, commitments, []curve.G1Affine{*foldedProof}, combinationCoeff); err != nil {
		panic(err)
	}

	fmt.Println("verified")
	// Output: verified
}
