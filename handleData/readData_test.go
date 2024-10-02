package linears

import (
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name:    "Valid data file",
			args:    args{path: "text.txt"},
			want:    []float64{1.0, 2.5, 3.0, 4.5, 5.0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadData(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexData(t *testing.T) {
	type args struct {
		data  []float64
		value float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Value does not exist",
			args: args{data: []float64{1.0, 2.5, 3.0, 4.5, 5.0}, value: 6.0},
			want: -1,
		},
		{
			name: "Duplicate values",
			args: args{data: []float64{1.0, 2.5, 3.0, 2.5, 5.0}, value: 2.5},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexData(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("IndexData() = %v, want %v", got, tt.want)
			}
		})
	}
}
