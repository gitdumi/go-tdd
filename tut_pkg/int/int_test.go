package int

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assertResults(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(12, 28)
	fmt.Println(sum)
	// Output: 40
}

func assertResults(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
