package main

import "testing"

func TestDLDist(t *testing.T) {
	cases := []struct {
		a, b string
		d    int
	}{
		{"", "", 0},
		{"", "ab", 2},
		{"a", "ab", 1},
		{"abc", "adc", 1},
		{"ab", "ba", 1},
		{"ca", "abc", 2},
	}

	for _, c := range cases {
		dist1, dist2 := DLDist(c.a, c.b), DLDist(c.b, c.a)
		if dist1 != c.d {
			t.Errorf("dist(%s, %s) = %d != %d", c.a, c.b, dist1, c.d)
		}
		if dist2 != c.d {
			t.Errorf("dist(%s, %s) = %d != %d", c.b, c.a, dist2, c.d)
		}
	}
}
