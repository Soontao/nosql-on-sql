package nosql

import (
	"reflect"
	"strconv"
)

// ObjectInfo struct
type ObjectInfo struct {
	isArray  bool
	isObject bool
	isString bool
	isNumber bool
	value    interface{}
}

// NumberInfo struct
type NumberInfo struct {
	isFloat   bool
	isInteger bool
	length    int
}

// Diff Object Type
// return diff information
func (o *ObjectInfo) Diff(before ObjectInfo, after ObjectInfo) []ObjectInfo {
	rt := []ObjectInfo{}
	if before.isObject && after.isObject {
		return rt
	}
	return rt
}

// GetInnerInfo for object or array object
func (o *ObjectInfo) GetInnerInfo() map[string]ObjectInfo {
	rt := map[string]ObjectInfo{}
	objectType := reflect.TypeOf(o.value)
	objectValue := reflect.ValueOf(o.value)

	switch objectType.Kind() {
	case reflect.Slice:
		for i := 0; i < objectValue.Len(); i++ {
			rt[strconv.Itoa(i)] = InspectObject(objectValue.Index(i).Interface())
		}
	case reflect.Map:
		for f, v := range o.value.(map[string]interface{}) {
			rt[f] = InspectObject(v)
		}
	case reflect.Struct:
		for index := 0; index < objectType.NumField(); index++ {
			f := objectType.Field(index)
			v := objectValue.Field(index)
			rt[f.Name] = InspectObject(v.Interface())
		}
	}

	return rt
}

// Value info
func (o *ObjectInfo) Value() interface{} {
	return o.value
}

func (o *ObjectInfo) IsArray() bool {
	return o.isArray
}

func (o *ObjectInfo) IsString() bool {
	return o.isString
}

func (o *ObjectInfo) IsObject() bool {
	return o.isObject
}

func (o *ObjectInfo) IsNumber() bool {
	return o.isNumber
}

// InspectObject information
func InspectObject(any interface{}) ObjectInfo {
	rt := ObjectInfo{false, false, false, false, any}
	kind := reflect.TypeOf(any).Kind()
	switch kind {
	case reflect.Struct, reflect.Map:
		rt.isObject = true
	case reflect.Array, reflect.Slice:
		rt.isArray = true
	case reflect.String:
		rt.isString = true
	case
		reflect.Uint,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		rt.isNumber = true
	default:
		panic("not support kind:" + kind.String())
	}
	return rt
}
