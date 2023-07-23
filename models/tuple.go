package models

import (
	"fmt"
	"math"
)

func SayHello() {
	fmt.Println("Hello everyone from package1")
}

type Tuple struct {
	x, y, z float64
	w       int8
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func areEqual(tuple1, tuple2 Tuple) bool {
	return tuple1 == tuple2
}

func add(tuple1, tuple2 Tuple) Tuple {
	return Tuple{tuple1.x + tuple2.x, tuple1.y + tuple2.y, tuple1.z + tuple2.z, tuple1.w + tuple2.w}
}

func subtract(tuple1, tuple2 Tuple) Tuple {
	return Tuple{tuple1.x - tuple2.x, tuple1.y - tuple2.y, tuple1.z - tuple2.z, tuple1.w - tuple2.w}
}

func negate(tuple Tuple) Tuple {
	return subtract(Tuple{0, 0, 0, 0}, tuple)
}

func multiply(tuple Tuple, scalar float64) Tuple {
	return Tuple{tuple.x * scalar, tuple.y * scalar, tuple.z * scalar, (int8)(float64(tuple.w) * scalar)}
}

func divide(tuple Tuple, scalar float64) Tuple {
	return Tuple{tuple.x / scalar, tuple.y / scalar, tuple.z / scalar, (int8)(float64(tuple.w) / scalar)}
}

func magnitude(tuple Tuple) float64 {
	return math.Sqrt(math.Pow(tuple.x, 2) + math.Pow(tuple.y, 2) + math.Pow(tuple.z, 2))
}

func normalize(tuple Tuple) Tuple {
	return divide(tuple, magnitude(tuple))
}

func dot(tuple1, tuple2 Tuple) float64 {
	return tuple1.x*tuple2.x + tuple1.y*tuple2.y + tuple1.z*tuple2.z + float64(tuple1.w*tuple2.w)
}

func cross(tuple1, tuple2 Tuple) Tuple {
	return Vector(
		tuple1.y*tuple2.z-tuple1.z*tuple2.y,
		tuple1.z*tuple2.x-tuple1.x*tuple2.z,
		tuple1.x*tuple2.y-tuple1.y*tuple2.x)
}
