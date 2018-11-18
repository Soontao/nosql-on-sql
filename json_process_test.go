package nosql

import (
	"reflect"
	"testing"
)

func TestParseJSON(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// test cases
		{
			"simple object",
			args{
				inputString: `{"hello":"world"}`,
			},
			JSONObject{"hello": "world"},
		},
		{
			"object with list",
			args{
				inputString: `{"hello":"world", "hello2": []}`,
			},
			JSONObject{"hello": "world", "hello2": JSONArray{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseJSON(tt.args.inputString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
