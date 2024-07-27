package query

import (
	"testing"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		wantQ string
		wantV []interface{}
	}{
		{
			name:  "Empty",
			items: []Item{},
			wantQ: "",
			wantV: nil,
		},
		{
			name: "Single Item",
			items: []Item{
				{
					Key:    "name",
					Values: []interface{}{"John"},
					Skip:   false,
				},
			},
			wantQ: "name",
			wantV: []interface{}{"John"},
		},
		{
			name: "Multiple Items",
			items: []Item{
				{
					Key:    "name",
					Values: []interface{}{"John"},
					Skip:   false,
				},
				{
					Key:    "age",
					Values: []interface{}{30},
					Skip:   false,
				},
				{
					Key:    "city",
					Values: []interface{}{"New York"},
					Skip:   false,
				},
			},
			wantQ: "name AND age AND city",
			wantV: []interface{}{"John", 30, "New York"},
		},
		{
			name: "Skip Item",
			items: []Item{
				{
					Key:    "name",
					Values: []interface{}{"John"},
					Skip:   true,
				},
				{
					Key:    "age",
					Values: []interface{}{30},
					Skip:   false,
				},
				{
					Key:    "city",
					Values: []interface{}{"New York"},
					Skip:   false,
				},
			},
			wantQ: "age AND city",
			wantV: []interface{}{30, "New York"},
		},
		{
			name: "Empty Value",
			items: []Item{
				{
					Key:    "name",
					Values: []interface{}{""},
					Skip:   false,
				},
				{
					Key:    "age",
					Values: []interface{}{30},
					Skip:   false,
				},
				{
					Key:    "city",
					Values: []interface{}{"New York"},
					Skip:   false,
				},
			},
			wantQ: "age AND city",
			wantV: []interface{}{30, "New York"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQ, gotV := Build(tt.items)
			if gotQ != tt.wantQ || !sliceEqual(gotV, tt.wantV) {
				t.Errorf("Build() = %q, %v, want %q, %v", gotQ, gotV, tt.wantQ, tt.wantV)
			}
		})
	}
}

func sliceEqual(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
