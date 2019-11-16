package caterinian

import (
	"math"
)

// IsCaterinian returns if a number is
// a caterinian number. A caterinian number
// is a natural number, represented as
// a_n*10^n + a_{n-1}*10^(n-1) + ... + a_1*10 + a_0,
// for which the sum of the sum of |a_i - a_{n-k}| as k
// goes from 0 to n and as i goes from 0 to n minus (n+1)
// is equal to a_n.
func IsCaterinian(n int) bool {
	var a []int8

	if n <= 0 {
		return false
	}

	for n != 0 {
		r := n % 10
		a = append(a, int8(r))
		n = n / 10
	}

	if len(a) == 0 {
		return false
	}

	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}

	z := 0
	for i := 0; i < len(a); i++ {
		for k := 0; k < len(a); k++ {
			z += int(math.Abs(float64(a[i]-a[len(a)-1-k])))
		}
	}

	return z/2-len(a) == int(a[0])
}
