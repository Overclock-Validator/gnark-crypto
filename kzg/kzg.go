// Package kzg provides constructor for curved-typed KZG SRS
//
// For more details, see ecc/XXX/fr/kzg package
package kzg

import (
	"io"

	"github.com/Overclock-Validator/gnark-crypto/ecc"

	kzg_bls12377 "github.com/Overclock-Validator/gnark-crypto/ecc/bls12-377/kzg"
	kzg_bls12381 "github.com/Overclock-Validator/gnark-crypto/ecc/bls12-381/kzg"
	kzg_bls24315 "github.com/Overclock-Validator/gnark-crypto/ecc/bls24-315/kzg"
	kzg_bls24317 "github.com/Overclock-Validator/gnark-crypto/ecc/bls24-317/kzg"
	kzg_bn254 "github.com/Overclock-Validator/gnark-crypto/ecc/bn254/kzg"
	kzg_bw6633 "github.com/Overclock-Validator/gnark-crypto/ecc/bw6-633/kzg"
	kzg_bw6761 "github.com/Overclock-Validator/gnark-crypto/ecc/bw6-761/kzg"
)

type Serializable interface {
	io.ReaderFrom
	io.WriterTo
	BinaryDumper

	WriteRawTo(w io.Writer) (n int64, err error)
	UnsafeReadFrom(r io.Reader) (int64, error)
}

type BinaryDumper interface {
	WriteDump(w io.Writer, maxPkPoints ...int) error
	ReadDump(r io.Reader, maxPkPoints ...int) error
}

type SRS Serializable

// NewSRS returns an empty curved-typed SRS object
// that implements io.ReaderFrom and io.WriterTo interfaces
func NewSRS(curveID ecc.ID) SRS {
	switch curveID {
	case ecc.BN254:
		return &kzg_bn254.SRS{}
	case ecc.BLS12_377:
		return &kzg_bls12377.SRS{}
	case ecc.BLS12_381:
		return &kzg_bls12381.SRS{}
	case ecc.BLS24_315:
		return &kzg_bls24315.SRS{}
	case ecc.BLS24_317:
		return &kzg_bls24317.SRS{}
	case ecc.BW6_761:
		return &kzg_bw6761.SRS{}
	case ecc.BW6_633:
		return &kzg_bw6633.SRS{}
	default:
		panic("not implemented")
	}
}
