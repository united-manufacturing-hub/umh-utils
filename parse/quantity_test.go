package parse

import "testing"

func TestParseQuantity(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"empty", args{""}, 0, true},
		{"zero", args{"0"}, 0, false},
		{"zeroes", args{"000"}, 0, false},
		{"one", args{"1"}, 1, false},
		{"negative one", args{"-1"}, -1, false},
		{"decimal", args{"1.1"}, 0, true},
		{"decimal suffix milli", args{"1.1m"}, 0, true},
		{"decimal suffix million", args{"1.1M"}, 1100000, false},
		{"decimal suffix million zero", args{"0.1M"}, 100000, false},
		{"binary", args{"10Mi"}, 10485760, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Quantity(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Quantity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("%s: Quantity() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
