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

package sis

import (
	"github.com/Overclock-Validator/gnark-crypto/ecc/bn254/fr"
	"math/big"
)

// FFT64 is generated by gnark-crypto and contains the unrolled code for FFT (DIF) on 64 elements
// equivalent code: r.Domain.FFT(k, fft.DIF, fft.OnCoset(), fft.WithNbTasks(1))
// twiddlesCoset must be pre-computed from twiddles and coset table, see PrecomputeTwiddlesCoset
func FFT64(a []fr.Element, twiddlesCoset []fr.Element) {

	a[32].Mul(&a[32], &twiddlesCoset[0])
	a[33].Mul(&a[33], &twiddlesCoset[0])
	a[34].Mul(&a[34], &twiddlesCoset[0])
	a[35].Mul(&a[35], &twiddlesCoset[0])
	a[36].Mul(&a[36], &twiddlesCoset[0])
	a[37].Mul(&a[37], &twiddlesCoset[0])
	a[38].Mul(&a[38], &twiddlesCoset[0])
	a[39].Mul(&a[39], &twiddlesCoset[0])
	a[40].Mul(&a[40], &twiddlesCoset[0])
	a[41].Mul(&a[41], &twiddlesCoset[0])
	a[42].Mul(&a[42], &twiddlesCoset[0])
	a[43].Mul(&a[43], &twiddlesCoset[0])
	a[44].Mul(&a[44], &twiddlesCoset[0])
	a[45].Mul(&a[45], &twiddlesCoset[0])
	a[46].Mul(&a[46], &twiddlesCoset[0])
	a[47].Mul(&a[47], &twiddlesCoset[0])
	a[48].Mul(&a[48], &twiddlesCoset[0])
	a[49].Mul(&a[49], &twiddlesCoset[0])
	a[50].Mul(&a[50], &twiddlesCoset[0])
	a[51].Mul(&a[51], &twiddlesCoset[0])
	a[52].Mul(&a[52], &twiddlesCoset[0])
	a[53].Mul(&a[53], &twiddlesCoset[0])
	a[54].Mul(&a[54], &twiddlesCoset[0])
	a[55].Mul(&a[55], &twiddlesCoset[0])
	a[56].Mul(&a[56], &twiddlesCoset[0])
	a[57].Mul(&a[57], &twiddlesCoset[0])
	a[58].Mul(&a[58], &twiddlesCoset[0])
	a[59].Mul(&a[59], &twiddlesCoset[0])
	a[60].Mul(&a[60], &twiddlesCoset[0])
	a[61].Mul(&a[61], &twiddlesCoset[0])
	a[62].Mul(&a[62], &twiddlesCoset[0])
	a[63].Mul(&a[63], &twiddlesCoset[0])
	fr.Butterfly(&a[0], &a[32])
	fr.Butterfly(&a[1], &a[33])
	fr.Butterfly(&a[2], &a[34])
	fr.Butterfly(&a[3], &a[35])
	fr.Butterfly(&a[4], &a[36])
	fr.Butterfly(&a[5], &a[37])
	fr.Butterfly(&a[6], &a[38])
	fr.Butterfly(&a[7], &a[39])
	fr.Butterfly(&a[8], &a[40])
	fr.Butterfly(&a[9], &a[41])
	fr.Butterfly(&a[10], &a[42])
	fr.Butterfly(&a[11], &a[43])
	fr.Butterfly(&a[12], &a[44])
	fr.Butterfly(&a[13], &a[45])
	fr.Butterfly(&a[14], &a[46])
	fr.Butterfly(&a[15], &a[47])
	fr.Butterfly(&a[16], &a[48])
	fr.Butterfly(&a[17], &a[49])
	fr.Butterfly(&a[18], &a[50])
	fr.Butterfly(&a[19], &a[51])
	fr.Butterfly(&a[20], &a[52])
	fr.Butterfly(&a[21], &a[53])
	fr.Butterfly(&a[22], &a[54])
	fr.Butterfly(&a[23], &a[55])
	fr.Butterfly(&a[24], &a[56])
	fr.Butterfly(&a[25], &a[57])
	fr.Butterfly(&a[26], &a[58])
	fr.Butterfly(&a[27], &a[59])
	fr.Butterfly(&a[28], &a[60])
	fr.Butterfly(&a[29], &a[61])
	fr.Butterfly(&a[30], &a[62])
	fr.Butterfly(&a[31], &a[63])
	a[16].Mul(&a[16], &twiddlesCoset[1])
	a[17].Mul(&a[17], &twiddlesCoset[1])
	a[18].Mul(&a[18], &twiddlesCoset[1])
	a[19].Mul(&a[19], &twiddlesCoset[1])
	a[20].Mul(&a[20], &twiddlesCoset[1])
	a[21].Mul(&a[21], &twiddlesCoset[1])
	a[22].Mul(&a[22], &twiddlesCoset[1])
	a[23].Mul(&a[23], &twiddlesCoset[1])
	a[24].Mul(&a[24], &twiddlesCoset[1])
	a[25].Mul(&a[25], &twiddlesCoset[1])
	a[26].Mul(&a[26], &twiddlesCoset[1])
	a[27].Mul(&a[27], &twiddlesCoset[1])
	a[28].Mul(&a[28], &twiddlesCoset[1])
	a[29].Mul(&a[29], &twiddlesCoset[1])
	a[30].Mul(&a[30], &twiddlesCoset[1])
	a[31].Mul(&a[31], &twiddlesCoset[1])
	a[48].Mul(&a[48], &twiddlesCoset[2])
	a[49].Mul(&a[49], &twiddlesCoset[2])
	a[50].Mul(&a[50], &twiddlesCoset[2])
	a[51].Mul(&a[51], &twiddlesCoset[2])
	a[52].Mul(&a[52], &twiddlesCoset[2])
	a[53].Mul(&a[53], &twiddlesCoset[2])
	a[54].Mul(&a[54], &twiddlesCoset[2])
	a[55].Mul(&a[55], &twiddlesCoset[2])
	a[56].Mul(&a[56], &twiddlesCoset[2])
	a[57].Mul(&a[57], &twiddlesCoset[2])
	a[58].Mul(&a[58], &twiddlesCoset[2])
	a[59].Mul(&a[59], &twiddlesCoset[2])
	a[60].Mul(&a[60], &twiddlesCoset[2])
	a[61].Mul(&a[61], &twiddlesCoset[2])
	a[62].Mul(&a[62], &twiddlesCoset[2])
	a[63].Mul(&a[63], &twiddlesCoset[2])
	fr.Butterfly(&a[0], &a[16])
	fr.Butterfly(&a[1], &a[17])
	fr.Butterfly(&a[2], &a[18])
	fr.Butterfly(&a[3], &a[19])
	fr.Butterfly(&a[4], &a[20])
	fr.Butterfly(&a[5], &a[21])
	fr.Butterfly(&a[6], &a[22])
	fr.Butterfly(&a[7], &a[23])
	fr.Butterfly(&a[8], &a[24])
	fr.Butterfly(&a[9], &a[25])
	fr.Butterfly(&a[10], &a[26])
	fr.Butterfly(&a[11], &a[27])
	fr.Butterfly(&a[12], &a[28])
	fr.Butterfly(&a[13], &a[29])
	fr.Butterfly(&a[14], &a[30])
	fr.Butterfly(&a[15], &a[31])
	fr.Butterfly(&a[32], &a[48])
	fr.Butterfly(&a[33], &a[49])
	fr.Butterfly(&a[34], &a[50])
	fr.Butterfly(&a[35], &a[51])
	fr.Butterfly(&a[36], &a[52])
	fr.Butterfly(&a[37], &a[53])
	fr.Butterfly(&a[38], &a[54])
	fr.Butterfly(&a[39], &a[55])
	fr.Butterfly(&a[40], &a[56])
	fr.Butterfly(&a[41], &a[57])
	fr.Butterfly(&a[42], &a[58])
	fr.Butterfly(&a[43], &a[59])
	fr.Butterfly(&a[44], &a[60])
	fr.Butterfly(&a[45], &a[61])
	fr.Butterfly(&a[46], &a[62])
	fr.Butterfly(&a[47], &a[63])
	a[8].Mul(&a[8], &twiddlesCoset[3])
	a[9].Mul(&a[9], &twiddlesCoset[3])
	a[10].Mul(&a[10], &twiddlesCoset[3])
	a[11].Mul(&a[11], &twiddlesCoset[3])
	a[12].Mul(&a[12], &twiddlesCoset[3])
	a[13].Mul(&a[13], &twiddlesCoset[3])
	a[14].Mul(&a[14], &twiddlesCoset[3])
	a[15].Mul(&a[15], &twiddlesCoset[3])
	a[24].Mul(&a[24], &twiddlesCoset[4])
	a[25].Mul(&a[25], &twiddlesCoset[4])
	a[26].Mul(&a[26], &twiddlesCoset[4])
	a[27].Mul(&a[27], &twiddlesCoset[4])
	a[28].Mul(&a[28], &twiddlesCoset[4])
	a[29].Mul(&a[29], &twiddlesCoset[4])
	a[30].Mul(&a[30], &twiddlesCoset[4])
	a[31].Mul(&a[31], &twiddlesCoset[4])
	a[40].Mul(&a[40], &twiddlesCoset[5])
	a[41].Mul(&a[41], &twiddlesCoset[5])
	a[42].Mul(&a[42], &twiddlesCoset[5])
	a[43].Mul(&a[43], &twiddlesCoset[5])
	a[44].Mul(&a[44], &twiddlesCoset[5])
	a[45].Mul(&a[45], &twiddlesCoset[5])
	a[46].Mul(&a[46], &twiddlesCoset[5])
	a[47].Mul(&a[47], &twiddlesCoset[5])
	a[56].Mul(&a[56], &twiddlesCoset[6])
	a[57].Mul(&a[57], &twiddlesCoset[6])
	a[58].Mul(&a[58], &twiddlesCoset[6])
	a[59].Mul(&a[59], &twiddlesCoset[6])
	a[60].Mul(&a[60], &twiddlesCoset[6])
	a[61].Mul(&a[61], &twiddlesCoset[6])
	a[62].Mul(&a[62], &twiddlesCoset[6])
	a[63].Mul(&a[63], &twiddlesCoset[6])
	fr.Butterfly(&a[0], &a[8])
	fr.Butterfly(&a[1], &a[9])
	fr.Butterfly(&a[2], &a[10])
	fr.Butterfly(&a[3], &a[11])
	fr.Butterfly(&a[4], &a[12])
	fr.Butterfly(&a[5], &a[13])
	fr.Butterfly(&a[6], &a[14])
	fr.Butterfly(&a[7], &a[15])
	fr.Butterfly(&a[16], &a[24])
	fr.Butterfly(&a[17], &a[25])
	fr.Butterfly(&a[18], &a[26])
	fr.Butterfly(&a[19], &a[27])
	fr.Butterfly(&a[20], &a[28])
	fr.Butterfly(&a[21], &a[29])
	fr.Butterfly(&a[22], &a[30])
	fr.Butterfly(&a[23], &a[31])
	fr.Butterfly(&a[32], &a[40])
	fr.Butterfly(&a[33], &a[41])
	fr.Butterfly(&a[34], &a[42])
	fr.Butterfly(&a[35], &a[43])
	fr.Butterfly(&a[36], &a[44])
	fr.Butterfly(&a[37], &a[45])
	fr.Butterfly(&a[38], &a[46])
	fr.Butterfly(&a[39], &a[47])
	fr.Butterfly(&a[48], &a[56])
	fr.Butterfly(&a[49], &a[57])
	fr.Butterfly(&a[50], &a[58])
	fr.Butterfly(&a[51], &a[59])
	fr.Butterfly(&a[52], &a[60])
	fr.Butterfly(&a[53], &a[61])
	fr.Butterfly(&a[54], &a[62])
	fr.Butterfly(&a[55], &a[63])
	a[4].Mul(&a[4], &twiddlesCoset[7])
	a[5].Mul(&a[5], &twiddlesCoset[7])
	a[6].Mul(&a[6], &twiddlesCoset[7])
	a[7].Mul(&a[7], &twiddlesCoset[7])
	a[12].Mul(&a[12], &twiddlesCoset[8])
	a[13].Mul(&a[13], &twiddlesCoset[8])
	a[14].Mul(&a[14], &twiddlesCoset[8])
	a[15].Mul(&a[15], &twiddlesCoset[8])
	a[20].Mul(&a[20], &twiddlesCoset[9])
	a[21].Mul(&a[21], &twiddlesCoset[9])
	a[22].Mul(&a[22], &twiddlesCoset[9])
	a[23].Mul(&a[23], &twiddlesCoset[9])
	a[28].Mul(&a[28], &twiddlesCoset[10])
	a[29].Mul(&a[29], &twiddlesCoset[10])
	a[30].Mul(&a[30], &twiddlesCoset[10])
	a[31].Mul(&a[31], &twiddlesCoset[10])
	a[36].Mul(&a[36], &twiddlesCoset[11])
	a[37].Mul(&a[37], &twiddlesCoset[11])
	a[38].Mul(&a[38], &twiddlesCoset[11])
	a[39].Mul(&a[39], &twiddlesCoset[11])
	a[44].Mul(&a[44], &twiddlesCoset[12])
	a[45].Mul(&a[45], &twiddlesCoset[12])
	a[46].Mul(&a[46], &twiddlesCoset[12])
	a[47].Mul(&a[47], &twiddlesCoset[12])
	a[52].Mul(&a[52], &twiddlesCoset[13])
	a[53].Mul(&a[53], &twiddlesCoset[13])
	a[54].Mul(&a[54], &twiddlesCoset[13])
	a[55].Mul(&a[55], &twiddlesCoset[13])
	a[60].Mul(&a[60], &twiddlesCoset[14])
	a[61].Mul(&a[61], &twiddlesCoset[14])
	a[62].Mul(&a[62], &twiddlesCoset[14])
	a[63].Mul(&a[63], &twiddlesCoset[14])
	fr.Butterfly(&a[0], &a[4])
	fr.Butterfly(&a[1], &a[5])
	fr.Butterfly(&a[2], &a[6])
	fr.Butterfly(&a[3], &a[7])
	fr.Butterfly(&a[8], &a[12])
	fr.Butterfly(&a[9], &a[13])
	fr.Butterfly(&a[10], &a[14])
	fr.Butterfly(&a[11], &a[15])
	fr.Butterfly(&a[16], &a[20])
	fr.Butterfly(&a[17], &a[21])
	fr.Butterfly(&a[18], &a[22])
	fr.Butterfly(&a[19], &a[23])
	fr.Butterfly(&a[24], &a[28])
	fr.Butterfly(&a[25], &a[29])
	fr.Butterfly(&a[26], &a[30])
	fr.Butterfly(&a[27], &a[31])
	fr.Butterfly(&a[32], &a[36])
	fr.Butterfly(&a[33], &a[37])
	fr.Butterfly(&a[34], &a[38])
	fr.Butterfly(&a[35], &a[39])
	fr.Butterfly(&a[40], &a[44])
	fr.Butterfly(&a[41], &a[45])
	fr.Butterfly(&a[42], &a[46])
	fr.Butterfly(&a[43], &a[47])
	fr.Butterfly(&a[48], &a[52])
	fr.Butterfly(&a[49], &a[53])
	fr.Butterfly(&a[50], &a[54])
	fr.Butterfly(&a[51], &a[55])
	fr.Butterfly(&a[56], &a[60])
	fr.Butterfly(&a[57], &a[61])
	fr.Butterfly(&a[58], &a[62])
	fr.Butterfly(&a[59], &a[63])
	a[2].Mul(&a[2], &twiddlesCoset[15])
	a[3].Mul(&a[3], &twiddlesCoset[15])
	a[6].Mul(&a[6], &twiddlesCoset[16])
	a[7].Mul(&a[7], &twiddlesCoset[16])
	a[10].Mul(&a[10], &twiddlesCoset[17])
	a[11].Mul(&a[11], &twiddlesCoset[17])
	a[14].Mul(&a[14], &twiddlesCoset[18])
	a[15].Mul(&a[15], &twiddlesCoset[18])
	a[18].Mul(&a[18], &twiddlesCoset[19])
	a[19].Mul(&a[19], &twiddlesCoset[19])
	a[22].Mul(&a[22], &twiddlesCoset[20])
	a[23].Mul(&a[23], &twiddlesCoset[20])
	a[26].Mul(&a[26], &twiddlesCoset[21])
	a[27].Mul(&a[27], &twiddlesCoset[21])
	a[30].Mul(&a[30], &twiddlesCoset[22])
	a[31].Mul(&a[31], &twiddlesCoset[22])
	a[34].Mul(&a[34], &twiddlesCoset[23])
	a[35].Mul(&a[35], &twiddlesCoset[23])
	a[38].Mul(&a[38], &twiddlesCoset[24])
	a[39].Mul(&a[39], &twiddlesCoset[24])
	a[42].Mul(&a[42], &twiddlesCoset[25])
	a[43].Mul(&a[43], &twiddlesCoset[25])
	a[46].Mul(&a[46], &twiddlesCoset[26])
	a[47].Mul(&a[47], &twiddlesCoset[26])
	a[50].Mul(&a[50], &twiddlesCoset[27])
	a[51].Mul(&a[51], &twiddlesCoset[27])
	a[54].Mul(&a[54], &twiddlesCoset[28])
	a[55].Mul(&a[55], &twiddlesCoset[28])
	a[58].Mul(&a[58], &twiddlesCoset[29])
	a[59].Mul(&a[59], &twiddlesCoset[29])
	a[62].Mul(&a[62], &twiddlesCoset[30])
	a[63].Mul(&a[63], &twiddlesCoset[30])
	fr.Butterfly(&a[0], &a[2])
	fr.Butterfly(&a[1], &a[3])
	fr.Butterfly(&a[4], &a[6])
	fr.Butterfly(&a[5], &a[7])
	fr.Butterfly(&a[8], &a[10])
	fr.Butterfly(&a[9], &a[11])
	fr.Butterfly(&a[12], &a[14])
	fr.Butterfly(&a[13], &a[15])
	fr.Butterfly(&a[16], &a[18])
	fr.Butterfly(&a[17], &a[19])
	fr.Butterfly(&a[20], &a[22])
	fr.Butterfly(&a[21], &a[23])
	fr.Butterfly(&a[24], &a[26])
	fr.Butterfly(&a[25], &a[27])
	fr.Butterfly(&a[28], &a[30])
	fr.Butterfly(&a[29], &a[31])
	fr.Butterfly(&a[32], &a[34])
	fr.Butterfly(&a[33], &a[35])
	fr.Butterfly(&a[36], &a[38])
	fr.Butterfly(&a[37], &a[39])
	fr.Butterfly(&a[40], &a[42])
	fr.Butterfly(&a[41], &a[43])
	fr.Butterfly(&a[44], &a[46])
	fr.Butterfly(&a[45], &a[47])
	fr.Butterfly(&a[48], &a[50])
	fr.Butterfly(&a[49], &a[51])
	fr.Butterfly(&a[52], &a[54])
	fr.Butterfly(&a[53], &a[55])
	fr.Butterfly(&a[56], &a[58])
	fr.Butterfly(&a[57], &a[59])
	fr.Butterfly(&a[60], &a[62])
	fr.Butterfly(&a[61], &a[63])
	a[1].Mul(&a[1], &twiddlesCoset[31])
	a[3].Mul(&a[3], &twiddlesCoset[32])
	a[5].Mul(&a[5], &twiddlesCoset[33])
	a[7].Mul(&a[7], &twiddlesCoset[34])
	a[9].Mul(&a[9], &twiddlesCoset[35])
	a[11].Mul(&a[11], &twiddlesCoset[36])
	a[13].Mul(&a[13], &twiddlesCoset[37])
	a[15].Mul(&a[15], &twiddlesCoset[38])
	a[17].Mul(&a[17], &twiddlesCoset[39])
	a[19].Mul(&a[19], &twiddlesCoset[40])
	a[21].Mul(&a[21], &twiddlesCoset[41])
	a[23].Mul(&a[23], &twiddlesCoset[42])
	a[25].Mul(&a[25], &twiddlesCoset[43])
	a[27].Mul(&a[27], &twiddlesCoset[44])
	a[29].Mul(&a[29], &twiddlesCoset[45])
	a[31].Mul(&a[31], &twiddlesCoset[46])
	a[33].Mul(&a[33], &twiddlesCoset[47])
	a[35].Mul(&a[35], &twiddlesCoset[48])
	a[37].Mul(&a[37], &twiddlesCoset[49])
	a[39].Mul(&a[39], &twiddlesCoset[50])
	a[41].Mul(&a[41], &twiddlesCoset[51])
	a[43].Mul(&a[43], &twiddlesCoset[52])
	a[45].Mul(&a[45], &twiddlesCoset[53])
	a[47].Mul(&a[47], &twiddlesCoset[54])
	a[49].Mul(&a[49], &twiddlesCoset[55])
	a[51].Mul(&a[51], &twiddlesCoset[56])
	a[53].Mul(&a[53], &twiddlesCoset[57])
	a[55].Mul(&a[55], &twiddlesCoset[58])
	a[57].Mul(&a[57], &twiddlesCoset[59])
	a[59].Mul(&a[59], &twiddlesCoset[60])
	a[61].Mul(&a[61], &twiddlesCoset[61])
	a[63].Mul(&a[63], &twiddlesCoset[62])
	fr.Butterfly(&a[0], &a[1])
	fr.Butterfly(&a[2], &a[3])
	fr.Butterfly(&a[4], &a[5])
	fr.Butterfly(&a[6], &a[7])
	fr.Butterfly(&a[8], &a[9])
	fr.Butterfly(&a[10], &a[11])
	fr.Butterfly(&a[12], &a[13])
	fr.Butterfly(&a[14], &a[15])
	fr.Butterfly(&a[16], &a[17])
	fr.Butterfly(&a[18], &a[19])
	fr.Butterfly(&a[20], &a[21])
	fr.Butterfly(&a[22], &a[23])
	fr.Butterfly(&a[24], &a[25])
	fr.Butterfly(&a[26], &a[27])
	fr.Butterfly(&a[28], &a[29])
	fr.Butterfly(&a[30], &a[31])
	fr.Butterfly(&a[32], &a[33])
	fr.Butterfly(&a[34], &a[35])
	fr.Butterfly(&a[36], &a[37])
	fr.Butterfly(&a[38], &a[39])
	fr.Butterfly(&a[40], &a[41])
	fr.Butterfly(&a[42], &a[43])
	fr.Butterfly(&a[44], &a[45])
	fr.Butterfly(&a[46], &a[47])
	fr.Butterfly(&a[48], &a[49])
	fr.Butterfly(&a[50], &a[51])
	fr.Butterfly(&a[52], &a[53])
	fr.Butterfly(&a[54], &a[55])
	fr.Butterfly(&a[56], &a[57])
	fr.Butterfly(&a[58], &a[59])
	fr.Butterfly(&a[60], &a[61])
	fr.Butterfly(&a[62], &a[63])
}

// PrecomputeTwiddlesCoset precomputes twiddlesCoset from twiddles and coset table
// it then return all elements in the correct order for the unrolled FFT.
func PrecomputeTwiddlesCoset(generator, shifter fr.Element) []fr.Element {
	toReturn := make([]fr.Element, 63)
	var r, s fr.Element
	e := new(big.Int)

	s = shifter
	for k := 0; k < 5; k++ {
		s.Square(&s)
	}
	toReturn[0] = s
	s = shifter
	for k := 0; k < 4; k++ {
		s.Square(&s)
	}
	toReturn[1] = s
	r.Exp(generator, e.SetUint64(uint64(1<<4*1)))
	toReturn[2].Mul(&r, &s)
	s = shifter
	for k := 0; k < 3; k++ {
		s.Square(&s)
	}
	toReturn[3] = s
	r.Exp(generator, e.SetUint64(uint64(1<<3*2)))
	toReturn[4].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<3*1)))
	toReturn[5].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<3*3)))
	toReturn[6].Mul(&r, &s)
	s = shifter
	for k := 0; k < 2; k++ {
		s.Square(&s)
	}
	toReturn[7] = s
	r.Exp(generator, e.SetUint64(uint64(1<<2*4)))
	toReturn[8].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<2*2)))
	toReturn[9].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<2*6)))
	toReturn[10].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<2*1)))
	toReturn[11].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<2*5)))
	toReturn[12].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<2*3)))
	toReturn[13].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<2*7)))
	toReturn[14].Mul(&r, &s)
	s = shifter
	for k := 0; k < 1; k++ {
		s.Square(&s)
	}
	toReturn[15] = s
	r.Exp(generator, e.SetUint64(uint64(1<<1*8)))
	toReturn[16].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*4)))
	toReturn[17].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*12)))
	toReturn[18].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*2)))
	toReturn[19].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*10)))
	toReturn[20].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*6)))
	toReturn[21].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*14)))
	toReturn[22].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*1)))
	toReturn[23].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*9)))
	toReturn[24].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*5)))
	toReturn[25].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*13)))
	toReturn[26].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*3)))
	toReturn[27].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*11)))
	toReturn[28].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*7)))
	toReturn[29].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<1*15)))
	toReturn[30].Mul(&r, &s)
	s = shifter
	for k := 0; k < 0; k++ {
		s.Square(&s)
	}
	toReturn[31] = s
	r.Exp(generator, e.SetUint64(uint64(1<<0*16)))
	toReturn[32].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*8)))
	toReturn[33].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*24)))
	toReturn[34].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*4)))
	toReturn[35].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*20)))
	toReturn[36].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*12)))
	toReturn[37].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*28)))
	toReturn[38].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*2)))
	toReturn[39].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*18)))
	toReturn[40].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*10)))
	toReturn[41].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*26)))
	toReturn[42].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*6)))
	toReturn[43].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*22)))
	toReturn[44].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*14)))
	toReturn[45].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*30)))
	toReturn[46].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*1)))
	toReturn[47].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*17)))
	toReturn[48].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*9)))
	toReturn[49].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*25)))
	toReturn[50].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*5)))
	toReturn[51].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*21)))
	toReturn[52].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*13)))
	toReturn[53].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*29)))
	toReturn[54].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*3)))
	toReturn[55].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*19)))
	toReturn[56].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*11)))
	toReturn[57].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*27)))
	toReturn[58].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*7)))
	toReturn[59].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*23)))
	toReturn[60].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*15)))
	toReturn[61].Mul(&r, &s)
	r.Exp(generator, e.SetUint64(uint64(1<<0*31)))
	toReturn[62].Mul(&r, &s)
	return toReturn
}
