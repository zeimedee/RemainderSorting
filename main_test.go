package main

import "testing"

func TestRemainderSorting(t *testing.T) {
	arr := []string{"a", "ab", "bc", "abc"}
	sorted := []string{"abc", "a", "ab", "bc"}

	results := RemainderSorting(arr)
	for i := range results {
		if results[i] != sorted[i] {
			t.Errorf("failed")
			break
		}
	}
}

func TestFullRemainderSorting(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"alex", "cat", "diego"}, []string{"cat", "alex", "diego"}},
		{[]string{"a", "ab", "bc", "abc"}, []string{"abc", "a", "ab", "bc"}},
		{[]string{"colorado", "utah", "wisconsin", "oregon"}, []string{"oregon", "wisconsin", "utah", "colorado"}},
	}

	for _, test := range tests {
		output := RemainderSorting(test.input)
		for i := range output {
			if output[i] != test.expected[i] {
				t.Error("Test failed: {} inputted: {}, {} expected, received: {}", test.input, test.expected, output)
			}
		}

	}
}
