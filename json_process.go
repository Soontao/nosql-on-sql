package nosql

import (
	"encoding/json"
)

// JSONObject type
type JSONObject = map[string]interface{}

// JSONArray type
type JSONArray = []interface{}

// ParseJSON Object
// must input a json object string like {}
func ParseJSON(inputString string) JSONObject {
	rt := JSONObject{}
	err := json.Unmarshal([]byte(inputString), &rt)
	if err != nil {
		panic(err)
	}
	return rt
}

// ParseJSONArray object
// must be a json array string like []
func ParseJSONArray(inputString string) JSONArray {
	rt := JSONArray{}
	err := json.Unmarshal([]byte(inputString), &rt)
	if err != nil {
		panic(err)
	}
	return rt
}
