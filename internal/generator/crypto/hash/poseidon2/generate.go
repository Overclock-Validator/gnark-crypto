package poseidon2

import (
	"path/filepath"

	"github.com/Overclock-Validator/gnark-crypto/internal/generator/config"
	"github.com/consensys/bavard"
)

func Generate(conf config.Curve, baseDir string, bgen *bavard.BatchGenerator) error {

	conf.Package = "poseidon2"
	entries := []bavard.Entry{
		{File: filepath.Join(baseDir, "poseidon2.go"), Templates: []string{"poseidon2.go.tmpl"}},
		{File: filepath.Join(baseDir, "poseidon2_test.go"), Templates: []string{"poseidon2.test.go.tmpl"}},
	}

	return bgen.Generate(conf, conf.Package, "./crypto/hash/poseidon2/template", entries...)

}
