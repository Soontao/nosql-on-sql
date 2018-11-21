package nosql

import (
	"fmt"
	"strings"
)

// GenerateDDLFromDefination used for generate Table Create DDL
func GenerateDDLFromDefination(table Table) string {
	columns := []string{}
	for _, column := range table.columns {
		columns = append(columns, fmt.Sprintf("%s %s(%d)", column.columnName, column.columnType, column.columnLength))
	}
	rt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", table.tableName, strings.Join(columns, ","))
	return rt
}
