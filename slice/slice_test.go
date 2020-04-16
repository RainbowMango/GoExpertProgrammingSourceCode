package _slice

func ExampleSliceDeclare() {
	SliceDeclare()
	// Output:
	// true
}

func ExampleSliceLiteral() {
	SliceLiteral()
	// Output:
	// false
	// false
}

func ExampleSliceInitMake() {
	SliceInitMake()
	// Output:
	// false
	// false
}

func ExampleSliceCut() {
	SliceCut()
	// Output:
	// [1 2]
	// [1]
}

func ExampleSliceInitNew() {
	SliceInitNew()
	// Output:
	// true
}
