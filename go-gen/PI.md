# Computing Pi to 20 Decimal Places (Go)

`pi.go` prints pi to 20 decimal places (or any number you choose) using the
Chudnovsky algorithm with exact big-integer arithmetic. `pi_test.go` verifies
the output.

## 1. Usage

Requires Go 1.21+ and only the standard library (`math/big`). No third-party
packages are needed.

Run directly:

```bash
go run .            # 3.14159265358979323846
go run . -n 50      # 50 decimal places
go run . -h         # help
```

Compile to a binary:

```bash
go build -o pi .
./pi                # 3.14159265358979323846
./pi -n 1000
```

From other Go code in the package:

```go
PiDigits(20)   // "3.14159265358979323846"
```

Run the tests:

```bash
go test -v
```

## 2. Method

### Chudnovsky formula

$$
\frac{1}{\pi} = 12 \sum_{k=0}^{\infty}
\frac{(-1)^k (6k)!\,(13591409 + 545140134k)}{(3k)!\,(k!)^3\,640320^{3k+3/2}}
$$

Each term adds about 14.18 correct digits, so 20 decimals (plus 10 guard
digits) takes only 4 terms. Record-setting computations of pi (trillions of
digits, e.g. y-cruncher) use this formula.

### Why the result is exact

- **Binary splitting**: the series is summed as three `big.Int` values (P, Q,
  T) by recursively combining halves of the term range. No division happens
  during summation, so nothing is rounded.
- **Final step**: pi = 426880 · √10005 · Q / T. The square root is computed
  with `big.Int.Sqrt` on an integer scaled by 10^(2·precision), so it is also
  an exact integer operation.
- **Guard digits**: 10 extra digits are computed and then dropped. This keeps
  the floor in the integer division and square root from affecting the last
  printed digit.
- **Truncation**: the output is truncated, not rounded, so every printed digit
  is a true digit of pi. At 20 places the two are the same, because the 21st
  digit is 2.

Go's `float64` (IEEE 754 double) carries only about 15–17 significant digits,
so `math.Pi` cannot give 20 decimals. `big.Float` could, but its binary
mantissa still rounds; integer arithmetic avoids that entirely.

## 3. Verification

`pi_test.go` checks the result in several independent ways:

| Test | What it checks |
|------|----------------|
| `TestTwentyPlaces` | `PiDigits(20)` equals `3.14159265358979323846` |
| `TestAgainstReference` | Every `n` from 0 to 100 matches the published first 100 decimals (OEIS [A000796](https://oeis.org/A000796)) |
| `TestPrefixConsistency` | `PiDigits(n)` for n = 1..500 is a prefix of `PiDigits(2000)`, catching errors in trailing digits |
| `TestMatchesMathPi` | First 15 decimals agree with Go's `math.Pi` |
| `TestLength` | Output length is exactly `n + 2` characters |
| `TestNegativePanics` | Negative `n` is rejected |

You can also check the output against external sources:

- OEIS A000796: <https://oeis.org/A000796>
- The `bc` calculator, which uses arbitrary precision:
  `echo "scale=25; 4*a(1)" | bc -l` → `3.1415926535897932384626432` (the
  last digit differs from pi because `bc` truncates intermediate results,
  but the first 20 decimals match)
