"""Compute pi to a given number of decimal places using the Chudnovsky algorithm.

All arithmetic is exact integer arithmetic (binary splitting plus an integer
square root), so the result is not affected by floating-point rounding.
"""

import argparse
from math import isqrt

GUARD_DIGITS = 10  # extra digits computed, then discarded, to protect the last printed digit
DIGITS_PER_TERM = 14.181647462725477  # each Chudnovsky term adds ~14.18 correct digits

C = 640320
C3_OVER_24 = C**3 // 24


def _binary_split(a: int, b: int) -> tuple[int, int, int]:
    """Return (P, Q, T) for Chudnovsky terms in the half-open range [a, b)."""
    if b - a == 1:
        if a == 0:
            p = q = 1
        else:
            p = (6 * a - 5) * (2 * a - 1) * (6 * a - 1)
            q = a * a * a * C3_OVER_24
        t = p * (13591409 + 545140134 * a)
        return p, q, -t if a % 2 else t

    m = (a + b) // 2
    p1, q1, t1 = _binary_split(a, m)
    p2, q2, t2 = _binary_split(m, b)
    return p1 * p2, q1 * q2, q2 * t1 + p1 * t2


def pi_digits(decimals: int) -> str:
    """Return pi as a string truncated (not rounded) to `decimals` decimal places."""
    if decimals < 0:
        raise ValueError("decimals must be non-negative")

    precision = decimals + GUARD_DIGITS
    terms = int(precision / DIGITS_PER_TERM) + 2
    _, q, t = _binary_split(0, terms)

    # pi = (426880 * sqrt(10005) * Q) / T, scaled by 10**precision as an integer.
    scale = 10**precision
    sqrt_10005 = isqrt(10005 * scale * scale)
    pi_scaled = (426880 * sqrt_10005 * q) // t

    digits = str(pi_scaled // 10**GUARD_DIGITS)
    return digits[0] + "." + digits[1:] if decimals else digits[0]


def main() -> None:
    parser = argparse.ArgumentParser(description="Print pi to N decimal places.")
    parser.add_argument("-n", "--decimals", type=int, default=20,
                        help="number of decimal places (default: 20)")
    args = parser.parse_args()
    print(pi_digits(args.decimals))


if __name__ == "__main__":
    main()
