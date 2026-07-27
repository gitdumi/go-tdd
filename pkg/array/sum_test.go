package array

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assert(t, got, want, numbers[:])
	})
}

func assert(t testing.TB, got, want int, input []int) {
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, input)
	}
}
