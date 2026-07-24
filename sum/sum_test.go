package sum

import "testing"

var tests = []struct {
	input []int
	want  int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{1, 3, 2}, 6},
}

func TestLoop(t *testing.T) {
	for _, test := range tests {
		if got := Loop(test.input); got != test.want {
			t.Errorf("Loop(%v) = %d, want %d",
				test.input, got, test.want)
		}
	}
}

func TestDaC(t *testing.T) {
	for _, test := range tests {
		if got := DaC(test.input); got != test.want {
			t.Errorf("DaC(%v) = %d, want %d",
				test.input, got, test.want)
		}
	}
}

func getInts(n int) []int {
	var ints []int
	for i := 0; i < n; i++ {
		ints = append(ints, i)
	}
	return ints
}

func BenchmarkLoop_10(b *testing.B) {
	ints := getInts(10)
	for b.Loop() {
		Loop(ints)
	}
}

func BenchmarkLoop_20(b *testing.B) {
	ints := getInts(20)
	for b.Loop() {
		Loop(ints)
	}
}

func BenchmarkLoop_30(b *testing.B) {
	ints := getInts(30)
	for b.Loop() {
		Loop(ints)
	}
}

func BenchmarkDaC_10(b *testing.B) {
	ints := getInts(10)
	for b.Loop() {
		DaC(ints)
	}
}

func BenchmarkDaC_20(b *testing.B) {
	ints := getInts(20)
	for b.Loop() {
		DaC(ints)
	}
}

func BenchmarkDaC_30(b *testing.B) {
	ints := getInts(30)
	for b.Loop() {
		DaC(ints)
	}
}
