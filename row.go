package nosql

// Row type
type Row = map[string]interface{}

// ObjectID type
type ObjectID = string

// NewRow instance
func NewRow() *Row {
	return &Row{"ObjectID": nil}
}
