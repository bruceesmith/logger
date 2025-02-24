// Code generated by "stringer -type LogID"; DO NOT EDIT.

package logger

import "testing"

func TestLogID_String(t *testing.T) {
	tests := []struct {
		name string
		i    LogID
		want string
	}{
		{
			name: "norm",
			i:    Norm,
			want: "Norm",
		},
		{
			name: "tracy",
			i:    Tracy,
			want: "Tracy",
		},
		{
			name: "whatthe",
			i:    77,
			want: "LogID(77)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("LogID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
