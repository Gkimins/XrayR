package sspanel

import "testing"

func TestVersionOlderThan(t *testing.T) {
	cases := []struct {
		version   string
		threshold string
		want      bool
	}{
		// New calendar-semver scheme must never be treated as older.
		{"25.1.0", "2021.11", false},
		{"25.1.0", "2023.2", false},
		{"24.12.31", "2023.2", false},
		// Old date-based scheme keeps working.
		{"2021.11", "2021.11", false},
		{"2020.5", "2021.11", true},
		{"2023.1", "2023.2", true},
		{"2023.2", "2023.2", false},
		// Empty / unknown version (legacy panels) -> treated as older.
		{"", "2021.11", true},
	}
	for _, tc := range cases {
		c := &APIClient{version: tc.version}
		if got := c.versionOlderThan(tc.threshold); got != tc.want {
			t.Errorf("versionOlderThan(version=%q, threshold=%q) = %v, want %v",
				tc.version, tc.threshold, got, tc.want)
		}
	}
}
