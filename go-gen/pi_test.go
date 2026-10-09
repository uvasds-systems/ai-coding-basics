package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// First 100 decimals of pi (OEIS A000796).
const pi100 = "3.1415926535897932384626433832795028841971693993751" +
	"058209749445923078164062862089986280348253421170679"

func TestTwentyPlaces(t *testing.T) {
	if got, want := PiDigits(20), "3.14159265358979323846"; got != want {
		t.Fatalf("PiDigits(20) = %s, want %s", got, want)
	}
}

func TestAgainstReference(t *testing.T) {
	for n := 0; n <= 100; n++ {
		want := pi100[:2+n]
		if n == 0 {
			want = "3"
		}
		if got := PiDigits(n); got != want {
			t.Errorf("PiDigits(%d) = %s, want %s", n, got, want)
		}
	}
}

// Every shorter result must be a prefix of a much longer one. This catches
// errors in the last digits (insufficient terms or guard digits).
func TestPrefixConsistency(t *testing.T) {
	long := PiDigits(2000)
	for n := 1; n <= 500; n++ {
		if got := PiDigits(n); !strings.HasPrefix(long, got) {
			t.Errorf("PiDigits(%d) = %s is not a prefix of PiDigits(2000)", n, got)
		}
	}
}

// Independent check: agrees with float64 math.Pi to 15 decimals.
func TestMatchesMathPi(t *testing.T) {
	want := fmt.Sprintf("%.16f", math.Pi)[:17] // truncate to 15 decimals
	if got := PiDigits(15); got != want {
		t.Fatalf("PiDigits(15) = %s, want %s", got, want)
	}
}

func TestLength(t *testing.T) {
	for _, n := range []int{1, 20, 1000} {
		if got := len(PiDigits(n)); got != n+2 {
			t.Errorf("len(PiDigits(%d)) = %d, want %d", n, got, n+2)
		}
	}
}

func TestNegativePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("PiDigits(-1) did not panic")
		}
	}()
	PiDigits(-1)
}
