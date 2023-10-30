package binarySearch

import "testing"

func TestBs0704(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	// test 1
	if got := Bs0704(nums, -2); got != -1 {
		t.Errorf("test 1, got: %d, want: -1", got)
	}

	// test 2
	if got := Bs0704(nums, -1); got != 0 {
		t.Errorf("test 2, got: %d, want: 0", got)
	}

	// test 3
	if got := Bs0704(nums, 3); got != 2 {
		t.Errorf("test 3, got: %d, want: 2", got)
	}

	// test 4
	if got := Bs0704(nums, 12); got != 5 {
		t.Errorf("test 4, got: %d, want: 6", got)
	}

	// test 5
	if got := Bs0704(nums, 15); got != -1 {
		t.Errorf("test 5, got: %d, want: -1", got)
	}

	// test 7
	if got := Bs0704([]int{}, 2); got != -1 {
		t.Errorf("test 6, got: %d, want: -1", got)
	}
}

func TestBs0035(t *testing.T) {
	// test 1
	nums := []int{1, 3, 5, 6}
	if got := Bs0035(nums, -1); got != 0 {
		t.Errorf("test 1, got: %d, want: 0", got)
	}

	// test 2
	if got := Bs0035(nums, 1); got != 0 {
		t.Errorf("test 2, got: %d, want: 0", got)
	}

	// test 3
	if got := Bs0035(nums, 6); got != 3 {
		t.Errorf("test 3, got: %d, want: 3", got)
	}

	if got := Bs0035([]int{}, 8); got != 0 {
		t.Errorf("test 3, got: %d, want: 0", got)
	}
}
