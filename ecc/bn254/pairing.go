// Copyright 2020 ConsenSys AG
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

package bn254

import (
	"errors"

	"github.com/Overclock-Validator/gnark-crypto/ecc/bn254/fp"
	"github.com/Overclock-Validator/gnark-crypto/ecc/bn254/fptower"
)

// GT target group of the pairing
type GT = fptower.E12

type lineEvaluation struct {
	r0 fptower.E2
	r1 fptower.E2
	r2 fptower.E2
}

// Pair calculates the reduced pairing for a set of points
// ∏ᵢ e(Pᵢ, Qᵢ).
//
// This function doesn't check that the inputs are in the correct subgroup. See IsInSubGroup.
func Pair(P []G1Affine, Q []G2Affine) (GT, error) {
	f, err := MillerLoop(P, Q)
	if err != nil {
		return GT{}, err
	}
	return FinalExponentiation(&f), nil
}

// PairingCheck calculates the reduced pairing for a set of points and returns True if the result is One
// ∏ᵢ e(Pᵢ, Qᵢ) =? 1
//
// This function doesn't check that the inputs are in the correct subgroup. See IsInSubGroup.
func PairingCheck(P []G1Affine, Q []G2Affine) (bool, error) {
	f, err := Pair(P, Q)
	if err != nil {
		return false, err
	}
	var one GT
	one.SetOne()
	return f.Equal(&one), nil
}

// FinalExponentiation computes the exponentiation (∏ᵢ zᵢ)ᵈ
// where d = (p¹²-1)/r = (p¹²-1)/Φ₁₂(p) ⋅ Φ₁₂(p)/r = (p⁶-1)(p²+1)(p⁴ - p² +1)/r
// we use instead d=s ⋅ (p⁶-1)(p²+1)(p⁴ - p² +1)/r
// where s is the cofactor 2x₀(6x₀²+3x₀+1)
func FinalExponentiation(z *GT, _z ...*GT) GT {

	var result GT
	result.Set(z)

	for _, e := range _z {
		result.Mul(&result, e)
	}

	var t [5]GT

	// Easy part
	// (p⁶-1)(p²+1)
	t[0].Conjugate(&result)
	result.Inverse(&result)
	t[0].Mul(&t[0], &result)
	result.FrobeniusSquare(&t[0]).Mul(&result, &t[0])

	var one GT
	one.SetOne()
	if result.Equal(&one) {
		return result
	}

	// Hard part (up to permutation)
	// 2x₀(6x₀²+3x₀+1)(p⁴-p²+1)/r
	// Duquesne and Ghammam
	// https://eprint.iacr.org/2015/192.pdf
	// Fuentes et al. (alg. 6)
	t[0].Expt(&result).
		Conjugate(&t[0])
	t[0].CyclotomicSquare(&t[0])
	t[1].CyclotomicSquare(&t[0])
	t[1].Mul(&t[0], &t[1])
	t[2].Expt(&t[1])
	t[2].Conjugate(&t[2])
	t[3].Conjugate(&t[1])
	t[1].Mul(&t[2], &t[3])
	t[3].CyclotomicSquare(&t[2])
	t[4].Expt(&t[3])
	t[4].Mul(&t[1], &t[4])
	t[3].Mul(&t[0], &t[4])
	t[0].Mul(&t[2], &t[4])
	t[0].Mul(&result, &t[0])
	t[2].Frobenius(&t[3])
	t[0].Mul(&t[2], &t[0])
	t[2].FrobeniusSquare(&t[4])
	t[0].Mul(&t[2], &t[0])
	t[2].Conjugate(&result)
	t[2].Mul(&t[2], &t[3])
	t[2].FrobeniusCube(&t[2])
	t[0].Mul(&t[2], &t[0])

	return t[0]
}

// MillerLoop computes the multi-Miller loop
// ∏ᵢ MillerLoop(Pᵢ, Qᵢ) =
// ∏ᵢ { fᵢ_{6x₀+2,Qᵢ}(Pᵢ) · ℓᵢ_{[6x₀+2]Qᵢ,π(Qᵢ)}(Pᵢ) · ℓᵢ_{[6x₀+2]Qᵢ+π(Qᵢ),-π²(Qᵢ)}(Pᵢ) }
func MillerLoop(P []G1Affine, Q []G2Affine) (GT, error) {
	n := len(P)
	if n == 0 || n != len(Q) {
		return GT{}, errors.New("invalid inputs sizes")
	}

	// filter infinity points
	p := make([]G1Affine, 0, n)
	q := make([]G2Affine, 0, n)

	for k := 0; k < n; k++ {
		if P[k].IsInfinity() || Q[k].IsInfinity() {
			continue
		}
		p = append(p, P[k])
		q = append(q, Q[k])
	}
	n = len(p)

	// projective points for Q
	qProj := make([]g2Proj, n)
	qNeg := make([]G2Affine, n)
	for k := 0; k < n; k++ {
		qProj[k].FromAffine(&q[k])
		qNeg[k].Neg(&q[k])
	}

	var result GT
	result.SetOne()
	var l2, l1 lineEvaluation
	var prodLines [5]E2

	// Compute ∏ᵢ { fᵢ_{6x₀+2,Q}(P) }
	if n >= 1 {
		// i = 64, separately to avoid an E12 Square
		// (Square(res) = 1² = 1)
		// LoopCounter[64] = 0
		// k = 0, separately to avoid MulBy034 (res × ℓ)
		// (assign line to res)

		// qProj[0] ← 2qProj[0] and l1 the tangent ℓ passing 2qProj[0]
		qProj[0].doubleStep(&l1)
		// line evaluation at P[0] (assign)
		result.C0.B0.MulByElement(&l1.r0, &p[0].Y)
		result.C1.B0.MulByElement(&l1.r1, &p[0].X)
		result.C1.B1.Set(&l1.r2)
	}

	if n >= 2 {
		// k = 1, separately to avoid MulBy034 (res × ℓ)
		// (res is also a line at this point, so we use Mul034By034 ℓ × ℓ)

		// qProj[1] ← 2qProj[1] and l1 the tangent ℓ passing 2qProj[1]
		qProj[1].doubleStep(&l1)
		// line evaluation at P[1]
		l1.r0.MulByElement(&l1.r0, &p[1].Y)
		l1.r1.MulByElement(&l1.r1, &p[1].X)
		// ℓ × res
		prodLines = fptower.Mul034By034(&l1.r0, &l1.r1, &l1.r2, &result.C0.B0, &result.C1.B0, &result.C1.B1)
		result.C0.B0 = prodLines[0]
		result.C0.B1 = prodLines[1]
		result.C0.B2 = prodLines[2]
		result.C1.B0 = prodLines[3]
		result.C1.B1 = prodLines[4]
	}

	// k >= 2
	for k := 2; k < n; k++ {
		// qProj[k] ← 2qProj[k] and l1 the tangent ℓ passing 2qProj[k]
		qProj[k].doubleStep(&l1)
		// line evaluation at P[k]
		l1.r0.MulByElement(&l1.r0, &p[k].Y)
		l1.r1.MulByElement(&l1.r1, &p[k].X)
		// ℓ × res
		result.MulBy034(&l1.r0, &l1.r1, &l1.r2)
	}

	// i = 63, separately to avoid a doubleStep (LoopCounter[63]=-1)
	// (at this point qProj = 2Q, so 2qProj-Q=3Q is equivalent to qProj+Q=3Q
	// this means doubleStep followed by an addMixedStep is equivalent to an
	// addMixedStep here)

	result.Square(&result)
	for k := 0; k < n; k++ {
		// l2 the line passing qProj[k] and -Q
		// (avoids a point addition: qProj[k]-Q)
		qProj[k].lineCompute(&l2, &qNeg[k])
		// line evaluation at P[k]
		l2.r0.MulByElement(&l2.r0, &p[k].Y)
		l2.r1.MulByElement(&l2.r1, &p[k].X)
		// qProj[k] ← qProj[k]+Q[k] and
		// l1 the line ℓ passing qProj[k] and Q[k]
		qProj[k].addMixedStep(&l1, &q[k])
		// line evaluation at P[k]
		l1.r0.MulByElement(&l1.r0, &p[k].Y)
		l1.r1.MulByElement(&l1.r1, &p[k].X)
		// ℓ × ℓ
		prodLines = fptower.Mul034By034(&l1.r0, &l1.r1, &l1.r2, &l2.r0, &l2.r1, &l2.r2)
		// (ℓ × ℓ) × res
		result.MulBy01234(&prodLines)
	}

	// i <= 62
	for i := len(LoopCounter) - 4; i >= 0; i-- {
		// mutualize the square among n Miller loops
		// (∏ᵢfᵢ)²
		result.Square(&result)

		for k := 0; k < n; k++ {
			// qProj[k] ← 2qProj[k] and l1 the tangent ℓ passing 2qProj[k]
			qProj[k].doubleStep(&l1)
			// line evaluation at P[k]
			l1.r0.MulByElement(&l1.r0, &p[k].Y)
			l1.r1.MulByElement(&l1.r1, &p[k].X)

			if LoopCounter[i] == 1 {
				// qProj[k] ← qProj[k]+Q[k] and
				// l2 the line ℓ passing qProj[k] and Q[k]
				qProj[k].addMixedStep(&l2, &q[k])
				// line evaluation at P[k]
				l2.r0.MulByElement(&l2.r0, &p[k].Y)
				l2.r1.MulByElement(&l2.r1, &p[k].X)
				// ℓ × ℓ
				prodLines = fptower.Mul034By034(&l1.r0, &l1.r1, &l1.r2, &l2.r0, &l2.r1, &l2.r2)
				// (ℓ × ℓ) × res
				result.MulBy01234(&prodLines)

			} else if LoopCounter[i] == -1 {
				// qProj[k] ← qProj[k]-Q[k] and
				// l2 the line ℓ passing qProj[k] and -Q[k]
				qProj[k].addMixedStep(&l2, &qNeg[k])
				// line evaluation at P[k]
				l2.r0.MulByElement(&l2.r0, &p[k].Y)
				l2.r1.MulByElement(&l2.r1, &p[k].X)
				// ℓ × ℓ
				prodLines = fptower.Mul034By034(&l1.r0, &l1.r1, &l1.r2, &l2.r0, &l2.r1, &l2.r2)
				// (ℓ × ℓ) × res
				result.MulBy01234(&prodLines)
			} else {
				// ℓ × res
				result.MulBy034(&l1.r0, &l1.r1, &l1.r2)
			}
		}
	}

	// Compute  ∏ᵢ { ℓᵢ_{[6x₀+2]Q,π(Q)}(P) · ℓᵢ_{[6x₀+2]Q+π(Q),-π²(Q)}(P) }
	var Q1, Q2 G2Affine
	for k := 0; k < n; k++ {
		//Q1 = π(Q)
		Q1.X.Conjugate(&q[k].X).MulByNonResidue1Power2(&Q1.X)
		Q1.Y.Conjugate(&q[k].Y).MulByNonResidue1Power3(&Q1.Y)

		// Q2 = -π²(Q)
		Q2.X.MulByNonResidue2Power2(&q[k].X)
		Q2.Y.MulByNonResidue2Power3(&q[k].Y).Neg(&Q2.Y)

		// qProj[k] ← qProj[k]+π(Q) and
		// l1 the line passing qProj[k] and π(Q)
		qProj[k].addMixedStep(&l2, &Q1)
		// line evaluation at P[k]
		l2.r0.MulByElement(&l2.r0, &p[k].Y)
		l2.r1.MulByElement(&l2.r1, &p[k].X)

		// l2 the line passing qProj[k] and -π²(Q)
		// (avoids a point addition: qProj[k]-π²(Q))
		qProj[k].lineCompute(&l1, &Q2)
		// line evaluation at P[k]
		l1.r0.MulByElement(&l1.r0, &p[k].Y)
		l1.r1.MulByElement(&l1.r1, &p[k].X)

		// ℓ × ℓ
		prodLines = fptower.Mul034By034(&l1.r0, &l1.r1, &l1.r2, &l2.r0, &l2.r1, &l2.r2)
		// (ℓ × ℓ) × res
		result.MulBy01234(&prodLines)
	}

	return result, nil
}

// doubleStep doubles a point in Homogenous projective coordinates, and evaluates the line in Miller loop
// https://eprint.iacr.org/2013/722.pdf (Section 4.3)
func (p *g2Proj) doubleStep(evaluations *lineEvaluation) {

	// get some Element from our pool
	var t1, A, B, C, D, E, EE, F, G, H, I, J, K fptower.E2
	A.Mul(&p.x, &p.y)
	A.Halve()
	B.Square(&p.y)
	C.Square(&p.z)
	D.Double(&C).
		Add(&D, &C)
	E.MulBybTwistCurveCoeff(&D)
	F.Double(&E).
		Add(&F, &E)
	G.Add(&B, &F)
	G.Halve()
	H.Add(&p.y, &p.z).
		Square(&H)
	t1.Add(&B, &C)
	H.Sub(&H, &t1)
	I.Sub(&E, &B)
	J.Square(&p.x)
	EE.Square(&E)
	K.Double(&EE).
		Add(&K, &EE)

	// X, Y, Z
	p.x.Sub(&B, &F).
		Mul(&p.x, &A)
	p.y.Square(&G).
		Sub(&p.y, &K)
	p.z.Mul(&B, &H)

	// Line evaluation
	evaluations.r0.Neg(&H)
	evaluations.r1.Double(&J).
		Add(&evaluations.r1, &J)
	evaluations.r2.Set(&I)
}

// addMixedStep point addition in Mixed Homogenous projective and Affine coordinates
// https://eprint.iacr.org/2013/722.pdf (Section 4.3)
func (p *g2Proj) addMixedStep(evaluations *lineEvaluation, a *G2Affine) {

	// get some Element from our pool
	var Y2Z1, X2Z1, O, L, C, D, E, F, G, H, t0, t1, t2, J fptower.E2
	Y2Z1.Mul(&a.Y, &p.z)
	O.Sub(&p.y, &Y2Z1)
	X2Z1.Mul(&a.X, &p.z)
	L.Sub(&p.x, &X2Z1)
	C.Square(&O)
	D.Square(&L)
	E.Mul(&L, &D)
	F.Mul(&p.z, &C)
	G.Mul(&p.x, &D)
	t0.Double(&G)
	H.Add(&E, &F).
		Sub(&H, &t0)
	t1.Mul(&p.y, &E)

	// X, Y, Z
	p.x.Mul(&L, &H)
	p.y.Sub(&G, &H).
		Mul(&p.y, &O).
		Sub(&p.y, &t1)
	p.z.Mul(&E, &p.z)

	t2.Mul(&L, &a.Y)
	J.Mul(&a.X, &O).
		Sub(&J, &t2)

	// Line evaluation
	evaluations.r0.Set(&L)
	evaluations.r1.Neg(&O)
	evaluations.r2.Set(&J)
}

// lineCompute computes the line through p in Homogenous projective coordinates
// and a in affine coordinates. It does not compute the resulting point p+a.
func (p *g2Proj) lineCompute(evaluations *lineEvaluation, a *G2Affine) {

	// get some Element from our pool
	var Y2Z1, X2Z1, O, L, t2, J fptower.E2
	Y2Z1.Mul(&a.Y, &p.z)
	O.Sub(&p.y, &Y2Z1)
	X2Z1.Mul(&a.X, &p.z)
	L.Sub(&p.x, &X2Z1)
	t2.Mul(&L, &a.Y)
	J.Mul(&a.X, &O).
		Sub(&J, &t2)

	// Line evaluation
	evaluations.r0.Set(&L)
	evaluations.r1.Neg(&O)
	evaluations.r2.Set(&J)
}

// ----------------------
// Fixed-argument pairing
// ----------------------

type LineEvaluationAff struct {
	R0 fptower.E2
	R1 fptower.E2
}

// PairFixedQ calculates the reduced pairing for a set of points
// ∏ᵢ e(Pᵢ, Qᵢ) where Q are fixed points in G2.
//
// This function doesn't check that the inputs are in the correct subgroup. See IsInSubGroup.
func PairFixedQ(P []G1Affine, lines [][2][len(LoopCounter)]LineEvaluationAff) (GT, error) {
	f, err := MillerLoopFixedQ(P, lines)
	if err != nil {
		return GT{}, err
	}
	return FinalExponentiation(&f), nil
}

// PairingCheckFixedQ calculates the reduced pairing for a set of points and returns True if the result is One
// ∏ᵢ e(Pᵢ, Qᵢ) =? 1 where Q are fixed points in G2.
//
// This function doesn't check that the inputs are in the correct subgroup. See IsInSubGroup.
func PairingCheckFixedQ(P []G1Affine, lines [][2][len(LoopCounter)]LineEvaluationAff) (bool, error) {
	f, err := PairFixedQ(P, lines)
	if err != nil {
		return false, err
	}
	var one GT
	one.SetOne()
	return f.Equal(&one), nil
}

// PrecomputeLines precomputes the lines for the fixed-argument Miller loop
func PrecomputeLines(Q G2Affine) (PrecomputedLines [2][len(LoopCounter)]LineEvaluationAff) {
	var accQ, negQ G2Affine
	accQ.Set(&Q)
	negQ.Neg(&Q)

	n := len(LoopCounter)
	for i := n - 2; i >= 0; i-- {
		switch LoopCounter[i] {
		case 0:
			accQ.doubleStep(&PrecomputedLines[0][i])
		case 1:
			accQ.doubleAndAddStep(&PrecomputedLines[0][i], &PrecomputedLines[1][i], &Q)
		case -1:
			accQ.doubleAndAddStep(&PrecomputedLines[0][i], &PrecomputedLines[1][i], &negQ)
		default:
			return [2][len(LoopCounter)]LineEvaluationAff{}
		}
	}

	var psiQ, phiQ G2Affine
	phiQ.X.Conjugate(&Q.X).MulByNonResidue1Power2(&phiQ.X)
	phiQ.Y.Conjugate(&Q.Y).MulByNonResidue1Power3(&phiQ.Y)
	psiQ.X.MulByNonResidue2Power2(&Q.X)
	psiQ.Y.MulByNonResidue2Power3(&Q.Y).Neg(&psiQ.Y)

	accQ.addStep(&PrecomputedLines[1][n-1], &phiQ)
	accQ.addStep(&PrecomputedLines[0][n-1], &psiQ)

	return PrecomputedLines
}

// MillerLoopFixedQ computes the multi-Miller loop as in MillerLoop
// but Qᵢ are fixed points in G2 known in advance.
func MillerLoopFixedQ(P []G1Affine, lines [][2][len(LoopCounter)]LineEvaluationAff) (GT, error) {
	n := len(P)
	if n == 0 || n != len(lines) {
		return GT{}, errors.New("invalid inputs sizes")
	}

	// no need to filter infinity points:
	// 		1. if Pᵢ=(0,0) then -x/y=1/y=0 by gnark-crypto convention and so
	// 		lines R0 and R1 are 0. At the end it happens that result will stay
	// 		1 through the Miller loop because MulBy34(1,0,0)==1
	// 		Mul34By34(1,0,0,1,0,0)==1 and MulBy01234(1,0,0,0,0)==1.
	//
	// 		2. if Qᵢ=(0,0) then PrecomputeLines(Qᵢ) will return lines R0 and R1
	// 		that are 0 because of gnark-convention (*/0==0) in doubleStep and
	// 		addStep. Similarly to Pᵢ=(0,0) it happens that result stays 1
	// 		throughout the MillerLoop.

	// precomputations
	yInv := make([]fp.Element, n)
	xNegOverY := make([]fp.Element, n)
	for k := 0; k < n; k++ {
		yInv[k].Set(&P[k].Y)
	}
	yInv = fp.BatchInvert(yInv)
	for k := 0; k < n; k++ {
		xNegOverY[k].Mul(&P[k].X, &yInv[k]).
			Neg(&xNegOverY[k])
	}

	var result GT
	result.SetOne()
	var prodLines [5]E2

	// Compute ∏ᵢ { fᵢ_{6x₀+2,Q}(P) }
	if n >= 1 {
		// i = 64, separately to avoid an E12 Square
		// (Square(res) = 1² = 1)
		// LoopCounter[64] = 0
		// k = 0, separately to avoid MulBy34 (res × ℓ)
		// (assign line to res)

		// line evaluation at P[0] (assign)
		result.C1.B0.MulByElement(&lines[0][0][64].R0, &xNegOverY[0])
		result.C1.B1.MulByElement(&lines[0][0][64].R1, &yInv[0])
		// the coefficient which MulBy34 sets to 1 happens to be already 1 (result = 1)
	}

	if n >= 2 {
		// k = 1, separately to avoid MulBy34 (res × ℓ)
		// (res is also a line at this point, so we use Mul34By34 ℓ × ℓ)
		// line evaluation at P[1]
		lines[1][0][64].R0.MulByElement(&lines[1][0][64].R0, &xNegOverY[1])
		lines[1][0][64].R1.MulByElement(&lines[1][0][64].R1, &yInv[1])
		// ℓ × res
		prodLines = fptower.Mul34By34(&lines[1][0][64].R0, &lines[1][0][64].R1, &result.C1.B0, &result.C1.B1)
		result.C0.B0 = prodLines[0]
		result.C0.B1 = prodLines[1]
		result.C0.B2 = prodLines[2]
		result.C1.B0 = prodLines[3]
		result.C1.B1 = prodLines[4]
	}

	// k >= 2
	for k := 2; k < n; k++ {
		// line evaluation at P[k]
		lines[k][0][64].R0.MulByElement(&lines[k][0][64].R0, &xNegOverY[k])
		lines[k][0][64].R1.MulByElement(&lines[k][0][64].R1, &yInv[k])
		// ℓ × res
		result.MulBy34(
			&lines[k][0][64].R0,
			&lines[k][0][64].R1,
		)
	}

	for i := len(LoopCounter) - 3; i >= 0; i-- {
		// mutualize the square among n Miller loops
		// (∏ᵢfᵢ)²
		result.Square(&result)

		for k := 0; k < n; k++ {
			// line evaluation at P[k]
			lines[k][0][i].R0.
				MulByElement(
					&lines[k][0][i].R0,
					&xNegOverY[k],
				)
			lines[k][0][i].R1.
				MulByElement(
					&lines[k][0][i].R1,
					&yInv[k],
				)

			if LoopCounter[i] == 0 {
				// ℓ × res
				result.MulBy34(
					&lines[k][0][i].R0,
					&lines[k][0][i].R1,
				)
			} else {
				// line evaluation at P[k]
				lines[k][1][i].R0.
					MulByElement(
						&lines[k][1][i].R0,
						&xNegOverY[k],
					)
				lines[k][1][i].R1.
					MulByElement(
						&lines[k][1][i].R1,
						&yInv[k],
					)
				// ℓ × ℓ
				prodLines = fptower.Mul34By34(
					&lines[k][0][i].R0, &lines[k][0][i].R1,
					&lines[k][1][i].R0, &lines[k][1][i].R1,
				)
				// (ℓ × ℓ) × res
				result.MulBy01234(&prodLines)
			}
		}
	}

	// Compute  ∏ᵢ { ℓᵢ_{[6x₀+2]Q,π(Q)}(P) · ℓᵢ_{[6x₀+2]Q+π(Q),-π²(Q)}(P) }
	for k := 0; k < n; k++ {
		// line evaluation at P[k]
		lines[k][1][65].R0.
			MulByElement(
				&lines[k][1][65].R0,
				&xNegOverY[k],
			)
		lines[k][1][65].R1.
			MulByElement(
				&lines[k][1][65].R1,
				&yInv[k],
			)
		// line evaluation at P[k]
		lines[k][0][65].R0.
			MulByElement(
				&lines[k][0][65].R0,
				&xNegOverY[k],
			)
		lines[k][0][65].R1.
			MulByElement(
				&lines[k][0][65].R1,
				&yInv[k],
			)
		// ℓ × ℓ
		prodLines = fptower.Mul34By34(
			&lines[k][1][65].R0, &lines[k][1][65].R1,
			&lines[k][0][65].R0, &lines[k][0][65].R1,
		)
		// (ℓ × ℓ) × res
		result.MulBy01234(&prodLines)
	}

	return result, nil
}

func (p *G2Affine) doubleStep(evaluations *LineEvaluationAff) {

	var n, d, λ, xr, yr fptower.E2
	// λ = 3x²/2y
	n.Square(&p.X)
	λ.Double(&n).
		Add(&λ, &n)
	d.Double(&p.Y)
	λ.Div(&λ, &d)

	// xr = λ²-2x
	xr.Square(&λ).
		Sub(&xr, &p.X).
		Sub(&xr, &p.X)

	// yr = λ(x-xr)-y
	yr.Sub(&p.X, &xr).
		Mul(&yr, &λ).
		Sub(&yr, &p.Y)

	evaluations.R0.Set(&λ)
	evaluations.R1.Mul(&λ, &p.X).
		Sub(&evaluations.R1, &p.Y)

	p.X.Set(&xr)
	p.Y.Set(&yr)
}

func (p *G2Affine) addStep(evaluations *LineEvaluationAff, a *G2Affine) {
	var n, d, λ, λλ, xr, yr fptower.E2

	// compute λ = (y2-y1)/(x2-x1)
	n.Sub(&a.Y, &p.Y)
	d.Sub(&a.X, &p.X)
	λ.Div(&n, &d)

	// xr = λ²-x1-x2
	λλ.Square(&λ)
	n.Add(&p.X, &a.X)
	xr.Sub(&λλ, &n)

	// yr = λ(x1-xr) - y1
	yr.Sub(&p.X, &xr).
		Mul(&yr, &λ).
		Sub(&yr, &p.Y)

	evaluations.R0.Set(&λ)
	evaluations.R1.Mul(&λ, &p.X).
		Sub(&evaluations.R1, &p.Y)

	p.X.Set(&xr)
	p.Y.Set(&yr)
}

func (p *G2Affine) doubleAndAddStep(evaluations1, evaluations2 *LineEvaluationAff, a *G2Affine) {
	var n, d, l1, x3, l2, x4, y4 fptower.E2

	// compute λ1 = (y2-y1)/(x2-x1)
	n.Sub(&p.Y, &a.Y)
	d.Sub(&p.X, &a.X)
	l1.Div(&n, &d)

	// compute x3 =λ1²-x1-x2
	x3.Square(&l1)
	x3.Sub(&x3, &p.X)
	x3.Sub(&x3, &a.X)

	// omit y3 computation

	// compute line1
	evaluations1.R0.Set(&l1)
	evaluations1.R1.Mul(&l1, &p.X)
	evaluations1.R1.Sub(&evaluations1.R1, &p.Y)

	// compute λ2 = -λ1-2y1/(x3-x1)
	n.Double(&p.Y)
	d.Sub(&x3, &p.X)
	l2.Div(&n, &d)
	l2.Add(&l2, &l1)
	l2.Neg(&l2)

	// compute x4 = λ2²-x1-x3
	x4.Square(&l2)
	x4.Sub(&x4, &p.X)
	x4.Sub(&x4, &x3)

	// compute y4 = λ2(x1 - x4)-y1
	y4.Sub(&p.X, &x4)
	y4.Mul(&l2, &y4)
	y4.Sub(&y4, &p.Y)

	// compute line2
	evaluations2.R0.Set(&l2)
	evaluations2.R1.Mul(&l2, &p.X)
	evaluations2.R1.Sub(&evaluations2.R1, &p.Y)

	p.X.Set(&x4)
	p.Y.Set(&y4)
}
