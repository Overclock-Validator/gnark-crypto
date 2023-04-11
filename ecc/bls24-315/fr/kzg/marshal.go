// Copyright 2020 ConsenSys Software Inc.
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
	"github.com/consensys/gnark-crypto/ecc/bls24-315"
	"io"
)

// WriteTo writes binary encoding of the ProvingKey
func (pk *ProvingKey) WriteTo(w io.Writer) (int64, error) {
	// encode the ProvingKey
	enc := bls24315.NewEncoder(w)
	if err := enc.Encode(pk.G1); err != nil {
		return enc.BytesWritten(), err
	}
	return enc.BytesWritten(), nil
}

// WriteRawTo writes binary encoding of Proof to w without point compression
func (vk *VerifyingKey) WriteRawTo(w io.Writer) (int64, error) {
	return vk.writeTo(w, bls24315.RawEncoding())
}

// WriteTo writes binary encoding of the VerifyingKey
func (vk *VerifyingKey) WriteTo(w io.Writer) (int64, error) {
	return vk.writeTo(w)
}

func (vk *VerifyingKey) writeTo(w io.Writer, options ...func(*bls24315.Encoder)) (int64, error) {
	// encode the VerifyingKey
	enc := bls24315.NewEncoder(w, options...)

	toEncode := []interface{}{
		&vk.G2[0],
		&vk.G2[1],
		&vk.G1,
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
	// encode the VerifyingKey
	enc := bls24315.NewEncoder(w)

	toEncode := []interface{}{
		&srs.Vk.G2[0],
		&srs.Vk.G2[1],
		srs.Pk.G1,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return enc.BytesWritten(), err
		}
	}

	return enc.BytesWritten(), nil
}

// ReadFrom decodes ProvingKey data from reader.
func (pk *ProvingKey) ReadFrom(r io.Reader) (int64, error) {
	// decode the ProvingKey
	dec := bls24315.NewDecoder(r)

	toDecode := []interface{}{
		&pk.G1,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	return dec.BytesRead(), nil
}

// ReadFrom decodes VerifyingKey data from reader.
func (vk *VerifyingKey) ReadFrom(r io.Reader) (int64, error) {
	// decode the VerifyingKey
	dec := bls24315.NewDecoder(r)

	toDecode := []interface{}{
		&vk.G2[0],
		&vk.G2[1],
		&vk.G1,
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
	dec := bls24315.NewDecoder(r)

	toDecode := []interface{}{
		&srs.Vk.G2[0],
		&srs.Vk.G2[1],
		&srs.Pk.G1,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}

	srs.Vk.G1 = srs.Pk.G1[0]

	return dec.BytesRead(), nil
}

// WriteTo writes binary encoding of a OpeningProof
func (proof *OpeningProof) WriteTo(w io.Writer) (int64, error) {
	enc := bls24315.NewEncoder(w)

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
	dec := bls24315.NewDecoder(r)

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
	enc := bls24315.NewEncoder(w)

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
	dec := bls24315.NewDecoder(r)
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
