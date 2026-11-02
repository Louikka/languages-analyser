package main

import "testing"

func TestIsSubpath(t *testing.T) {
	tests := []struct {
		name     string
		parent   string
		child    string
		expected bool
	}{
		{
			parent:   "/some/absolute/path",
			child:    "/some/absolute/path/child",
			expected: true,
		},
		{
			parent:   "./some/relative/path",
			child:    "./some/relative/path/child",
			expected: true,
		},
		{
			parent:   "/some/path",
			child:    "/another/path",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is, err := isSubpath(tt.parent, tt.child)
			if err != nil {
				t.Errorf("error: %s", err)
			}
			if is != tt.expected {
				t.Errorf("expected %t, but got %t", tt.expected, is)
			}
		})
	}
}
