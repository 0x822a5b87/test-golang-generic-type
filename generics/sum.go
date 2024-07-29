package generics

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Sum sums the values of map m.
// It supports both int64 and float64 as typed of map values
func Sum[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, value := range m {
		s += value
	}
	return s
}

// Number declare Number interface type to use as a type constraint
// Essentially, you’re moving the union from the function declaration into a new type constraint.
// That way, when you want to constrain a type parameter to either int64 or float64,
// you can use this Number type constraint instead of writing out int64 | float64.
type Number interface {
	int64 | float64
}

// SumNumber a function with generics with constraint that indicate map's value has type of Number
func SumNumber[K comparable, V Number](data map[K]V) V {
	var s V
	for _, value := range data {
		s += value
	}

	return s
}

type Id[T interface{}] interface {
	GetId() T
}

type PageView struct {
	IntId int
}

func (p PageView) GetId() int {
	return p.IntId
}

type IntNumber interface {
	int | int64
}

type GenericPageView[T IntNumber] struct {
	IntNumberId T
}

func (g GenericPageView[T]) GetId() T {
	return g.IntNumberId
}

// generics with interface return type

type Gender interface {
	GetGender() string
}

type Male struct {
}

func (m Male) GetGender() string {
	return "male"
}

type Female struct {
}

func (m Female) GetGender() string {
	return "male"
}

func NewGenderValue[T Gender](gender string) T {
	var g Gender
	if gender == "male" {
		g = Male{}
	} else {
		g = Female{}
	}
	return g.(T)
}
