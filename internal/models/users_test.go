package models

import (
	"testing"

	"snippetbox.sinantalebi.net/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	// Set up a suite of table-driven tests and expected results.
	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-existent ID",
			userID: 2,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get connection pool to the test database.
			// It will be set up and torn down for each sub-test.
			db := NewTestDB(t)
			m := UserModel{db}

			exists, err := m.Exists(tt.userID)
			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}
