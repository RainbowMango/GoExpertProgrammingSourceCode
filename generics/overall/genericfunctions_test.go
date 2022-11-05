package overall

import "testing"

func TestSumIntsOrFloats_Explicitly(t *testing.T) {
	scores := map[string]int64{"Jim": 80, "John": 90}
	if SumIntsOrFloats[string, int64](scores) != 170 {
		t.Fatalf("expected: 170, but got: SumIntsOrFloats(scores)")
	}

	stocks := map[int]float64{600718: 10.92, 600519: 1516.57}
	if SumIntsOrFloats[int, float64](stocks) != (10.92 + 1516.57) {
		t.Fatalf("expected: %f, bug got: %f", 10.92+1516.57, SumIntsOrFloats(stocks))
	}
}

func TestSumIntsOrFloats_Implicitly(t *testing.T) {
	scores := map[string]int64{"Jim": 80, "John": 90}
	if SumIntsOrFloats(scores) != 170 {
		t.Fatalf("expected: 170, but got: SumIntsOrFloats(scores)")
	}

	stocks := map[int]float64{600718: 10.92, 600519: 1516.57}
	if SumIntsOrFloats(stocks) != (10.92 + 1516.57) {
		t.Fatalf("expected: %f, bug got: %f", 10.92+1516.57, SumIntsOrFloats(stocks))
	}
}
