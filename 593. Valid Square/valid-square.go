package main

import (
	"math"
)

// https://leetcode.com/problems/valid-square/

type Vector struct {
	x int
	y int
}

func NearEquals(a, b float64) bool {
	if math.Abs(a-b) < 1e-9 {
		return true
	}
	return false
}

func (a Vector) Dot(b Vector) float64 {
	return float64((a.x*b.x)+(a.y*b.y)) / (a.Norm() * b.Norm())
}

func (a Vector) Norm() float64 {
	return math.Sqrt(float64(a.x*a.x + a.y*a.y))
}

func PointsToVector(p1, p2 []int) Vector {
	return Vector{
		x: p2[0] - p1[0],
		y: p2[1] - p1[1],
	}
}

func CheckLengths(v1, v2, v3 Vector) bool {
	a, b, c := v1.Norm(), v2.Norm(), v3.Norm()

	// order a -> diagonal, b and c sides
	if b > a {
		a, b = b, a
	} else if c > a {
		a, c = c, a
	}

	if NearEquals(a, math.Sqrt(2)*b) && b == c {
		return true
	}
	return false
}

func CheckDotProducts(v1, v2, v3 Vector) bool {
	a, b, c := v1.Dot(v2), v2.Dot(v3), v3.Dot(v1)

	if a > b {
		a, b = b, a
	} else if a > c {
		a, c = c, a
	}

	if a == 0 && b == c && NearEquals(b, 1/math.Sqrt(2)) {
		return true
	} else {
		return false
	}
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// using p1 as pivot
	v1 := PointsToVector(p1, p2)
	v2 := PointsToVector(p1, p3)
	v3 := PointsToVector(p1, p4)

	// check dot relationships
	if CheckDotProducts(v1, v2, v3) {
		// check length
		if CheckLengths(v1, v2, v3) {
			return true
		}
	}
	return false
}
