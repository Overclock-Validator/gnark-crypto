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

package kzg

import (
	"github.com/consensys/gnark-crypto/ecc/bls12-378"
	"io"
)

// WriteTo writes binary encoding of the ProvingKey
func (pk *ProvingKey) WriteTo(w io.Writer) (int64, error) {
	return pk.writeTo(w)
}

// WriteRawTo writes binary encoding of ProvingKey to w without point compression
func (pk *ProvingKey) WriteRawTo(w io.Writer) (int64, error) {
	return pk.writeTo(w, bls12378.RawEncoding())
}

func (pk *ProvingKey) writeTo(w io.Writer, options ...func(*bls12378.Encoder)) (int64, error) {
	// encode the ProvingKey
	enc := bls12378.NewEncoder(w, options...)
	if err := enc.Encode(pk.G1); err != nil {
		return enc.BytesWritten(), err
	}
	return enc.BytesWritten(), nil
}

// WriteRawTo writes binary encoding of VerifyingKey to w without point compression
func (vk *VerifyingKey) WriteRawTo(w io.Writer) (int64, error) {
	return vk.writeTo(w, bls12378.RawEncoding())
}

// WriteTo writes binary encoding of the VerifyingKey
func (vk *VerifyingKey) WriteTo(w io.Writer) (int64, error) {
	return vk.writeTo(w)
}

func (vk *VerifyingKey) writeTo(w io.Writer, options ...func(*bls12378.Encoder)) (int64, error) {
	// encode the VerifyingKey
	enc := bls12378.NewEncoder(w, options...)

	toEncode := []interface{}{
		&vk.G2[0],
		&vk.G2[1],
		&vk.G1,
		&vk.Lines[0][0][62].R0,
		&vk.Lines[0][0][62].R1,
		&vk.Lines[0][0][61].R0,
		&vk.Lines[0][0][61].R1,
		&vk.Lines[0][0][60].R0,
		&vk.Lines[0][0][60].R1,
		&vk.Lines[0][0][59].R0,
		&vk.Lines[0][0][59].R1,
		&vk.Lines[0][0][58].R0,
		&vk.Lines[0][0][58].R1,
		&vk.Lines[0][0][57].R0,
		&vk.Lines[0][0][57].R1,
		&vk.Lines[0][0][56].R0,
		&vk.Lines[0][0][56].R1,
		&vk.Lines[0][0][55].R0,
		&vk.Lines[0][0][55].R1,
		&vk.Lines[0][0][54].R0,
		&vk.Lines[0][0][54].R1,
		&vk.Lines[0][0][53].R0,
		&vk.Lines[0][0][53].R1,
		&vk.Lines[0][0][52].R0,
		&vk.Lines[0][0][52].R1,
		&vk.Lines[0][0][51].R0,
		&vk.Lines[0][0][51].R1,
		&vk.Lines[0][0][50].R0,
		&vk.Lines[0][0][50].R1,
		&vk.Lines[0][0][49].R0,
		&vk.Lines[0][0][49].R1,
		&vk.Lines[0][0][48].R0,
		&vk.Lines[0][0][48].R1,
		&vk.Lines[0][0][47].R0,
		&vk.Lines[0][0][47].R1,
		&vk.Lines[0][0][46].R0,
		&vk.Lines[0][0][46].R1,
		&vk.Lines[0][0][45].R0,
		&vk.Lines[0][0][45].R1,
		&vk.Lines[0][0][44].R0,
		&vk.Lines[0][0][44].R1,
		&vk.Lines[0][0][43].R0,
		&vk.Lines[0][0][43].R1,
		&vk.Lines[0][0][42].R0,
		&vk.Lines[0][0][42].R1,
		&vk.Lines[0][0][41].R0,
		&vk.Lines[0][0][41].R1,
		&vk.Lines[0][0][40].R0,
		&vk.Lines[0][0][40].R1,
		&vk.Lines[0][0][39].R0,
		&vk.Lines[0][0][39].R1,
		&vk.Lines[0][0][38].R0,
		&vk.Lines[0][0][38].R1,
		&vk.Lines[0][0][37].R0,
		&vk.Lines[0][0][37].R1,
		&vk.Lines[0][0][36].R0,
		&vk.Lines[0][0][36].R1,
		&vk.Lines[0][0][35].R0,
		&vk.Lines[0][0][35].R1,
		&vk.Lines[0][0][34].R0,
		&vk.Lines[0][0][34].R1,
		&vk.Lines[0][0][33].R0,
		&vk.Lines[0][0][33].R1,
		&vk.Lines[0][0][32].R0,
		&vk.Lines[0][0][32].R1,
		&vk.Lines[0][0][31].R0,
		&vk.Lines[0][0][31].R1,
		&vk.Lines[0][0][30].R0,
		&vk.Lines[0][0][30].R1,
		&vk.Lines[0][0][29].R0,
		&vk.Lines[0][0][29].R1,
		&vk.Lines[0][0][28].R0,
		&vk.Lines[0][0][28].R1,
		&vk.Lines[0][0][27].R0,
		&vk.Lines[0][0][27].R1,
		&vk.Lines[0][0][26].R0,
		&vk.Lines[0][0][26].R1,
		&vk.Lines[0][0][25].R0,
		&vk.Lines[0][0][25].R1,
		&vk.Lines[0][0][24].R0,
		&vk.Lines[0][0][24].R1,
		&vk.Lines[0][0][23].R0,
		&vk.Lines[0][0][23].R1,
		&vk.Lines[0][0][22].R0,
		&vk.Lines[0][0][22].R1,
		&vk.Lines[0][0][21].R0,
		&vk.Lines[0][0][21].R1,
		&vk.Lines[0][0][20].R0,
		&vk.Lines[0][0][20].R1,
		&vk.Lines[0][0][19].R0,
		&vk.Lines[0][0][19].R1,
		&vk.Lines[0][0][18].R0,
		&vk.Lines[0][0][18].R1,
		&vk.Lines[0][0][17].R0,
		&vk.Lines[0][0][17].R1,
		&vk.Lines[0][0][16].R0,
		&vk.Lines[0][0][16].R1,
		&vk.Lines[0][0][15].R0,
		&vk.Lines[0][0][15].R1,
		&vk.Lines[0][0][14].R0,
		&vk.Lines[0][0][14].R1,
		&vk.Lines[0][0][13].R0,
		&vk.Lines[0][0][13].R1,
		&vk.Lines[0][0][12].R0,
		&vk.Lines[0][0][12].R1,
		&vk.Lines[0][0][11].R0,
		&vk.Lines[0][0][11].R1,
		&vk.Lines[0][0][10].R0,
		&vk.Lines[0][0][10].R1,
		&vk.Lines[0][0][9].R0,
		&vk.Lines[0][0][9].R1,
		&vk.Lines[0][0][8].R0,
		&vk.Lines[0][0][8].R1,
		&vk.Lines[0][0][7].R0,
		&vk.Lines[0][0][7].R1,
		&vk.Lines[0][0][6].R0,
		&vk.Lines[0][0][6].R1,
		&vk.Lines[0][0][5].R0,
		&vk.Lines[0][0][5].R1,
		&vk.Lines[0][0][4].R0,
		&vk.Lines[0][0][4].R1,
		&vk.Lines[0][0][3].R0,
		&vk.Lines[0][0][3].R1,
		&vk.Lines[0][0][2].R0,
		&vk.Lines[0][0][2].R1,
		&vk.Lines[0][0][1].R0,
		&vk.Lines[0][0][1].R1,
		&vk.Lines[0][0][0].R0,
		&vk.Lines[0][0][0].R1,
		&vk.Lines[0][1][62].R0,
		&vk.Lines[0][1][62].R1,
		&vk.Lines[0][1][61].R0,
		&vk.Lines[0][1][61].R1,
		&vk.Lines[0][1][60].R0,
		&vk.Lines[0][1][60].R1,
		&vk.Lines[0][1][59].R0,
		&vk.Lines[0][1][59].R1,
		&vk.Lines[0][1][58].R0,
		&vk.Lines[0][1][58].R1,
		&vk.Lines[0][1][57].R0,
		&vk.Lines[0][1][57].R1,
		&vk.Lines[0][1][56].R0,
		&vk.Lines[0][1][56].R1,
		&vk.Lines[0][1][55].R0,
		&vk.Lines[0][1][55].R1,
		&vk.Lines[0][1][54].R0,
		&vk.Lines[0][1][54].R1,
		&vk.Lines[0][1][53].R0,
		&vk.Lines[0][1][53].R1,
		&vk.Lines[0][1][52].R0,
		&vk.Lines[0][1][52].R1,
		&vk.Lines[0][1][51].R0,
		&vk.Lines[0][1][51].R1,
		&vk.Lines[0][1][50].R0,
		&vk.Lines[0][1][50].R1,
		&vk.Lines[0][1][49].R0,
		&vk.Lines[0][1][49].R1,
		&vk.Lines[0][1][48].R0,
		&vk.Lines[0][1][48].R1,
		&vk.Lines[0][1][47].R0,
		&vk.Lines[0][1][47].R1,
		&vk.Lines[0][1][46].R0,
		&vk.Lines[0][1][46].R1,
		&vk.Lines[0][1][45].R0,
		&vk.Lines[0][1][45].R1,
		&vk.Lines[0][1][44].R0,
		&vk.Lines[0][1][44].R1,
		&vk.Lines[0][1][43].R0,
		&vk.Lines[0][1][43].R1,
		&vk.Lines[0][1][42].R0,
		&vk.Lines[0][1][42].R1,
		&vk.Lines[0][1][41].R0,
		&vk.Lines[0][1][41].R1,
		&vk.Lines[0][1][40].R0,
		&vk.Lines[0][1][40].R1,
		&vk.Lines[0][1][39].R0,
		&vk.Lines[0][1][39].R1,
		&vk.Lines[0][1][38].R0,
		&vk.Lines[0][1][38].R1,
		&vk.Lines[0][1][37].R0,
		&vk.Lines[0][1][37].R1,
		&vk.Lines[0][1][36].R0,
		&vk.Lines[0][1][36].R1,
		&vk.Lines[0][1][35].R0,
		&vk.Lines[0][1][35].R1,
		&vk.Lines[0][1][34].R0,
		&vk.Lines[0][1][34].R1,
		&vk.Lines[0][1][33].R0,
		&vk.Lines[0][1][33].R1,
		&vk.Lines[0][1][32].R0,
		&vk.Lines[0][1][32].R1,
		&vk.Lines[0][1][31].R0,
		&vk.Lines[0][1][31].R1,
		&vk.Lines[0][1][30].R0,
		&vk.Lines[0][1][30].R1,
		&vk.Lines[0][1][29].R0,
		&vk.Lines[0][1][29].R1,
		&vk.Lines[0][1][28].R0,
		&vk.Lines[0][1][28].R1,
		&vk.Lines[0][1][27].R0,
		&vk.Lines[0][1][27].R1,
		&vk.Lines[0][1][26].R0,
		&vk.Lines[0][1][26].R1,
		&vk.Lines[0][1][25].R0,
		&vk.Lines[0][1][25].R1,
		&vk.Lines[0][1][24].R0,
		&vk.Lines[0][1][24].R1,
		&vk.Lines[0][1][23].R0,
		&vk.Lines[0][1][23].R1,
		&vk.Lines[0][1][22].R0,
		&vk.Lines[0][1][22].R1,
		&vk.Lines[0][1][21].R0,
		&vk.Lines[0][1][21].R1,
		&vk.Lines[0][1][20].R0,
		&vk.Lines[0][1][20].R1,
		&vk.Lines[0][1][19].R0,
		&vk.Lines[0][1][19].R1,
		&vk.Lines[0][1][18].R0,
		&vk.Lines[0][1][18].R1,
		&vk.Lines[0][1][17].R0,
		&vk.Lines[0][1][17].R1,
		&vk.Lines[0][1][16].R0,
		&vk.Lines[0][1][16].R1,
		&vk.Lines[0][1][15].R0,
		&vk.Lines[0][1][15].R1,
		&vk.Lines[0][1][14].R0,
		&vk.Lines[0][1][14].R1,
		&vk.Lines[0][1][13].R0,
		&vk.Lines[0][1][13].R1,
		&vk.Lines[0][1][12].R0,
		&vk.Lines[0][1][12].R1,
		&vk.Lines[0][1][11].R0,
		&vk.Lines[0][1][11].R1,
		&vk.Lines[0][1][10].R0,
		&vk.Lines[0][1][10].R1,
		&vk.Lines[0][1][9].R0,
		&vk.Lines[0][1][9].R1,
		&vk.Lines[0][1][8].R0,
		&vk.Lines[0][1][8].R1,
		&vk.Lines[0][1][7].R0,
		&vk.Lines[0][1][7].R1,
		&vk.Lines[0][1][6].R0,
		&vk.Lines[0][1][6].R1,
		&vk.Lines[0][1][5].R0,
		&vk.Lines[0][1][5].R1,
		&vk.Lines[0][1][4].R0,
		&vk.Lines[0][1][4].R1,
		&vk.Lines[0][1][3].R0,
		&vk.Lines[0][1][3].R1,
		&vk.Lines[0][1][2].R0,
		&vk.Lines[0][1][2].R1,
		&vk.Lines[0][1][1].R0,
		&vk.Lines[0][1][1].R1,
		&vk.Lines[0][1][0].R0,
		&vk.Lines[0][1][0].R1,
		&vk.Lines[1][0][62].R0,
		&vk.Lines[1][0][62].R1,
		&vk.Lines[1][0][61].R0,
		&vk.Lines[1][0][61].R1,
		&vk.Lines[1][0][60].R0,
		&vk.Lines[1][0][60].R1,
		&vk.Lines[1][0][59].R0,
		&vk.Lines[1][0][59].R1,
		&vk.Lines[1][0][58].R0,
		&vk.Lines[1][0][58].R1,
		&vk.Lines[1][0][57].R0,
		&vk.Lines[1][0][57].R1,
		&vk.Lines[1][0][56].R0,
		&vk.Lines[1][0][56].R1,
		&vk.Lines[1][0][55].R0,
		&vk.Lines[1][0][55].R1,
		&vk.Lines[1][0][54].R0,
		&vk.Lines[1][0][54].R1,
		&vk.Lines[1][0][53].R0,
		&vk.Lines[1][0][53].R1,
		&vk.Lines[1][0][52].R0,
		&vk.Lines[1][0][52].R1,
		&vk.Lines[1][0][51].R0,
		&vk.Lines[1][0][51].R1,
		&vk.Lines[1][0][50].R0,
		&vk.Lines[1][0][50].R1,
		&vk.Lines[1][0][49].R0,
		&vk.Lines[1][0][49].R1,
		&vk.Lines[1][0][48].R0,
		&vk.Lines[1][0][48].R1,
		&vk.Lines[1][0][47].R0,
		&vk.Lines[1][0][47].R1,
		&vk.Lines[1][0][46].R0,
		&vk.Lines[1][0][46].R1,
		&vk.Lines[1][0][45].R0,
		&vk.Lines[1][0][45].R1,
		&vk.Lines[1][0][44].R0,
		&vk.Lines[1][0][44].R1,
		&vk.Lines[1][0][43].R0,
		&vk.Lines[1][0][43].R1,
		&vk.Lines[1][0][42].R0,
		&vk.Lines[1][0][42].R1,
		&vk.Lines[1][0][41].R0,
		&vk.Lines[1][0][41].R1,
		&vk.Lines[1][0][40].R0,
		&vk.Lines[1][0][40].R1,
		&vk.Lines[1][0][39].R0,
		&vk.Lines[1][0][39].R1,
		&vk.Lines[1][0][38].R0,
		&vk.Lines[1][0][38].R1,
		&vk.Lines[1][0][37].R0,
		&vk.Lines[1][0][37].R1,
		&vk.Lines[1][0][36].R0,
		&vk.Lines[1][0][36].R1,
		&vk.Lines[1][0][35].R0,
		&vk.Lines[1][0][35].R1,
		&vk.Lines[1][0][34].R0,
		&vk.Lines[1][0][34].R1,
		&vk.Lines[1][0][33].R0,
		&vk.Lines[1][0][33].R1,
		&vk.Lines[1][0][32].R0,
		&vk.Lines[1][0][32].R1,
		&vk.Lines[1][0][31].R0,
		&vk.Lines[1][0][31].R1,
		&vk.Lines[1][0][30].R0,
		&vk.Lines[1][0][30].R1,
		&vk.Lines[1][0][29].R0,
		&vk.Lines[1][0][29].R1,
		&vk.Lines[1][0][28].R0,
		&vk.Lines[1][0][28].R1,
		&vk.Lines[1][0][27].R0,
		&vk.Lines[1][0][27].R1,
		&vk.Lines[1][0][26].R0,
		&vk.Lines[1][0][26].R1,
		&vk.Lines[1][0][25].R0,
		&vk.Lines[1][0][25].R1,
		&vk.Lines[1][0][24].R0,
		&vk.Lines[1][0][24].R1,
		&vk.Lines[1][0][23].R0,
		&vk.Lines[1][0][23].R1,
		&vk.Lines[1][0][22].R0,
		&vk.Lines[1][0][22].R1,
		&vk.Lines[1][0][21].R0,
		&vk.Lines[1][0][21].R1,
		&vk.Lines[1][0][20].R0,
		&vk.Lines[1][0][20].R1,
		&vk.Lines[1][0][19].R0,
		&vk.Lines[1][0][19].R1,
		&vk.Lines[1][0][18].R0,
		&vk.Lines[1][0][18].R1,
		&vk.Lines[1][0][17].R0,
		&vk.Lines[1][0][17].R1,
		&vk.Lines[1][0][16].R0,
		&vk.Lines[1][0][16].R1,
		&vk.Lines[1][0][15].R0,
		&vk.Lines[1][0][15].R1,
		&vk.Lines[1][0][14].R0,
		&vk.Lines[1][0][14].R1,
		&vk.Lines[1][0][13].R0,
		&vk.Lines[1][0][13].R1,
		&vk.Lines[1][0][12].R0,
		&vk.Lines[1][0][12].R1,
		&vk.Lines[1][0][11].R0,
		&vk.Lines[1][0][11].R1,
		&vk.Lines[1][0][10].R0,
		&vk.Lines[1][0][10].R1,
		&vk.Lines[1][0][9].R0,
		&vk.Lines[1][0][9].R1,
		&vk.Lines[1][0][8].R0,
		&vk.Lines[1][0][8].R1,
		&vk.Lines[1][0][7].R0,
		&vk.Lines[1][0][7].R1,
		&vk.Lines[1][0][6].R0,
		&vk.Lines[1][0][6].R1,
		&vk.Lines[1][0][5].R0,
		&vk.Lines[1][0][5].R1,
		&vk.Lines[1][0][4].R0,
		&vk.Lines[1][0][4].R1,
		&vk.Lines[1][0][3].R0,
		&vk.Lines[1][0][3].R1,
		&vk.Lines[1][0][2].R0,
		&vk.Lines[1][0][2].R1,
		&vk.Lines[1][0][1].R0,
		&vk.Lines[1][0][1].R1,
		&vk.Lines[1][0][0].R0,
		&vk.Lines[1][0][0].R1,
		&vk.Lines[1][1][62].R0,
		&vk.Lines[1][1][62].R1,
		&vk.Lines[1][1][61].R0,
		&vk.Lines[1][1][61].R1,
		&vk.Lines[1][1][60].R0,
		&vk.Lines[1][1][60].R1,
		&vk.Lines[1][1][59].R0,
		&vk.Lines[1][1][59].R1,
		&vk.Lines[1][1][58].R0,
		&vk.Lines[1][1][58].R1,
		&vk.Lines[1][1][57].R0,
		&vk.Lines[1][1][57].R1,
		&vk.Lines[1][1][56].R0,
		&vk.Lines[1][1][56].R1,
		&vk.Lines[1][1][55].R0,
		&vk.Lines[1][1][55].R1,
		&vk.Lines[1][1][54].R0,
		&vk.Lines[1][1][54].R1,
		&vk.Lines[1][1][53].R0,
		&vk.Lines[1][1][53].R1,
		&vk.Lines[1][1][52].R0,
		&vk.Lines[1][1][52].R1,
		&vk.Lines[1][1][51].R0,
		&vk.Lines[1][1][51].R1,
		&vk.Lines[1][1][50].R0,
		&vk.Lines[1][1][50].R1,
		&vk.Lines[1][1][49].R0,
		&vk.Lines[1][1][49].R1,
		&vk.Lines[1][1][48].R0,
		&vk.Lines[1][1][48].R1,
		&vk.Lines[1][1][47].R0,
		&vk.Lines[1][1][47].R1,
		&vk.Lines[1][1][46].R0,
		&vk.Lines[1][1][46].R1,
		&vk.Lines[1][1][45].R0,
		&vk.Lines[1][1][45].R1,
		&vk.Lines[1][1][44].R0,
		&vk.Lines[1][1][44].R1,
		&vk.Lines[1][1][43].R0,
		&vk.Lines[1][1][43].R1,
		&vk.Lines[1][1][42].R0,
		&vk.Lines[1][1][42].R1,
		&vk.Lines[1][1][41].R0,
		&vk.Lines[1][1][41].R1,
		&vk.Lines[1][1][40].R0,
		&vk.Lines[1][1][40].R1,
		&vk.Lines[1][1][39].R0,
		&vk.Lines[1][1][39].R1,
		&vk.Lines[1][1][38].R0,
		&vk.Lines[1][1][38].R1,
		&vk.Lines[1][1][37].R0,
		&vk.Lines[1][1][37].R1,
		&vk.Lines[1][1][36].R0,
		&vk.Lines[1][1][36].R1,
		&vk.Lines[1][1][35].R0,
		&vk.Lines[1][1][35].R1,
		&vk.Lines[1][1][34].R0,
		&vk.Lines[1][1][34].R1,
		&vk.Lines[1][1][33].R0,
		&vk.Lines[1][1][33].R1,
		&vk.Lines[1][1][32].R0,
		&vk.Lines[1][1][32].R1,
		&vk.Lines[1][1][31].R0,
		&vk.Lines[1][1][31].R1,
		&vk.Lines[1][1][30].R0,
		&vk.Lines[1][1][30].R1,
		&vk.Lines[1][1][29].R0,
		&vk.Lines[1][1][29].R1,
		&vk.Lines[1][1][28].R0,
		&vk.Lines[1][1][28].R1,
		&vk.Lines[1][1][27].R0,
		&vk.Lines[1][1][27].R1,
		&vk.Lines[1][1][26].R0,
		&vk.Lines[1][1][26].R1,
		&vk.Lines[1][1][25].R0,
		&vk.Lines[1][1][25].R1,
		&vk.Lines[1][1][24].R0,
		&vk.Lines[1][1][24].R1,
		&vk.Lines[1][1][23].R0,
		&vk.Lines[1][1][23].R1,
		&vk.Lines[1][1][22].R0,
		&vk.Lines[1][1][22].R1,
		&vk.Lines[1][1][21].R0,
		&vk.Lines[1][1][21].R1,
		&vk.Lines[1][1][20].R0,
		&vk.Lines[1][1][20].R1,
		&vk.Lines[1][1][19].R0,
		&vk.Lines[1][1][19].R1,
		&vk.Lines[1][1][18].R0,
		&vk.Lines[1][1][18].R1,
		&vk.Lines[1][1][17].R0,
		&vk.Lines[1][1][17].R1,
		&vk.Lines[1][1][16].R0,
		&vk.Lines[1][1][16].R1,
		&vk.Lines[1][1][15].R0,
		&vk.Lines[1][1][15].R1,
		&vk.Lines[1][1][14].R0,
		&vk.Lines[1][1][14].R1,
		&vk.Lines[1][1][13].R0,
		&vk.Lines[1][1][13].R1,
		&vk.Lines[1][1][12].R0,
		&vk.Lines[1][1][12].R1,
		&vk.Lines[1][1][11].R0,
		&vk.Lines[1][1][11].R1,
		&vk.Lines[1][1][10].R0,
		&vk.Lines[1][1][10].R1,
		&vk.Lines[1][1][9].R0,
		&vk.Lines[1][1][9].R1,
		&vk.Lines[1][1][8].R0,
		&vk.Lines[1][1][8].R1,
		&vk.Lines[1][1][7].R0,
		&vk.Lines[1][1][7].R1,
		&vk.Lines[1][1][6].R0,
		&vk.Lines[1][1][6].R1,
		&vk.Lines[1][1][5].R0,
		&vk.Lines[1][1][5].R1,
		&vk.Lines[1][1][4].R0,
		&vk.Lines[1][1][4].R1,
		&vk.Lines[1][1][3].R0,
		&vk.Lines[1][1][3].R1,
		&vk.Lines[1][1][2].R0,
		&vk.Lines[1][1][2].R1,
		&vk.Lines[1][1][1].R0,
		&vk.Lines[1][1][1].R1,
		&vk.Lines[1][1][0].R0,
		&vk.Lines[1][1][0].R1,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// WriteTo writes binary encoding of the entire SRS
func (srs *SRS) WriteTo(w io.Writer) (int64, error) {
	// encode the SRS
	var pn, vn int64
	var err error
	if pn, err = srs.Pk.WriteTo(w); err != nil {
		return pn, err
	}
	vn, err = srs.Vk.WriteTo(w)
	return pn + vn, err
}

// WriteRawTo writes binary encoding of the entire SRS without point compression
func (srs *SRS) WriteRawTo(w io.Writer) (int64, error) {
	// encode the SRS
	var pn, vn int64
	var err error
	if pn, err = srs.Pk.WriteRawTo(w); err != nil {
		return pn, err
	}
	vn, err = srs.Vk.WriteRawTo(w)
	return pn + vn, err
}

// ReadFrom decodes ProvingKey data from reader.
func (pk *ProvingKey) ReadFrom(r io.Reader) (int64, error) {
	// decode the ProvingKey
	dec := bls12378.NewDecoder(r)
	if err := dec.Decode(&pk.G1); err != nil {
		return dec.BytesRead(), err
	}
	return dec.BytesRead(), nil
}

// UnsafeReadFrom decodes ProvingKey data from reader without checking
// that point are in the correct subgroup.
func (pk *ProvingKey) UnsafeReadFrom(r io.Reader) (int64, error) {
	// decode the ProvingKey
	dec := bls12378.NewDecoder(r, bls12378.NoSubgroupChecks())
	if err := dec.Decode(&pk.G1); err != nil {
		return dec.BytesRead(), err
	}
	return dec.BytesRead(), nil
}

// ReadFrom decodes VerifyingKey data from reader.
func (vk *VerifyingKey) ReadFrom(r io.Reader) (int64, error) {
	// decode the VerifyingKey
	dec := bls12378.NewDecoder(r)

	toDecode := []interface{}{
		&vk.G2[0],
		&vk.G2[1],
		&vk.G1,
		&vk.Lines[0][0][62].R0,
		&vk.Lines[0][0][62].R1,
		&vk.Lines[0][0][61].R0,
		&vk.Lines[0][0][61].R1,
		&vk.Lines[0][0][60].R0,
		&vk.Lines[0][0][60].R1,
		&vk.Lines[0][0][59].R0,
		&vk.Lines[0][0][59].R1,
		&vk.Lines[0][0][58].R0,
		&vk.Lines[0][0][58].R1,
		&vk.Lines[0][0][57].R0,
		&vk.Lines[0][0][57].R1,
		&vk.Lines[0][0][56].R0,
		&vk.Lines[0][0][56].R1,
		&vk.Lines[0][0][55].R0,
		&vk.Lines[0][0][55].R1,
		&vk.Lines[0][0][54].R0,
		&vk.Lines[0][0][54].R1,
		&vk.Lines[0][0][53].R0,
		&vk.Lines[0][0][53].R1,
		&vk.Lines[0][0][52].R0,
		&vk.Lines[0][0][52].R1,
		&vk.Lines[0][0][51].R0,
		&vk.Lines[0][0][51].R1,
		&vk.Lines[0][0][50].R0,
		&vk.Lines[0][0][50].R1,
		&vk.Lines[0][0][49].R0,
		&vk.Lines[0][0][49].R1,
		&vk.Lines[0][0][48].R0,
		&vk.Lines[0][0][48].R1,
		&vk.Lines[0][0][47].R0,
		&vk.Lines[0][0][47].R1,
		&vk.Lines[0][0][46].R0,
		&vk.Lines[0][0][46].R1,
		&vk.Lines[0][0][45].R0,
		&vk.Lines[0][0][45].R1,
		&vk.Lines[0][0][44].R0,
		&vk.Lines[0][0][44].R1,
		&vk.Lines[0][0][43].R0,
		&vk.Lines[0][0][43].R1,
		&vk.Lines[0][0][42].R0,
		&vk.Lines[0][0][42].R1,
		&vk.Lines[0][0][41].R0,
		&vk.Lines[0][0][41].R1,
		&vk.Lines[0][0][40].R0,
		&vk.Lines[0][0][40].R1,
		&vk.Lines[0][0][39].R0,
		&vk.Lines[0][0][39].R1,
		&vk.Lines[0][0][38].R0,
		&vk.Lines[0][0][38].R1,
		&vk.Lines[0][0][37].R0,
		&vk.Lines[0][0][37].R1,
		&vk.Lines[0][0][36].R0,
		&vk.Lines[0][0][36].R1,
		&vk.Lines[0][0][35].R0,
		&vk.Lines[0][0][35].R1,
		&vk.Lines[0][0][34].R0,
		&vk.Lines[0][0][34].R1,
		&vk.Lines[0][0][33].R0,
		&vk.Lines[0][0][33].R1,
		&vk.Lines[0][0][32].R0,
		&vk.Lines[0][0][32].R1,
		&vk.Lines[0][0][31].R0,
		&vk.Lines[0][0][31].R1,
		&vk.Lines[0][0][30].R0,
		&vk.Lines[0][0][30].R1,
		&vk.Lines[0][0][29].R0,
		&vk.Lines[0][0][29].R1,
		&vk.Lines[0][0][28].R0,
		&vk.Lines[0][0][28].R1,
		&vk.Lines[0][0][27].R0,
		&vk.Lines[0][0][27].R1,
		&vk.Lines[0][0][26].R0,
		&vk.Lines[0][0][26].R1,
		&vk.Lines[0][0][25].R0,
		&vk.Lines[0][0][25].R1,
		&vk.Lines[0][0][24].R0,
		&vk.Lines[0][0][24].R1,
		&vk.Lines[0][0][23].R0,
		&vk.Lines[0][0][23].R1,
		&vk.Lines[0][0][22].R0,
		&vk.Lines[0][0][22].R1,
		&vk.Lines[0][0][21].R0,
		&vk.Lines[0][0][21].R1,
		&vk.Lines[0][0][20].R0,
		&vk.Lines[0][0][20].R1,
		&vk.Lines[0][0][19].R0,
		&vk.Lines[0][0][19].R1,
		&vk.Lines[0][0][18].R0,
		&vk.Lines[0][0][18].R1,
		&vk.Lines[0][0][17].R0,
		&vk.Lines[0][0][17].R1,
		&vk.Lines[0][0][16].R0,
		&vk.Lines[0][0][16].R1,
		&vk.Lines[0][0][15].R0,
		&vk.Lines[0][0][15].R1,
		&vk.Lines[0][0][14].R0,
		&vk.Lines[0][0][14].R1,
		&vk.Lines[0][0][13].R0,
		&vk.Lines[0][0][13].R1,
		&vk.Lines[0][0][12].R0,
		&vk.Lines[0][0][12].R1,
		&vk.Lines[0][0][11].R0,
		&vk.Lines[0][0][11].R1,
		&vk.Lines[0][0][10].R0,
		&vk.Lines[0][0][10].R1,
		&vk.Lines[0][0][9].R0,
		&vk.Lines[0][0][9].R1,
		&vk.Lines[0][0][8].R0,
		&vk.Lines[0][0][8].R1,
		&vk.Lines[0][0][7].R0,
		&vk.Lines[0][0][7].R1,
		&vk.Lines[0][0][6].R0,
		&vk.Lines[0][0][6].R1,
		&vk.Lines[0][0][5].R0,
		&vk.Lines[0][0][5].R1,
		&vk.Lines[0][0][4].R0,
		&vk.Lines[0][0][4].R1,
		&vk.Lines[0][0][3].R0,
		&vk.Lines[0][0][3].R1,
		&vk.Lines[0][0][2].R0,
		&vk.Lines[0][0][2].R1,
		&vk.Lines[0][0][1].R0,
		&vk.Lines[0][0][1].R1,
		&vk.Lines[0][0][0].R0,
		&vk.Lines[0][0][0].R1,
		&vk.Lines[0][1][62].R0,
		&vk.Lines[0][1][62].R1,
		&vk.Lines[0][1][61].R0,
		&vk.Lines[0][1][61].R1,
		&vk.Lines[0][1][60].R0,
		&vk.Lines[0][1][60].R1,
		&vk.Lines[0][1][59].R0,
		&vk.Lines[0][1][59].R1,
		&vk.Lines[0][1][58].R0,
		&vk.Lines[0][1][58].R1,
		&vk.Lines[0][1][57].R0,
		&vk.Lines[0][1][57].R1,
		&vk.Lines[0][1][56].R0,
		&vk.Lines[0][1][56].R1,
		&vk.Lines[0][1][55].R0,
		&vk.Lines[0][1][55].R1,
		&vk.Lines[0][1][54].R0,
		&vk.Lines[0][1][54].R1,
		&vk.Lines[0][1][53].R0,
		&vk.Lines[0][1][53].R1,
		&vk.Lines[0][1][52].R0,
		&vk.Lines[0][1][52].R1,
		&vk.Lines[0][1][51].R0,
		&vk.Lines[0][1][51].R1,
		&vk.Lines[0][1][50].R0,
		&vk.Lines[0][1][50].R1,
		&vk.Lines[0][1][49].R0,
		&vk.Lines[0][1][49].R1,
		&vk.Lines[0][1][48].R0,
		&vk.Lines[0][1][48].R1,
		&vk.Lines[0][1][47].R0,
		&vk.Lines[0][1][47].R1,
		&vk.Lines[0][1][46].R0,
		&vk.Lines[0][1][46].R1,
		&vk.Lines[0][1][45].R0,
		&vk.Lines[0][1][45].R1,
		&vk.Lines[0][1][44].R0,
		&vk.Lines[0][1][44].R1,
		&vk.Lines[0][1][43].R0,
		&vk.Lines[0][1][43].R1,
		&vk.Lines[0][1][42].R0,
		&vk.Lines[0][1][42].R1,
		&vk.Lines[0][1][41].R0,
		&vk.Lines[0][1][41].R1,
		&vk.Lines[0][1][40].R0,
		&vk.Lines[0][1][40].R1,
		&vk.Lines[0][1][39].R0,
		&vk.Lines[0][1][39].R1,
		&vk.Lines[0][1][38].R0,
		&vk.Lines[0][1][38].R1,
		&vk.Lines[0][1][37].R0,
		&vk.Lines[0][1][37].R1,
		&vk.Lines[0][1][36].R0,
		&vk.Lines[0][1][36].R1,
		&vk.Lines[0][1][35].R0,
		&vk.Lines[0][1][35].R1,
		&vk.Lines[0][1][34].R0,
		&vk.Lines[0][1][34].R1,
		&vk.Lines[0][1][33].R0,
		&vk.Lines[0][1][33].R1,
		&vk.Lines[0][1][32].R0,
		&vk.Lines[0][1][32].R1,
		&vk.Lines[0][1][31].R0,
		&vk.Lines[0][1][31].R1,
		&vk.Lines[0][1][30].R0,
		&vk.Lines[0][1][30].R1,
		&vk.Lines[0][1][29].R0,
		&vk.Lines[0][1][29].R1,
		&vk.Lines[0][1][28].R0,
		&vk.Lines[0][1][28].R1,
		&vk.Lines[0][1][27].R0,
		&vk.Lines[0][1][27].R1,
		&vk.Lines[0][1][26].R0,
		&vk.Lines[0][1][26].R1,
		&vk.Lines[0][1][25].R0,
		&vk.Lines[0][1][25].R1,
		&vk.Lines[0][1][24].R0,
		&vk.Lines[0][1][24].R1,
		&vk.Lines[0][1][23].R0,
		&vk.Lines[0][1][23].R1,
		&vk.Lines[0][1][22].R0,
		&vk.Lines[0][1][22].R1,
		&vk.Lines[0][1][21].R0,
		&vk.Lines[0][1][21].R1,
		&vk.Lines[0][1][20].R0,
		&vk.Lines[0][1][20].R1,
		&vk.Lines[0][1][19].R0,
		&vk.Lines[0][1][19].R1,
		&vk.Lines[0][1][18].R0,
		&vk.Lines[0][1][18].R1,
		&vk.Lines[0][1][17].R0,
		&vk.Lines[0][1][17].R1,
		&vk.Lines[0][1][16].R0,
		&vk.Lines[0][1][16].R1,
		&vk.Lines[0][1][15].R0,
		&vk.Lines[0][1][15].R1,
		&vk.Lines[0][1][14].R0,
		&vk.Lines[0][1][14].R1,
		&vk.Lines[0][1][13].R0,
		&vk.Lines[0][1][13].R1,
		&vk.Lines[0][1][12].R0,
		&vk.Lines[0][1][12].R1,
		&vk.Lines[0][1][11].R0,
		&vk.Lines[0][1][11].R1,
		&vk.Lines[0][1][10].R0,
		&vk.Lines[0][1][10].R1,
		&vk.Lines[0][1][9].R0,
		&vk.Lines[0][1][9].R1,
		&vk.Lines[0][1][8].R0,
		&vk.Lines[0][1][8].R1,
		&vk.Lines[0][1][7].R0,
		&vk.Lines[0][1][7].R1,
		&vk.Lines[0][1][6].R0,
		&vk.Lines[0][1][6].R1,
		&vk.Lines[0][1][5].R0,
		&vk.Lines[0][1][5].R1,
		&vk.Lines[0][1][4].R0,
		&vk.Lines[0][1][4].R1,
		&vk.Lines[0][1][3].R0,
		&vk.Lines[0][1][3].R1,
		&vk.Lines[0][1][2].R0,
		&vk.Lines[0][1][2].R1,
		&vk.Lines[0][1][1].R0,
		&vk.Lines[0][1][1].R1,
		&vk.Lines[0][1][0].R0,
		&vk.Lines[0][1][0].R1,
		&vk.Lines[1][0][62].R0,
		&vk.Lines[1][0][62].R1,
		&vk.Lines[1][0][61].R0,
		&vk.Lines[1][0][61].R1,
		&vk.Lines[1][0][60].R0,
		&vk.Lines[1][0][60].R1,
		&vk.Lines[1][0][59].R0,
		&vk.Lines[1][0][59].R1,
		&vk.Lines[1][0][58].R0,
		&vk.Lines[1][0][58].R1,
		&vk.Lines[1][0][57].R0,
		&vk.Lines[1][0][57].R1,
		&vk.Lines[1][0][56].R0,
		&vk.Lines[1][0][56].R1,
		&vk.Lines[1][0][55].R0,
		&vk.Lines[1][0][55].R1,
		&vk.Lines[1][0][54].R0,
		&vk.Lines[1][0][54].R1,
		&vk.Lines[1][0][53].R0,
		&vk.Lines[1][0][53].R1,
		&vk.Lines[1][0][52].R0,
		&vk.Lines[1][0][52].R1,
		&vk.Lines[1][0][51].R0,
		&vk.Lines[1][0][51].R1,
		&vk.Lines[1][0][50].R0,
		&vk.Lines[1][0][50].R1,
		&vk.Lines[1][0][49].R0,
		&vk.Lines[1][0][49].R1,
		&vk.Lines[1][0][48].R0,
		&vk.Lines[1][0][48].R1,
		&vk.Lines[1][0][47].R0,
		&vk.Lines[1][0][47].R1,
		&vk.Lines[1][0][46].R0,
		&vk.Lines[1][0][46].R1,
		&vk.Lines[1][0][45].R0,
		&vk.Lines[1][0][45].R1,
		&vk.Lines[1][0][44].R0,
		&vk.Lines[1][0][44].R1,
		&vk.Lines[1][0][43].R0,
		&vk.Lines[1][0][43].R1,
		&vk.Lines[1][0][42].R0,
		&vk.Lines[1][0][42].R1,
		&vk.Lines[1][0][41].R0,
		&vk.Lines[1][0][41].R1,
		&vk.Lines[1][0][40].R0,
		&vk.Lines[1][0][40].R1,
		&vk.Lines[1][0][39].R0,
		&vk.Lines[1][0][39].R1,
		&vk.Lines[1][0][38].R0,
		&vk.Lines[1][0][38].R1,
		&vk.Lines[1][0][37].R0,
		&vk.Lines[1][0][37].R1,
		&vk.Lines[1][0][36].R0,
		&vk.Lines[1][0][36].R1,
		&vk.Lines[1][0][35].R0,
		&vk.Lines[1][0][35].R1,
		&vk.Lines[1][0][34].R0,
		&vk.Lines[1][0][34].R1,
		&vk.Lines[1][0][33].R0,
		&vk.Lines[1][0][33].R1,
		&vk.Lines[1][0][32].R0,
		&vk.Lines[1][0][32].R1,
		&vk.Lines[1][0][31].R0,
		&vk.Lines[1][0][31].R1,
		&vk.Lines[1][0][30].R0,
		&vk.Lines[1][0][30].R1,
		&vk.Lines[1][0][29].R0,
		&vk.Lines[1][0][29].R1,
		&vk.Lines[1][0][28].R0,
		&vk.Lines[1][0][28].R1,
		&vk.Lines[1][0][27].R0,
		&vk.Lines[1][0][27].R1,
		&vk.Lines[1][0][26].R0,
		&vk.Lines[1][0][26].R1,
		&vk.Lines[1][0][25].R0,
		&vk.Lines[1][0][25].R1,
		&vk.Lines[1][0][24].R0,
		&vk.Lines[1][0][24].R1,
		&vk.Lines[1][0][23].R0,
		&vk.Lines[1][0][23].R1,
		&vk.Lines[1][0][22].R0,
		&vk.Lines[1][0][22].R1,
		&vk.Lines[1][0][21].R0,
		&vk.Lines[1][0][21].R1,
		&vk.Lines[1][0][20].R0,
		&vk.Lines[1][0][20].R1,
		&vk.Lines[1][0][19].R0,
		&vk.Lines[1][0][19].R1,
		&vk.Lines[1][0][18].R0,
		&vk.Lines[1][0][18].R1,
		&vk.Lines[1][0][17].R0,
		&vk.Lines[1][0][17].R1,
		&vk.Lines[1][0][16].R0,
		&vk.Lines[1][0][16].R1,
		&vk.Lines[1][0][15].R0,
		&vk.Lines[1][0][15].R1,
		&vk.Lines[1][0][14].R0,
		&vk.Lines[1][0][14].R1,
		&vk.Lines[1][0][13].R0,
		&vk.Lines[1][0][13].R1,
		&vk.Lines[1][0][12].R0,
		&vk.Lines[1][0][12].R1,
		&vk.Lines[1][0][11].R0,
		&vk.Lines[1][0][11].R1,
		&vk.Lines[1][0][10].R0,
		&vk.Lines[1][0][10].R1,
		&vk.Lines[1][0][9].R0,
		&vk.Lines[1][0][9].R1,
		&vk.Lines[1][0][8].R0,
		&vk.Lines[1][0][8].R1,
		&vk.Lines[1][0][7].R0,
		&vk.Lines[1][0][7].R1,
		&vk.Lines[1][0][6].R0,
		&vk.Lines[1][0][6].R1,
		&vk.Lines[1][0][5].R0,
		&vk.Lines[1][0][5].R1,
		&vk.Lines[1][0][4].R0,
		&vk.Lines[1][0][4].R1,
		&vk.Lines[1][0][3].R0,
		&vk.Lines[1][0][3].R1,
		&vk.Lines[1][0][2].R0,
		&vk.Lines[1][0][2].R1,
		&vk.Lines[1][0][1].R0,
		&vk.Lines[1][0][1].R1,
		&vk.Lines[1][0][0].R0,
		&vk.Lines[1][0][0].R1,
		&vk.Lines[1][1][62].R0,
		&vk.Lines[1][1][62].R1,
		&vk.Lines[1][1][61].R0,
		&vk.Lines[1][1][61].R1,
		&vk.Lines[1][1][60].R0,
		&vk.Lines[1][1][60].R1,
		&vk.Lines[1][1][59].R0,
		&vk.Lines[1][1][59].R1,
		&vk.Lines[1][1][58].R0,
		&vk.Lines[1][1][58].R1,
		&vk.Lines[1][1][57].R0,
		&vk.Lines[1][1][57].R1,
		&vk.Lines[1][1][56].R0,
		&vk.Lines[1][1][56].R1,
		&vk.Lines[1][1][55].R0,
		&vk.Lines[1][1][55].R1,
		&vk.Lines[1][1][54].R0,
		&vk.Lines[1][1][54].R1,
		&vk.Lines[1][1][53].R0,
		&vk.Lines[1][1][53].R1,
		&vk.Lines[1][1][52].R0,
		&vk.Lines[1][1][52].R1,
		&vk.Lines[1][1][51].R0,
		&vk.Lines[1][1][51].R1,
		&vk.Lines[1][1][50].R0,
		&vk.Lines[1][1][50].R1,
		&vk.Lines[1][1][49].R0,
		&vk.Lines[1][1][49].R1,
		&vk.Lines[1][1][48].R0,
		&vk.Lines[1][1][48].R1,
		&vk.Lines[1][1][47].R0,
		&vk.Lines[1][1][47].R1,
		&vk.Lines[1][1][46].R0,
		&vk.Lines[1][1][46].R1,
		&vk.Lines[1][1][45].R0,
		&vk.Lines[1][1][45].R1,
		&vk.Lines[1][1][44].R0,
		&vk.Lines[1][1][44].R1,
		&vk.Lines[1][1][43].R0,
		&vk.Lines[1][1][43].R1,
		&vk.Lines[1][1][42].R0,
		&vk.Lines[1][1][42].R1,
		&vk.Lines[1][1][41].R0,
		&vk.Lines[1][1][41].R1,
		&vk.Lines[1][1][40].R0,
		&vk.Lines[1][1][40].R1,
		&vk.Lines[1][1][39].R0,
		&vk.Lines[1][1][39].R1,
		&vk.Lines[1][1][38].R0,
		&vk.Lines[1][1][38].R1,
		&vk.Lines[1][1][37].R0,
		&vk.Lines[1][1][37].R1,
		&vk.Lines[1][1][36].R0,
		&vk.Lines[1][1][36].R1,
		&vk.Lines[1][1][35].R0,
		&vk.Lines[1][1][35].R1,
		&vk.Lines[1][1][34].R0,
		&vk.Lines[1][1][34].R1,
		&vk.Lines[1][1][33].R0,
		&vk.Lines[1][1][33].R1,
		&vk.Lines[1][1][32].R0,
		&vk.Lines[1][1][32].R1,
		&vk.Lines[1][1][31].R0,
		&vk.Lines[1][1][31].R1,
		&vk.Lines[1][1][30].R0,
		&vk.Lines[1][1][30].R1,
		&vk.Lines[1][1][29].R0,
		&vk.Lines[1][1][29].R1,
		&vk.Lines[1][1][28].R0,
		&vk.Lines[1][1][28].R1,
		&vk.Lines[1][1][27].R0,
		&vk.Lines[1][1][27].R1,
		&vk.Lines[1][1][26].R0,
		&vk.Lines[1][1][26].R1,
		&vk.Lines[1][1][25].R0,
		&vk.Lines[1][1][25].R1,
		&vk.Lines[1][1][24].R0,
		&vk.Lines[1][1][24].R1,
		&vk.Lines[1][1][23].R0,
		&vk.Lines[1][1][23].R1,
		&vk.Lines[1][1][22].R0,
		&vk.Lines[1][1][22].R1,
		&vk.Lines[1][1][21].R0,
		&vk.Lines[1][1][21].R1,
		&vk.Lines[1][1][20].R0,
		&vk.Lines[1][1][20].R1,
		&vk.Lines[1][1][19].R0,
		&vk.Lines[1][1][19].R1,
		&vk.Lines[1][1][18].R0,
		&vk.Lines[1][1][18].R1,
		&vk.Lines[1][1][17].R0,
		&vk.Lines[1][1][17].R1,
		&vk.Lines[1][1][16].R0,
		&vk.Lines[1][1][16].R1,
		&vk.Lines[1][1][15].R0,
		&vk.Lines[1][1][15].R1,
		&vk.Lines[1][1][14].R0,
		&vk.Lines[1][1][14].R1,
		&vk.Lines[1][1][13].R0,
		&vk.Lines[1][1][13].R1,
		&vk.Lines[1][1][12].R0,
		&vk.Lines[1][1][12].R1,
		&vk.Lines[1][1][11].R0,
		&vk.Lines[1][1][11].R1,
		&vk.Lines[1][1][10].R0,
		&vk.Lines[1][1][10].R1,
		&vk.Lines[1][1][9].R0,
		&vk.Lines[1][1][9].R1,
		&vk.Lines[1][1][8].R0,
		&vk.Lines[1][1][8].R1,
		&vk.Lines[1][1][7].R0,
		&vk.Lines[1][1][7].R1,
		&vk.Lines[1][1][6].R0,
		&vk.Lines[1][1][6].R1,
		&vk.Lines[1][1][5].R0,
		&vk.Lines[1][1][5].R1,
		&vk.Lines[1][1][4].R0,
		&vk.Lines[1][1][4].R1,
		&vk.Lines[1][1][3].R0,
		&vk.Lines[1][1][3].R1,
		&vk.Lines[1][1][2].R0,
		&vk.Lines[1][1][2].R1,
		&vk.Lines[1][1][1].R0,
		&vk.Lines[1][1][1].R1,
		&vk.Lines[1][1][0].R0,
		&vk.Lines[1][1][0].R1,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	return dec.BytesRead(), nil
}

// ReadFrom decodes SRS data from reader.
func (srs *SRS) ReadFrom(r io.Reader) (int64, error) {
	// decode the VerifyingKey
	var pn, vn int64
	var err error
	if pn, err = srs.Pk.ReadFrom(r); err != nil {
		return pn, err
	}
	vn, err = srs.Vk.ReadFrom(r)
	return pn + vn, err
}

// WriteTo writes binary encoding of a OpeningProof
func (proof *OpeningProof) WriteTo(w io.Writer) (int64, error) {
	enc := bls12378.NewEncoder(w)

	toEncode := []interface{}{
		&proof.H,
		&proof.ClaimedValue,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// ReadFrom decodes OpeningProof data from reader.
func (proof *OpeningProof) ReadFrom(r io.Reader) (int64, error) {
	dec := bls12378.NewDecoder(r)

	toDecode := []interface{}{
		&proof.H,
		&proof.ClaimedValue,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	return dec.BytesRead(), nil
}

// WriteTo writes binary encoding of a BatchOpeningProof
func (proof *BatchOpeningProof) WriteTo(w io.Writer) (int64, error) {
	enc := bls12378.NewEncoder(w)

	toEncode := []interface{}{
		&proof.H,
		proof.ClaimedValues,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// ReadFrom decodes BatchOpeningProof data from reader.
func (proof *BatchOpeningProof) ReadFrom(r io.Reader) (int64, error) {
	dec := bls12378.NewDecoder(r)
	toDecode := []interface{}{
		&proof.H,
		&proof.ClaimedValues,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	return dec.BytesRead(), nil
}
