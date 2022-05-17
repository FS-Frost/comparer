package comparer_test

import (
	"testing"

	"github.com/FS-Frost/comparer"
)

func TestAreEqual(t *testing.T) {
	t.Skip()
	a1 := []string{"a", "b", "c"}
	a2 := []string{"a", "b", "c"}
	areEqual, index := comparer.AreEqual(a1, a2)

	if index != 0 {
		t.Errorf("index should be 0, instead: %d\n", index)
	}

	if areEqual == false {
		t.Errorf("should be equal, instead '%t' at index '%d'\n", areEqual, index)
	}
}

func TestAreNotEqual(t *testing.T) {
	t.Skip()
	a1 := []string{"a", "b", "c"}
	a2 := []string{"a", "f", "c"}
	areEqual, index := comparer.AreEqual(a1, a2)

	if index != 1 {
		t.Errorf("index should be 1, instead: %d\n", index)
	}

	if areEqual == true {
		t.Errorf("should not be equal, found difference at index '%d'\n", index)
	}
}

func TestDistinctLenArraysAreNotEqual(t *testing.T) {
	type testCase struct {
		list1          []string
		list2          []string
		expectedResult bool
		expectedIndex  int
	}

	cases := []testCase{
		{
			list1:          []string{"a"},
			list2:          []string{"a", "b"},
			expectedResult: false,
			expectedIndex:  1,
		},
		{
			list1:          []string{"a"},
			list2:          []string{"a"},
			expectedResult: true,
			expectedIndex:  0,
		},
		{
			list1:          []string{"a", "b", "c"},
			list2:          []string{"a", "b"},
			expectedResult: false,
			expectedIndex:  2,
		},
		{
			list1:          []string{"a", "", "b"},
			list2:          []string{"a", "", "b"},
			expectedResult: true,
			expectedIndex:  0,
		},
	}

	for i, c := range cases {
		compareDistinctLenArrays(t, i+1, c.list1, c.list2, c.expectedIndex, c.expectedResult)
	}
}

func compareDistinctLenArrays(t *testing.T, caseIndex int, list1, list2 []string, expectedIndex int, expectedResult bool) {
	areEqual, index := comparer.AreEqual(list1, list2)

	if index != expectedIndex {
		t.Errorf("case %d: index should be %d, instead: %d\n", caseIndex, index, expectedIndex)
	}

	if areEqual != expectedResult {
		t.Errorf("case %d: equality should be %t, instead found difference at index '%d'\n", caseIndex, expectedResult, index)
	}
}
