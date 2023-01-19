package config

var STARK_CURVE = Curve{
	Name:         "stark-curve",
	CurvePackage: "stark-curve",
	EnumID:       "STARK_CURVE",
	FrModulus:    "3618502788666131213697322783095070105526743751716087489154079457884512865583",
	FpModulus:    "3618502788666131213697322783095070105623107215331596699973092056135872020481",
	G1: Point{
		CoordType:        "fp.Element",
		CoordExtDegree:   1,
		PointName:        "g1",
		GLV:              false,
		CofactorCleaning: false,
		CRange:           defaultCRange(),
	},
	HashE1: &HashSuiteSvdw{
		z:  []string{"1"},
		c1: []string{"3141592653589793238462643383279502884197169399375105820974944592307816406667"},
		c2: []string{"1809251394333065606848661391547535052811553607665798349986546028067936010240"},
		c3: []string{"747120397548504753672821049844706693752799645928246271384591722031176001048"},
		c4: []string{"272520077186478842991245371323181269386250180546566216570369979330317493608"},
	},
}

func init() {
	addCurve(&STARK_CURVE)
}
