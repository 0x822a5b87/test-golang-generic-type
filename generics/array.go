package generics

type MyInt int

// IntArray 声明一个元素底层类型为int的数组
type IntArray[T ~int] []T

func FillIntArray() {
	var intArr IntArray[int]
	for i := 0; i < 10; i++ {
		intArr = append(intArr, i)
	}
}

func FillMyIntArray() {
	var intArr IntArray[MyInt]
	var i MyInt
	for i = 0; i < 10; i++ {
		intArr = append(intArr, i)
	}
}

// sum returns the sum (concatenation, for strings) of its arguments.
func sum[T ~int | ~float64 | ~string](x ...T) T {
	var t T
	for _, v := range x {
		t += v
	}
	return t
}

func o() {
	x := sum                  // illegal: the type of x is unknown
	intSum := sum[int]        // intSum has type func(x... int) int
	a := intSum(2, 3)         // a has value 5 of type int
	b := sum[float64](2.0, 3) // b has value 5.0 of type float64
	c := sum(b, -1)           // c has value 4.0 of type float64

	type sumFunc func(x ...string) string
	var f sumFunc = sum // same as var f sumFunc = sum[string]
	f = sum             // same as f = sum[string]
}
