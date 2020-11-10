package test

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 6)
	expect := 10

	if result != expect {
		t.Fatalf("Didn't work: We should have gotten %d but instead we got %d", expect, result)
	}
}
