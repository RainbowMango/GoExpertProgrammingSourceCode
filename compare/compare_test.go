package compare

func ExampleStructTypeCompare() {
	StructTypeCompare()
	// Output:
	// false
}

func ExampleInterfaceTypeCompare() {
	InterfaceTypeCompare()
	// Output:
	// true
	// false
}

func ExampleInterfaceAndStructCompare() {
	InterfaceAndStructCompare()
	// Output:
	// true
}

func ExampleArrayCompare() {
	ArrayCompare()
	// Output:
	// false
	// true
}
