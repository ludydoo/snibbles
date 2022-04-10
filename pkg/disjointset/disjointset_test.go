package disjointset

import "testing"

func TestFind(t *testing.T) {
	tcs := []struct {
		name string
		set  []int
		want int
	}{
		{
			name: "1",
			set:  []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 1,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			ds := MakeSet(tc.want)
			for _, i := range tc.set {
				ds.Union(tc.want, i)
			}
			for _, i := range tc.set {
				if ds.Find(i) != tc.want {
					t.Errorf("Find(%d) = %d, want %d", i, ds.Find(i), tc.want)
				}
			}
		})
	}
}
