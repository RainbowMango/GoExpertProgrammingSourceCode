package slice

// 检测多行输出
func ExampleSliceString() {
	SliceString()
	// OutPut:
	// baseStr: Hello World!
	// baseStr type: string
	// newStr: Hello
	// newStr type: string
}

func ExampleSliceArray() {
	SliceArray()
	// OutPut:
	// baseArray: [Hello World        ]
	// baseArray type: [10]string
	// newSlice: [Hello World   ]
	// newSlice type: []string
}

func ExampleSliceCap() {
	SliceCap()
	// OutPut:
	// newSlice: [0 0 0]
}

func ExampleExtendSliceArray() {
	ExtendSliceArray()
	// OutPut:
	// cap(b): 3
}
