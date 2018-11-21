package nosql

type Database struct {
	databaseName string
	tables       map[string]Table
}

type Table struct {
	tableName string
	columns   map[string]Column
}

type Column struct {
	columnName   string
	columnType   string
	columnLength int32
}
