# Computing Pi to 20 Decimal Places

`pi.py` prints pi to 20 decimal places (or any number you choose) using the
Chudnovsky algorithm with exact integer arithmetic. `test_pi.py` verifies the
output.

## 1. Usage

Requires Python 3.9+ and only the standard library. You do not need to install
or compile anything, and no virtual environment is needed.

```bash
python pi.py            # 3.14159265358979323846
python pi.py -n 50      # 50 decimal places
python pi.py --help
```

From other code:

```python
from pi import pi_digits
pi_digits(20)   # '3.14159265358979323846'
```

Run the tests:

```bash
python -m unittest -v test_pi
```

## 2. Method

### Chudnovsky formula

$$
\frac{1}{\pi} = 12 \sum_{k=0}^{\infty}
\frac{(-1)^k (6k)!\,(13591409 + 545140134k)}{(3k)!\,(k!)^3\,640320^{3k+3/2}}
$$

Each term adds about 14.18 correct digits, so 20 decimals takes only 3 terms.
Record-setting computations of pi (trillions of digits, e.g. y-cruncher) use
this formula.

### Why the result is exact

- **Binary splitting**: the series is summed as three big integers (P, Q, T)
  by recursively combining halves of the term range. No division happens
  during summation, so nothing is rounded.
- **Final step**: pi = 426880 · √10005 · Q / T. The square root is computed
  with `math.isqrt` on an integer scaled by 10^(2·precision), so it is also an
  exact integer operation.
- **Guard digits**: 10 extra digits are computed and then dropped. This keeps
  the floor in the integer division and square root from affecting the last
  printed digit.
- **Truncation**: the output is truncated, not rounded, so every printed digit
  is a true digit of pi. At 20 places the two are the same, because the 21st
  digit is 2.

Python floats (IEEE 754 doubles) carry only about 15–17 significant digits, so
`math.pi` cannot give 20 decimals. That is why the code uses integers.

## 3. Verification

`test_pi.py` checks the output in three independent ways:

| Test | What it checks |
|---|---|
| `test_twenty_decimals` | Exact match to `3.14159265358979323846` |
| `test_matches_reference_up_to_100` | Every length from 1 to 100 decimals matches the published digits of pi (OEIS [A000796](https://oeis.org/A000796)) |
| `test_matches_machin_formula` | At 20, 500 and 2000 decimals, matches pi computed separately with Machin's formula (π = 16·arctan(1/5) − 4·arctan(1/239)) in Python's `decimal` module |
| `test_zero_decimals`, `test_negative_rejected` | Edge cases |

Machin's formula is mathematically unrelated to Chudnovsky's and uses a
different arithmetic path (`decimal` instead of integers). If both agree to
2000 places, an error in either is very unlikely.

To check by hand, compare the output with any published table of pi digits,
for example OEIS A000796 or the NIST Digital Library of Mathematical Functions
(§3.12). You can also compare against `mpmath`, if it is installed:

```bash
python -c "import mpmath; mpmath.mp.dps = 30; print(mpmath.pi)"
```
