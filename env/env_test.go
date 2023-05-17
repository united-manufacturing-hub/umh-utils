package env

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

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func setup() {
	err := os.Setenv("EMPTY_VAR", "")
	if err != nil {
		panic(err)
	}

	fmt.Println("Set EMPTY_VAR to \"\"")
}

func teardown() {
	err := os.Unsetenv("EMPTY_VAR")
	if err != nil {
		panic(err)
	}

	fmt.Println("Unset EMPTY_VAR")
}

func setupEnv(correctValue, wrongValue string) {
	err := os.Setenv("EXISTING_VAR", correctValue)
	if err != nil {
		panic(err)
	}
	err = os.Setenv("WRONG_VAR", wrongValue)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Set EXISTING_VAR to %v\n", correctValue)
}

func teardownEnv() {
	// Unset environment variables
	err := os.Unsetenv("EXISTING_VAR")
	if err != nil {
		panic(err)
	}

	fmt.Println("Unset EXISTING_VAR")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// TestGetAsString tests the GetAsString function
func TestGetAsString(t *testing.T) {
	expected := "my-value"
	wrong := 123
	fallback := "fallback-value"
	setupEnv(expected, strconv.Itoa(wrong))
	defer teardownEnv()

	type args struct {
		key      string
		required bool
		fallback string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Case 1: Variable exists, return value",
			args: args{
				key:      "EXISTING_VAR",
				required: false,
				fallback: fallback,
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "Case 2: Variable is empty, return empty value",
			args: args{
				key:      "EMPTY_VAR",
				required: false,
				fallback: fallback,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Case 3: Variable does not exist and is not required, return fallback value",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: false,
		},
		{
			name: "Case 4: Variable does not exist and is required, return error",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: true,
				fallback: fallback,
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAsString(tt.args.key, tt.args.required, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetAsInt tests the GetAsInt function
func TestGetAsInt(t *testing.T) {
	expected := 123
	wrong := "wrong"
	fallback := 456
	setupEnv(strconv.Itoa(expected), wrong)
	defer teardownEnv()

	type args struct {
		key      string
		required bool
		fallback int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Case 1: Variable exists, return value",
			args: args{
				key:      "EXISTING_VAR",
				required: false,
				fallback: fallback,
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "Case 2: Variable is empty, return error",
			args: args{
				key:      "EMPTY_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: true,
		},
		{
			name: "Case 3: Variable does not exist and is not required, return fallback value",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: false,
		},
		{
			name: "Case 4: Variable does not exist and is required, return error",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: true,
				fallback: fallback,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Case 5: Variable is not an integer, return fallback value and error",
			args: args{
				key:      "WRONG_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAsInt(tt.args.key, tt.args.required, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsInt() error = %v, wantErr %v. got: %v", err, tt.wantErr, got)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetAsUint64 tests the GetAsUint64 function
func TestGetAsUint64(t *testing.T) {
	expected := uint64(123)
	wrong := "wrong"
	fallback := uint64(456)
	setupEnv(strconv.FormatUint(expected, 10), wrong)
	defer teardownEnv()

	type args struct {
		key      string
		required bool
		fallback uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Case 1: Variable exists, return value",
			args: args{
				key:      "EXISTING_VAR",
				required: false,
				fallback: fallback,
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "Case 2: Variable is empty, return error",
			args: args{
				key:      "EMPTY_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: true,
		},
		{
			name: "Case 3: Variable does not exist and is not required, return fallback value",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: false,
		},
		{
			name: "Case 4: Variable does not exist and is required, return error",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: true,
				fallback: fallback,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Case 5: Variable is not an integer, return fallback value and error",
			args: args{
				key:      "WRONG_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAsUint64(tt.args.key, tt.args.required, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetAsFloat64 tests the GetAsFloat64 function
func TestGetAsFloat64(t *testing.T) {
	expected := 123.456
	wrong := "wrong"
	fallback := 456.789
	setupEnv(strconv.FormatFloat(expected, 'f', -1, 64), wrong)
	defer teardownEnv()

	type args struct {
		key      string
		required bool
		fallback float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Case 1: Variable exists, return value",
			args: args{
				key:      "EXISTING_VAR",
				required: false,
				fallback: fallback,
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "Case 2: Variable is empty, return error",
			args: args{
				key:      "EMPTY_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: true,
		},
		{
			name: "Case 3: Variable does not exist and is not required, return fallback value",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: false,
		},
		{
			name: "Case 4: Variable does not exist and is required, return error",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: true,
				fallback: fallback,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Case 5: Variable is not a float, return fallback value and error",
			args: args{
				key:      "WRONG_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAsFloat64(tt.args.key, tt.args.required, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetAsBool tests the GetAsBool function
func TestGetAsBool(t *testing.T) {
	setupEnv(strconv.FormatBool(true), "wrong")
	defer teardownEnv()

	type args struct {
		key      string
		required bool
		fallback bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Case 1: Variable exists, return value",
			args: args{
				key:      "EXISTING_VAR",
				required: false,
				fallback: false,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Case 2: Variable is empty, return error",
			args: args{
				key:      "EMPTY_VAR",
				required: false,
				fallback: false,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Case 3: Variable does not exist and is not required, return fallback value",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: false,
				fallback: true,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Case 4: Variable does not exist and is required, return error",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: true,
				fallback: false,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Case 5: Variable is not a bool, return fallback value and error",
			args: args{
				key:      "WRONG_VAR",
				required: false,
				fallback: false,
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAsBool(tt.args.key, tt.args.required, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetAsType tests the GetAsType function
func TestGetAsType(t *testing.T) {
	type testStruct struct {
		BoolValue   bool   `json:"boolValue"`
		IntValue    int    `json:"intValue"`
		StringValue string `json:"stringValue"`
	}
	expected := testStruct{
		BoolValue:   true,
		IntValue:    123,
		StringValue: "test",
	}
	wrong := struct {
		WrongValue float64 `json:"wrongValue"`
	}{
		WrongValue: 1.23,
	}
	fallback := testStruct{
		BoolValue:   false,
		IntValue:    456,
		StringValue: "fallback",
	}
	strEnv, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("Error marshalling testStruct: %v", err)
	}
	strWrong, err := json.Marshal(wrong)
	if err != nil {
		t.Errorf("Error marshalling wrong: %v", err)
	}
	setupEnv(string(strEnv), string(strWrong))
	defer teardownEnv()

	type args struct {
		key      string
		required bool
		fallback testStruct
	}
	tests := []struct {
		name    string
		args    args
		want    testStruct
		wantErr bool
	}{
		{
			name: "Case 1: Variable exists, return value",
			args: args{
				key:      "EXISTING_VAR",
				required: false,
				fallback: fallback,
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "Case 2: Variable is empty, return error",
			args: args{
				key:      "EMPTY_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: true,
		},
		{
			name: "Case 3: Variable does not exist and is not required, return fallback value",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: false,
				fallback: fallback,
			},
			want:    fallback,
			wantErr: false,
		},
		{
			name: "Case 4: Variable does not exist and is required, return error",
			args: args{
				key:      "NONEXISTENT_VAR",
				required: true,
				fallback: fallback,
			},
			want:    testStruct{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got testStruct
			err := GetAsType(tt.args.key, &got, tt.args.required, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAsType() = %v, want %v", got, tt.want)
			}
		})
	}
}
