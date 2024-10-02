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
