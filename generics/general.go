package generics

// ERROR
// Cannot use a type parameter as RHS in type declaration
// type FloatNumber[T float64 | float32] T

// FloatNumber use FloatNumber interface instead of general type
type FloatNumber interface {
	float32 | float64
}

// IntSlice define generic type slice
type IntSlice[T int | int8 | int16 | int32 | int64] []T

func testSmallIntSlice() {
	var intSlice IntSlice[int]
	for i := 0; i < 10; i++ {
		intSlice[i] = i
	}
}

// GenericMap define generics map
type GenericMap[K comparable, V FloatNumber] map[K]V

// Tree define generics tree
type Tree[T interface{}] struct {
	Left  *Tree[T]
	Right *Tree[T]
	Val   T
}
