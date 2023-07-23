package models

import (
	"math"
	"testing"
)

// test creation of a point
func TestPoint(t *testing.T) {
	p := Point(4.3, -4.2, 3.1)
	if p.x != 4.3 || p.y != -4.2 || p.z != 3.1 || p.w != 1 {
		t.Errorf("Point = %v, want %v", p, Tuple{4.3, -4.2, 3.1, 1})
	}

	pointTuple := Tuple{4.3, -4.2, 3.1, 1}
	if p != pointTuple {
		t.Errorf("Point %v isn't the same as Tuple %v", p, pointTuple)
	}

}

// test creation of a vector
func TestVector(t *testing.T) {
	v := Vector(4.3, -4.2, 3.1)
	if v.x != 4.3 || v.y != -4.2 || v.z != 3.1 || v.w != 0 {
		t.Errorf("Vector = %v, want %v", v, Tuple{4.3, -4.2, 3.1, 0})
	}

	vectorTuple := Tuple{4.3, -4.2, 3.1, 0}
	if v != vectorTuple {
		t.Errorf("Vector %v isn't the same as Tuple %v", v, vectorTuple)
	}
}

func TestTupleComparison(t *testing.T) {

	tuple1 := Tuple{1.2, 3.4, 5.6, 0}
	tuple2 := Tuple{1.2, 3.4, 5.6, 0}

	if areEqual(tuple1, tuple2) == false {
		t.Errorf("Tuple %v isn't the same as Tuple %v", tuple1, tuple2)
	}
}

func TestAddition(t *testing.T) {
	a1 := Tuple{3.0, -2.0, 5.0, 1}
	a2 := Tuple{-2.0, 3.0, 1.0, 0}

	if (!areEqual(add(a1, a2), Tuple{1.0, 1.0, 6.0, 1})) {
		t.Errorf("Tuple %v isn't the same as the addition of Tuples %v and %v", Tuple{1.0, 1.0, 6.0, 1}, a1, a2)
	}

}

func TestSubtraction(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	if subtract(p1, p2) != Vector(-2, -4, -6) {
		t.Errorf("Tuple %v isn't the same as the subtraction of Tuples %v and %v", Vector(-2, -4, -6), p1, p2)
	}
}

func TestAddingAVectorToAPointReturnsAPoint(t *testing.T) {
	p := Point(3.2, -2.2, 5)
	v := Vector(-2.2, 3.2, 2)
	expectedResult := Point(1, 1, 7)

	if add(p, v) != expectedResult {
		t.Errorf("Tuple %v isn't the same as the addition of Point %v and Vector %v", expectedResult, p, v)
	}
}

func TestSubtractingAVectorFromAPointReturnsAPoint(t *testing.T) {
	p := Point(3.2, -2.2, 5)
	v := Vector(-2.2, 3.2, 2)
	expectedResult := Point(5.4, -5.4, 3)

	if subtract(p, v) != expectedResult {
		t.Errorf("Tuple %v isn't the same as the subtraction of Point %v and Vector %v", expectedResult, p, v)
	}
}

func TestSubtractingTwoVectorsIsAVector(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	expectedResult := Vector(-2, -4, -6)

	if subtract(v1, v2) != expectedResult {
		t.Errorf("Tuple %v isn't the same as the subtraction of Vector %v and Vector %v", expectedResult, v1, v2)
	}
}

func TestNegatingATuple(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	expectedResult := Tuple{-1, 2, -3, 4}

	if negate(tuple) != expectedResult {
		t.Errorf("Tuple %v isn't the same as the negation of Tuple %v", expectedResult, tuple)
	}
}

func TestMultiplyingATupleByAScalar(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	expectedResult := Tuple{3, -6, 9.0, -12}

	if multiply(tuple, 3) != expectedResult {
		t.Errorf("Tuple %v isn't the same as the multiplication of Tuple %v by 3", expectedResult, tuple)
	}
}

func TestMultiplyingATupleByAFraction(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	expectedResult := Tuple{0.5, -1, 1.5, -2}
	result := multiply(tuple, 0.5)

	if result != expectedResult {
		t.Errorf("Tuple %v isn't the same as the multiplication of Tuple %v by 0.5 (%v)", expectedResult, tuple, result)
	}
}

func TestDividingATupleByAScalar(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	expectedResult := Tuple{0.5, -1, 1.5, -2}
	result := divide(tuple, 2)

	if result != expectedResult {
		t.Errorf("Tuple %v isn't the same as the division of Tuple %v by 2 (%v)", expectedResult, tuple, result)
	}
}

func TestDividingATupleByAFraction(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	expectedResult := Tuple{2, -4, 6, -8}
	result := divide(tuple, 0.5)

	if result != expectedResult {
		t.Errorf("Tuple %v isn't the same as the division of Tuple %v by 0.5 (%v)", expectedResult, tuple, result)
	}
}

func TestComputingTheMagnitudeOfVector1_0_0(t *testing.T) {
	v := Vector(1, 0, 0)
	expectedResult := 1.0
	result := magnitude(v)

	if result != expectedResult {
		t.Errorf("Magnitude of %v is %v, want %v", v, result, expectedResult)
	}
}

func TestComputingTheMagnitudeOfVector0_1_0(t *testing.T) {
	v := Vector(0, 1, 0)
	expectedResult := 1.0
	result := magnitude(v)

	if result != expectedResult {
		t.Errorf("Magnitude of %v is %v, want %v", v, result, expectedResult)
	}
}

func TestComputingTheMagnitudeOfVector0_0_1(t *testing.T) {
	v := Vector(0, 0, 1)
	expectedResult := 1.0
	result := magnitude(v)

	if result != expectedResult {
		t.Errorf("Magnitude of %v is %v, want %v", v, result, expectedResult)
	}
}

func TestComputingTheMagnitudeOfVector1_2_3(t *testing.T) {
	v := Vector(1, 2, 3)
	expectedResult := math.Sqrt(14.0)
	result := magnitude(v)

	if result != expectedResult {
		t.Errorf("Magnitude of %v is %v, want %v", v, result, expectedResult)
	}
}

func TestComputingTheMagnitudeOfVectorNeg1_Neg2_Neg3(t *testing.T) {
	v := Vector(-1, -2, -3)
	expectedResult := math.Sqrt(14.0)
	result := magnitude(v)

	if result != expectedResult {
		t.Errorf("Magnitude of %v is %v, want %v", v, result, expectedResult)
	}
}

func TestNormalizingVector4_0_0(t *testing.T) {
	v := Vector(4, 0, 0)
	expectedResult := Vector(1, 0, 0)
	result := normalize(v)

	if result != expectedResult {
		t.Errorf("Normalizing %v is %v, want %v", v, result, expectedResult)
	}
}

func TestNormalizingVector1_2_3(t *testing.T) {
	v := Vector(1, 2, 3)
	expectedResult := Vector(0.2672612419124244, 0.5345224838248488, 0.8017837257372732)
	result := normalize(v)

	if result != expectedResult {
		t.Errorf("Normalizing %v is %v, want %v", v, result, expectedResult)
	}
}

func TestTheMagnitudeOfANormalizedVector(t *testing.T) {
	v := Vector(1, 2, 3)
	norm := normalize(v)
	expectedResult := 1.0
	result := magnitude(norm)

	if result != expectedResult {
		t.Errorf("Magnitude of %v is %v, want %v", norm, result, expectedResult)
	}
}

func TestTheDotProductOfTwoTuples(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	expectedResult := 20.0
	result := dot(a, b)

	if result != expectedResult {
		t.Errorf("Dot product of %v and %v is %v, want %v", a, b, result, expectedResult)
	}
}

// Scenario: The cross product of two vectors Given a ← vector(1, 2, 3)
// And b ← vector(2, 3, 4)
// Then cross(a, b) = vector(-1, 2, -1)
// And cross(b, a) = vector(1, -2, 1)
func TestTheCrossProductOfTwoVectors(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	expectedResult := Vector(-1, 2, -1)
	result := cross(a, b)

	if result != expectedResult {
		t.Errorf("Cross product of %v and %v is %v, want %v", a, b, result, expectedResult)
	}

	expectedResult = Vector(1, -2, 1)
	result = cross(b, a)

	if result != expectedResult {
		t.Errorf("Cross product of %v and %v is %v, want %v", a, b, result, expectedResult)
	}
}
