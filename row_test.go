package nosql

import (
	"reflect"
	"testing"
)

func TestNewRow(t *testing.T) {
	tests := []struct {
		name string
		want *Row
	}{
		{
			"create row instance",
			&map[string]interface{}{"ObjectID": nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRow(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
