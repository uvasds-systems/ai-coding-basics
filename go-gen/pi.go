// Command pi prints pi to a given number of decimal places (default 20)
// using the Chudnovsky algorithm with exact big-integer arithmetic.
package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

const (
	guardDigits    = 10    // extra digits computed, then dropped
	digitsPerTerm  = 14.18 // correct digits added by each Chudnovsky term
	chudA          = 13591409
	chudB          = 545140134
	chudC          = 640320
	chudFinalScale = 426880 // C^(3/2)/12 = 426880 * sqrt(10005)
)

// c3Over24 is 640320^3 / 24, used in the Q term of each series step.
var c3Over24 = new(big.Int).Div(
	new(big.Int).Exp(big.NewInt(chudC), big.NewInt(3), nil),
	big.NewInt(24),
)

// binarySplit returns P(a,b), Q(a,b), T(a,b) for the Chudnovsky series terms
// in [a, b). Combining halves recursively keeps every value an exact integer.
func binarySplit(a, b int64) (p, q, t *big.Int) {
	if b-a == 1 {
		if a == 0 {
			p, q = big.NewInt(1), big.NewInt(1)
		} else {
			p = big.NewInt((6*a - 5) * (2*a - 1) * (6*a - 1))
			q = new(big.Int).Mul(big.NewInt(a*a), big.NewInt(a))
			q.Mul(q, c3Over24)
		}
		t = new(big.Int).Mul(p, big.NewInt(chudA+chudB*a))
		if a%2 == 1 {
			t.Neg(t)
		}
		return p, q, t
	}
	m := (a + b) / 2
	pam, qam, tam := binarySplit(a, m)
	pmb, qmb, tmb := binarySplit(m, b)
	p = new(big.Int).Mul(pam, pmb)
	q = new(big.Int).Mul(qam, qmb)
	t = new(big.Int).Add(new(big.Int).Mul(tam, qmb), new(big.Int).Mul(pam, tmb))
	return p, q, t
}

// PiDigits returns pi truncated (not rounded) to n decimal places,
// e.g. PiDigits(5) == "3.14159". n must be >= 0.
func PiDigits(n int) string {
	if n < 0 {
		panic("PiDigits: n must be non-negative")
	}
	prec := int64(n + guardDigits)
	terms := int64(float64(prec)/digitsPerTerm) + 2

	_, q, t := binarySplit(0, terms)

	// sqrt(10005) scaled by 10^prec, computed as an exact integer square root.
	one := new(big.Int).Exp(big.NewInt(10), big.NewInt(prec), nil)
	sqrtC := new(big.Int).Mul(big.NewInt(10005), new(big.Int).Mul(one, one))
	sqrtC.Sqrt(sqrtC)

	// pi * 10^prec = 426880 * sqrt(10005) * Q / T
	pi := new(big.Int).Mul(q, big.NewInt(chudFinalScale))
	pi.Mul(pi, sqrtC)
	pi.Quo(pi, t)

	// Drop guard digits (truncation), then insert the decimal point.
	pi.Quo(pi, new(big.Int).Exp(big.NewInt(10), big.NewInt(guardDigits), nil))
	s := pi.String()
	if n == 0 {
		return s
	}
	return s[:1] + "." + s[1:]
}

func main() {
	n := flag.Int("n", 20, "number of decimal places")
	flag.Parse()
	if *n < 0 {
		fmt.Fprintln(os.Stderr, "error: -n must be non-negative")
		os.Exit(2)
	}
	fmt.Println(PiDigits(*n))
}
