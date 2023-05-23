package parse

/*
Copyright 2023 UMH Systems GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
