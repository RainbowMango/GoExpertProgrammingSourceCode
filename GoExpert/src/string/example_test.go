package stringExample

import "testing"

func TestGetSliceByString(t *testing.T) {
    slice := GetSliceByString("Hello,World")

    if string(slice) != "Hello,World" {
        t.Fatalf("GetSliceByString failed. expected: Hello,World, actual: %s", string(slice))
    }
}