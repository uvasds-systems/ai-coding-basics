"""Verification tests for pi.py. Run with: python -m unittest -v test_pi"""

import unittest
from decimal import Decimal, getcontext

from pi import pi_digits

# First 100 decimal places of pi, from published references (e.g. OEIS A000796).
REFERENCE_PI = (
    "3.14159265358979323846264338327950288419716939937510"
    "58209749445923078164062862089986280348253421170679"
)


def machin_pi(decimals: int) -> str:
    """Independent check: Machin's formula pi = 16*atan(1/5) - 4*atan(1/239)."""
    getcontext().prec = decimals + 15

    def atan_inv(x: int) -> Decimal:
        x2, term, total, n, sign = x * x, Decimal(1) / x, Decimal(0), 1, 1
        eps = Decimal(10) ** -(decimals + 12)
        while term > eps:
            total += sign * term / n
            term /= x2
            n += 2
            sign = -sign
        return total

    pi = 16 * atan_inv(5) - 4 * atan_inv(239)
    whole, frac = str(pi).split(".")
    return whole + "." + frac[:decimals]


class TestPi(unittest.TestCase):
    def test_twenty_decimals(self):
        self.assertEqual(pi_digits(20), "3.14159265358979323846")

    def test_matches_reference_up_to_100(self):
        for n in range(1, 101):
            with self.subTest(decimals=n):
                self.assertEqual(pi_digits(n), REFERENCE_PI[: n + 2])

    def test_matches_machin_formula(self):
        for n in (20, 500, 2000):
            with self.subTest(decimals=n):
                self.assertEqual(pi_digits(n), machin_pi(n))

    def test_zero_decimals(self):
        self.assertEqual(pi_digits(0), "3")

    def test_negative_rejected(self):
        with self.assertRaises(ValueError):
            pi_digits(-1)


if __name__ == "__main__":
    unittest.main()
