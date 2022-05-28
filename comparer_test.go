package comparer_test

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/FS-Frost/comparer"
)

func TestAreEqual(t *testing.T) {
	type testCase struct {
		list1          []string
		list2          []string
		expectedResult bool
		expectedIndex  int
	}

	cases := []testCase{
		{
			list1:          []string{"a"},
			list2:          []string{"a"},
			expectedResult: true,
			expectedIndex:  0,
		},
		{
			list1:          []string{"a"},
			list2:          []string{"b"},
			expectedResult: false,
			expectedIndex:  0,
		},
		{
			list1:          []string{"a"},
			list2:          []string{"a", "b"},
			expectedResult: false,
			expectedIndex:  1,
		},
		{
			list1:          []string{"a", "", "b"},
			list2:          []string{"a", "", "b"},
			expectedResult: true,
			expectedIndex:  0,
		},
		{
			list1:          []string{"a", "", "b"},
			list2:          []string{"a", "", "c"},
			expectedResult: false,
			expectedIndex:  2,
		},
		{
			list1:          []string{"a", ""},
			list2:          []string{"a"},
			expectedResult: true,
			expectedIndex:  0,
		},
		{
			list1:          []string{"a", "", ""},
			list2:          []string{"a", ""},
			expectedResult: true,
			expectedIndex:  0,
		},
		{
			list1:          []string{"a", "b", "c"},
			list2:          []string{"a", "f", "c"},
			expectedResult: false,
			expectedIndex:  1,
		},
		{
			list1:          []string{"a", "b", "c"},
			list2:          []string{"a", "b", "c"},
			expectedResult: true,
			expectedIndex:  0,
		},
	}

	for i, c := range cases {
		caseNumber := i + 1
		caseName := fmt.Sprintf("%d", caseNumber)
		compareLists(t, caseName, c.list1, c.list2, c.expectedIndex, c.expectedResult)

		caseName = fmt.Sprintf("%d inverted", caseNumber)
		compareLists(t, caseName, c.list2, c.list1, c.expectedIndex, c.expectedResult)
	}
}

func TestFilesAreEqual(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getting working directory: %v", err)
	}

	lines1, err := readLines(path.Join(cwd, "data/file1.ass"))
	if err != nil {
		t.Fatalf("error reading test file: %v", err)
	}

	lines2, err := readLines(path.Join(cwd, "data/file2.ass"))
	if err != nil {
		t.Fatalf("error reading test file: %v", err)
	}

	compareLists(t, "ASS files", lines1, lines2, 19, false)
}

func compareLists(t *testing.T, caseName string, list1, list2 []string, expectedIndex int, expectedResult bool) {
	actualResult, actualIndex := comparer.AreEqual(list1, list2)
	printLists := func() {
		t.Logf("%#v\n", list1)
		t.Logf("%#v\n", list2)
	}

	if actualResult != expectedResult {
		printLists()
		t.Fatalf("case %s: equality should be %t, instead found %t with index '%d'\n", caseName, expectedResult, actualResult, actualIndex)
	}

	if actualIndex != expectedIndex {
		printLists()
		t.Fatalf("case %s: index should be %d, instead found %d\n", caseName, expectedIndex, actualIndex)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
