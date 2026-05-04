package gcfg

import (
	"strings"
	"testing"
)

func TestMisspelledSecionNameDoesNotProduceDuplicateErrorMessages(t *testing.T) {
	var cfg struct {
		Profile map[string]*struct {
			Color string
		}
	}

	tests := []struct {
		name  string
		input string
	}{
		{
			name: "misspelled subsection style section",
			input: `
[profileX "A"]
color = white`,
		},
		{
			name: "misspelled struct style section",
			input: `
[profileX]
color = white`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ReadStringInto(&cfg, tt.input)
			if err == nil {
				t.Fatal("expected error for unknown section name, got nil")
			}
			errStr := err.Error()

			// Each distinct error message should only appear once.
			if n := strings.Count(errStr, `can't store data at section "profileX"`); n != 1 {
				t.Errorf(`"can't store data at section" appeared %d time(s), want 1 — duplicate errors indicate issue
 #14 regression:\n%s`, n, errStr)
			}
			if n := strings.Count(errStr, `can't store data into key "color" at section "profileX"`); n != 1 {
				t.Errorf(`"can't store data into key" appeared %d time(s), want 1 — duplicate errors indicate issue #14
 regression:\n%s`, n, errStr)
			}
		})
	}
}
