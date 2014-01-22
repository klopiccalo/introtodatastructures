package main

import "testing"

func TestAsOccuranceMap(t *testing.T) {
	words := []string{
		"foo", "bar", "baz",
		"foo", "bar",
		"foo",
	}
	expected := map[string]int{
		"foo": 3,
		"bar": 2,
		"baz": 1,
	}
	actual := asOccuranceMap(words)
	for word, expectedCount := range expected {
		if actualCount := actual[word]; actualCount != expectedCount {
			t.Fatalf("Expected word %v with count of %v, got %v",
				word, expectedCount, actualCount)
		}
	}
}

func TestReverseOccuranceMap(t *testing.T) {
	input := map[string]int{
		"foo":  3,
		"bar":  2,
		"baz":  2,
		"bang": 1,
		"bing": 1,
	}
	expectedMax := 3
	expected := map[int][]string{
		3: []string{"foo"},
		2: []string{"bar", "baz"},
		1: []string{"bang", "bing"},
	}
	actualMax, actual := reverseOccuranceMap(input)
	if actualMax != 3 {
		t.Fatalf("Expected max occurance %v, got %v", expectedMax, actualMax)
	}
	for count, expectedWords := range expected {
		for _, expectedWord := range expectedWords {
			found := false
			for _, actualWord := range actual[count] {
				found = found || actualWord == expectedWord
			}
			if !found {
				t.Fatalf("Expected word %v with count of %v",
					expectedWord, count)
			}
		}
	}
}
