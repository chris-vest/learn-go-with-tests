package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeatCount := 10
	repeated := Repeat("a", repeatCount)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}
