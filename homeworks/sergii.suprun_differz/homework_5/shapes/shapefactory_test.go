package shapes

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		conf map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    Shaper
		wantErr bool
	}{
		{name: "cone",
			args:    args{conf: map[string]string{"SHAPE": "cone", "R": "5", "H": "10"}},
			want:    &Cone{s: "Cone", m: 261, r: 5, h: 10},
			wantErr: false},
		{name: "cube",
			args:    args{conf: map[string]string{"SHAPE": "cube", "A": "10"}},
			want:    &Cube{s: "Cube", m: 1000, a: 10},
			wantErr: false},
		{name: "cylinder",
			args:    args{conf: map[string]string{"SHAPE": "cylinder", "R": "5", "H": "10"}},
			want:    &Cylinder{s: "Cylinder", m: 785, r: 5, h: 10},
			wantErr: false},
		{name: "prizm",
			args:    args{conf: map[string]string{"SHAPE": "prizm", "A": "10", "H": "10"}},
			want:    &Prizm{s: "Prizm", m: 433, a: 10, h: 10},
			wantErr: false},
		{name: "pyramid",
			args:    args{conf: map[string]string{"SHAPE": "pyramid", "A": "10", "H": "10"}},
			want:    &Pyramid{s: "Pyramid", m: 333, a: 10, h: 10},
			wantErr: false},
		{name: "sphere",
			args:    args{conf: map[string]string{"SHAPE": "sphere", "R": "5"}},
			want:    &Sphere{s: "Sphere", m: 523, r: 5},
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
