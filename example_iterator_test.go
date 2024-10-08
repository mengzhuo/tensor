package tensor

import "fmt"

// This is an example of how to use `IteratorFromDense` from a row-major Dense tensor
func Example_iteratorRowmajor() {
	T := New(WithShape(2, 3), WithBacking([]float64{0, 1, 2, 3, 4, 5}))
	it := IteratorFromDense(T)
	fmt.Printf("T:\n%v\n", T)

	for i, err := it.Start(); err == nil; i, err = it.Next() {
		fmt.Printf("i: %d, coord: %v\n", i, it.Coord())
	}

	// Output:
	// T:
	// ⎡0  1  2⎤
	// ⎣3  4  5⎦
	//
	// i: 0, coord: [0 1]
	// i: 1, coord: [0 2]
	// i: 2, coord: [1 0]
	// i: 3, coord: [1 1]
	// i: 4, coord: [1 2]
	// i: 5, coord: [0 0]

}

// This is an example of using `IteratorFromDense` on a col-major Dense tensor. More importantly
// this example shows the order of the iteration.
func Example_iteratorcolMajor() {
	T := New(WithShape(2, 3), WithBacking([]float64{0, 1, 2, 3, 4, 5}), AsFortran(nil))
	it := IteratorFromDense(T)
	fmt.Printf("T:\n%v\n", T)

	for i, err := it.Start(); err == nil; i, err = it.Next() {
		fmt.Printf("i: %d, coord: %v\n", i, it.Coord())
	}

	// Output:
	// T:
	// ⎡0  2  4⎤
	// ⎣1  3  5⎦
	//
	// i: 0, coord: [0 1]
	// i: 2, coord: [0 2]
	// i: 4, coord: [1 0]
	// i: 1, coord: [1 1]
	// i: 3, coord: [1 2]
	// i: 5, coord: [0 0]

}

func ExampleSlice() {
	T := New(WithShape(3, 3), WithBacking(Range(Float64, 0, 9)))
	S, err := T.Slice(makeRS(1, 3), makeRS(1, 3))
	if err != nil {
		fmt.Printf("Err %v\n", err)
		return
	}
	fmt.Printf("S (requires iterator? %t)\n%v\n", S.(*Dense).RequiresIterator(), S)
	it := IteratorFromDense(S.(*Dense))
	for i, err := it.Start(); err == nil; i, err = it.Next() {
		fmt.Printf("i %d, coord %v\n", i, it.Coord())
	}

	// Output:
	// S (requires iterator? true)
	// ⎡4  5⎤
	// ⎣7  8⎦
	//
	// i 0, coord [0 1]
	// i 1, coord [1 0]
	// i 3, coord [1 1]
	// i 4, coord [0 0]

}
