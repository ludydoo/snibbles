package bst2

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func clean(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimPrefix(str, "\n")
	return str
}

func TestTraverseInOrder(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		expect []int
	}{
		{
			name:   "empty tree",
			root:   New(),
			expect: nil,
		}, {
			name:   "single Node tree",
			root:   New(1),
			expect: []int{1},
		}, {
			name:   "two Node tree",
			root:   New(1, 2),
			expect: []int{1, 2},
		}, {
			name:   "three Node tree",
			root:   New(2, 1, 3),
			expect: []int{1, 2, 3},
		}, {
			name:   "four Node tree",
			root:   New(2, 1, 3, 4),
			expect: []int{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			assert.Equal(t, testCase.expect, TraverseInOrder(testCase.root))
		})
	}

}

func TestTraversePreOrder(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		expect []int
	}{
		{
			name:   "empty tree",
			root:   New(),
			expect: nil,
		}, {
			name:   "single Node tree",
			root:   New(1),
			expect: []int{1},
		}, {
			name:   "two Node tree",
			root:   New(1, 2),
			expect: []int{1, 2},
		}, {
			name:   "three Node tree",
			root:   New(2, 1, 3),
			expect: []int{2, 1, 3},
		}, {
			name:   "four Node tree",
			root:   New(2, 1, 3, 4),
			expect: []int{2, 1, 3, 4},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			assert.Equal(t, testCase.expect, TraversePreOrder(testCase.root))
		})
	}
}

func TestTraversePostOrder(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		expect []int
	}{
		{
			name:   "empty tree",
			root:   New(),
			expect: nil,
		}, {
			name:   "single Node tree",
			root:   New(1),
			expect: []int{1},
		}, {
			name:   "two Node tree",
			root:   New(1, 2),
			expect: []int{2, 1},
		}, {
			name:   "three Node tree",
			root:   New(2, 1, 3),
			expect: []int{1, 3, 2},
		}, {
			name:   "four Node tree",
			root:   New(2, 1, 3, 4),
			expect: []int{1, 4, 3, 2},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			assert.Equal(t, testCase.expect, TraversePostOrder(testCase.root))
		})
	}
}

func TestTraverseLevelOrder(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		expect []int
	}{
		{
			name:   "empty tree",
			root:   New(),
			expect: nil,
		}, {
			name:   "single Node tree",
			root:   New(1),
			expect: []int{1},
		}, {
			name:   "two Node tree",
			root:   New(1, 2),
			expect: []int{1, 2},
		}, {
			name:   "three Node tree",
			root:   New(2, 1, 3),
			expect: []int{2, 1, 3},
		}, {
			name:   "four Node tree",
			root:   New(2, 1, 3, 4),
			expect: []int{2, 1, 3, 4},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			assert.Equal(t, testCase.expect, TraverseLevelOrder(testCase.root))
		})
	}
}

func TestTraverseReverseLevelOrder(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		expect []int
	}{
		{
			name:   "empty tree",
			root:   New(),
			expect: nil,
		}, {
			name:   "single Node tree",
			root:   New(1),
			expect: []int{1},
		}, {
			name:   "two Node tree",
			root:   New(1, 2),
			expect: []int{2, 1},
		}, {
			name:   "three Node tree",
			root:   New(2, 1, 3),
			expect: []int{1, 3, 2},
		}, {
			name:   "four Node tree",
			root:   New(2, 1, 3, 4),
			expect: []int{4, 1, 3, 2},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			assert.Equal(t, testCase.expect, TraverseReverseLevelOrder(testCase.root))
		})
	}
}

func TestPrintFromLeft(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		expect []int
	}{
		{
			name:   "empty tree",
			root:   New(),
			expect: nil,
		}, {
			name:   "single Node tree",
			root:   New(1),
			expect: []int{1},
		}, {
			name:   "two Node tree",
			root:   New(1, 2),
			expect: []int{1, 2},
		}, {
			name:   "three Node tree",
			root:   New(2, 1, 3),
			expect: []int{2, 1},
		}, {
			name:   "four Node tree",
			root:   New(2, 1, 3, 4),
			expect: []int{2, 1, 4},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			assert.Equal(t, testCase.expect, PrintFromLeft(testCase.root))
		})
	}
}

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		expect bool
	}{
		{
			name:   "empty tree",
			root:   New(),
			expect: true,
		}, {
			name:   "single Node tree",
			root:   New(1),
			expect: true,
		}, {
			name:   "two Node tree",
			root:   New(1, 2),
			expect: false,
		}, {
			name:   "three Node tree",
			root:   New(2, 1, 3),
			expect: true,
		}, {
			name:   "four Node tree",
			root:   New(2, 1, 3, 4),
			expect: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			assert.Equal(t, testCase.expect, IsSymmetric(testCase.root))
		})
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		insert int
		expect *Node
	}{
		{
			name:   "empty",
			root:   New(),
			expect: New(1),
			insert: 1,
		}, {
			name:   "tree with 1 item",
			root:   New(1),
			expect: New(1, 2),
			insert: 2,
		}, {
			name:   "tree with 2 items",
			root:   New(1, 2),
			expect: New(1, 2, 3),
			insert: 3,
		}, {
			name:   "existing",
			root:   New(1, 2),
			expect: New(1, 2),
			insert: 2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("\n%s\n", Print(testCase.root))
			actual := Insert(testCase.root, testCase.insert)
			expected := testCase.expect
			assert.Equal(t, expected, actual)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name   string
		root   *Node
		delete int
		expect *Node
	}{
		{
			name:   "empty",
			root:   New(),
			expect: New(),
			delete: 1,
		}, {
			name:   "tree with 1 item",
			root:   New(1),
			expect: New(),
			delete: 1,
		}, {
			name:   "tree with 2 item",
			root:   New(1, 2),
			expect: New(1),
			delete: 2,
		}, {
			name:   "tree with 3 items",
			root:   New(2, 1, 3),
			expect: New(2, 1),
			delete: 3,
		}, {
			name:   "delete root",
			root:   New(2, 1, 3),
			expect: New(3, 1),
			delete: 2,
		}, {
			name:   "delete Left",
			root:   New(2, 1, 3),
			expect: New(2, 3),
			delete: 1,
		}, {
			name:   "delete Right",
			root:   New(2, 1, 3),
			expect: New(2, 1),
			delete: 3,
		}, {
			name:   "delete Left parent",
			root:   New(5, 3, 7, 2, 1),
			expect: New(5, 2, 7, 1),
			delete: 3,
		}, {
			name:   "delete Right parent",
			root:   New(5, 3, 7, 1, 9, 8),
			expect: New(5, 3, 1, 9, 8),
			delete: 7,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			t.Logf("\nbefore\n%s\n", Print(testCase.root))
			actual := DeleteKey(testCase.root, testCase.delete)
			expected := testCase.expect
			t.Logf("\nactual\n%s\n", Print(actual))
			t.Logf("\nexpected\n%s\n", Print(expected))

			assert.Equal(t, expected, actual)
		})
	}
}

func TestPrint(t *testing.T) {
	tcs := []struct {
		name   string
		n      *Node
		expect string
	}{
		{
			name: "single",
			n:    New(1),
			expect: clean(`
1
`),
		}, {
			name: "Right child",
			n:    New(1, 2),
			expect: clean(`
1
└─╮
  2
`),
		}, {
			name: "Left child",
			n:    New(2, 1),
			expect: clean(`
  2
╭─┘
1
`),
		}, {
			name: "both children",
			n:    New(2, 1, 3),
			expect: clean(`
  2
╭─┴─╮
1   3`),
		}, {
			name: "deep Right",
			n:    New(2, 1, 4, 3),
			expect: clean(`
  2
╭─┴───╮
1     4
    ╭─┘
    3
`),
		},
		{
			name: "larger numbers",
			n:    New(70, 60, 80),
			expect: clean(`
   70
╭──┴──╮
60    80
`),
		},

		{
			name: "big tree",
			n:    New(50, 30, 70, 20, 40, 60, 80, 45, 65, 78, 85, 79, 100, 200),
			expect: clean(`
            50
   ╭────────┴────────╮
   30                70
╭──┴──╮        ╭─────┴────────╮
20    40       60             80
      └──╮     └──╮     ╭─────┴──╮
         45       65    78       85
                        └──╮     └──╮
                           79       100
                                    └───╮
                                        200
`),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Print(tc.n)
			t.Logf("actual:\n%s\n", actual)
			t.Logf("expected:\n%s\n", tc.expect)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func TestMergeStrings(t *testing.T) {
	tcs := []struct {
		name   string
		s1     string
		s2     string
		x      int
		y      int
		expect string
	}{
		{
			name:   "empty",
			s1:     "",
			s2:     "",
			x:      0,
			y:      0,
			expect: "",
		}, {
			name:   "s1 empty",
			s1:     "",
			s2:     "a",
			x:      0,
			y:      0,
			expect: "a",
		}, {
			name:   "s2 empty",
			s1:     "a",
			s2:     "",
			x:      0,
			y:      0,
			expect: "a",
		}, {
			name:   "horizontal offset",
			s1:     "a",
			s2:     "b",
			x:      1,
			y:      0,
			expect: "ab",
		}, {
			name: "vertical offset",
			s1:   "a",
			s2:   "b",
			x:    0,
			y:    1,
			expect: clean(`
a
b
`),
		}, {
			name: "both offsets",
			s1:   "a",
			s2:   "b",
			x:    1,
			y:    1,
			expect: clean(`
a
 b
`),
		}, {
			name: "simple multiline",
			x:    0,
			y:    0,
			s1: clean(`
111
111
111
`),
			s2: clean(`
22
22
`),
			expect: clean(`
221
221
111
`),
		}, {
			name: "multiline offset",
			x:    1,
			y:    2,
			s1: clean(`
111
111
111
`),
			s2: clean(`
22
22
`),
			expect: clean(`
111
111
122
 22
`),
		}, {
			name: "negative offset",
			x:    -1,
			y:    -1,
			s1: clean(`
111
111
111
`),
			s2: clean(`
222
222
222
`),
			expect: clean(`
222 
2221
2221
 111
`),
		}, {
			name:   "with spaces",
			x:      0,
			y:      0,
			s1:     "1 1",
			s2:     " 2 ",
			expect: "121",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := mergeStrs(tc.s1, tc.s2, tc.x, tc.y)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
