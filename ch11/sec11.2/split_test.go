package sec11_2

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		separator string
		wantLen   int
	}{
		{
			name:      "a:b:c + : => 3",
			str:       "a:b:c",
			separator: ":",
			wantLen:   3,
		},
		{
			name:      "a=b=c + = => 3",
			str:       "a=b=c",
			separator: "=",
			wantLen:   3,
		},
		{
			name:      "a:b:c + = => 1",
			str:       "a:b:c",
			separator: "=",
			wantLen:   1,
		},
		{
			name:      "a:b:c + b => 2",
			str:       "a:b:c",
			separator: "b",
			wantLen:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			words := strings.Split(tt.str, tt.separator)
			gotLen := len(words)
			if gotLen != tt.wantLen {
				t.Errorf("Split(%s, %s) returned %d words, want: %d", tt.str, tt.separator, gotLen, tt.wantLen)
			}
		})
	}
}
