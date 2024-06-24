package lexorank

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRank(t *testing.T) {
	testCases := []struct {
		prev, next string
		want       string
		wantErr    bool
	}{
		{"", "", "U", false},
		{"", "0", "0", false},
		{"", "2", "1", false},
		{"x", "", "y", false},
		{"aaaa", "aaab", "aaaaU", false},
		{"aaaa", "aaac", "aaab", false},
		{"az", "b", "azU", false},
		{"a", "d", "b", false},
		{"a", "c", "b", false},
		{"aaaa", "aaaa", "aaaa", false},
		{"a", "a0", "a", false},
		{"a", "", "m", false},
		{"", "0", "0", false},
		{"z", "", "z", false},
		{"_", "", "", true},
		{"", "*", "", true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: (%q..%q)", i, tc.prev, tc.next), func(t *testing.T) {
			got, err := Rank(tc.prev, tc.next)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Errorf("unexpected error: %v", err)
				return
			}
			if got != tc.want {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}

func TestRankN(t *testing.T) {
	testCases := []struct {
		prev, next string
		n          int
		want       []string
		wantErr    bool
	}{
		{"", "", 5, []string{"U0", "U1", "U2", "U3", "U4"}, false},
		{"az", "b", 3, []string{"azU0", "azU1", "azU2"}, false},
		{"a", "d", 2, []string{"b0", "b1"}, false},
		{"a", "c", 4, []string{"b0", "b1", "b2", "b3"}, false},
		{"0", "z", 20, []string{"U00", "U01", "U02", "U03", "U04", "U05", "U06", "U07", "U08", "U09", "U10", "U11", "U12", "U13", "U14", "U15", "U16", "U17", "U18", "U19"}, false},
		{"a", "_", 4, []string{"b0", "b1", "b2", "b3"}, true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: (%q..%q)", i, tc.prev, tc.next), func(t *testing.T) {
			got, err := RankN(tc.prev, tc.next, tc.n)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Errorf("unexpected error: %v", err)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %#v, got %#v", tc.want, got)
			}
		})
	}
}
