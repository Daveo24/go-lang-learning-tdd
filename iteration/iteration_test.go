package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("exepected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeatB := Repeat("b", 5)
	fmt.Println(repeatB)
	//Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
