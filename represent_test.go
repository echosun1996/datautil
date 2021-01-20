package datautil

import "testing"

var testRecords = [][]string{
	{"first_name", "last_name", "username"},
	{"Rob", "Pike", "rob"},
	{"Ken", "Thompson", "ken"},
	{"Robert", "Griesemer", "gri"},
}

func TestDataRepresentation(t *testing.T) {
	type args struct {
		data [][]string
		top  int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{testRecords, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DataRepresentation(tt.args.data, tt.args.top)
		})
	}
}
