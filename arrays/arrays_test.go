package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertSumIsCorrect(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of tails returns [2,9]", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})
	t.Run("sum of tails with empty slice as input", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func TestSumAllHeads(t *testing.T) {
	t.Run("sum of heads should return [1, 0]", func(t *testing.T) {
		got := SumAllHeads([]int{1, 2}, []int{0, 9})
		want := []int{1, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("sum of heads with empty slice as input", func(t *testing.T) {
		got := SumAllHeads([]int{1, 9}, []int{})
		want := []int{1, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Outputs: 6
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func assertSumIsCorrect(t testing.TB, got int, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
