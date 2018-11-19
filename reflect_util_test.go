package nosql

import (
	"reflect"
	"testing"
)

func TestInspectObject(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name string
		args args
		want ObjectInfo
	}{
		{
			"is Array",
			args{
				any: []int{1},
			},
			ObjectInfo{isArray: true, value: []int{1}},
		},
		{
			"is Object",
			args{ObjectInfo{}},
			ObjectInfo{isObject: true, value: ObjectInfo{}},
		},
		{
			"is Object",
			args{map[string]string{}},
			ObjectInfo{isObject: true, value: map[string]string{}},
		},
		{
			"is String",
			args{"hello world"},
			ObjectInfo{isString: true, value: "hello world"},
		},
		{
			"is Number",
			args{1},
			ObjectInfo{isNumber: true, value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InspectObject(tt.args.any); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InspectObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectInfo_Value(t *testing.T) {
	type fields struct {
		isArray  bool
		isObject bool
		isString bool
		isNumber bool
		value    interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"Info.Value", fields{isString: true, value: "1"}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ObjectInfo{
				isArray:  tt.fields.isArray,
				isObject: tt.fields.isObject,
				isString: tt.fields.isString,
				isNumber: tt.fields.isNumber,
				value:    tt.fields.value,
			}
			if got := o.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObjectInfo.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectInfo_GetInnerInfo(t *testing.T) {
	type fields struct {
		isArray  bool
		isObject bool
		isString bool
		isNumber bool
		value    interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]ObjectInfo
	}{
		{
			"Info.GetInnerInfo for map",
			fields{isString: true, value: map[string]interface{}{"hello": "world"}},
			map[string]ObjectInfo{"hello": ObjectInfo{isString: true, value: "world"}},
		},
		{
			"Info.GetInnerInfo for slice",
			fields{isString: true, value: []string{"hello", "world"}},
			map[string]ObjectInfo{
				"0": ObjectInfo{isString: true, value: "hello"},
				"1": ObjectInfo{isString: true, value: "world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ObjectInfo{
				isArray:  tt.fields.isArray,
				isObject: tt.fields.isObject,
				isString: tt.fields.isString,
				isNumber: tt.fields.isNumber,
				value:    tt.fields.value,
			}
			if got := o.GetInnerInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObjectInfo.GetInnerInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
