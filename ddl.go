package nosql

import (
	"fmt"
	"strings"
)

// DDLType enum
type DDLType int

const (
	// Create new table
	Create DDLType = iota
	// Add new columns
	Add
	// Modify existed columns type
	Modify
	// Remove existed columns
	Remove
)

// GenerateDDL from table defination
func GenerateDDL(table Table, ddlType DDLType) string {
	switch ddlType {
	case Create:
		return GenerateDDLFromDefination(table)
	case Modify:
		return GenerateAlterModifyDDL(table)
	default:
		return ""
	}
}

// GenerateDDLFromDefination used for generate Table Create DDL
func GenerateDDLFromDefination(table Table) string {
	columns := []string{}
	for _, column := range table.columns {
		columns = append(columns, fmt.Sprintf("%s %s(%d)", column.columnName, column.columnType, column.columnLength))
	}
	rt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", table.tableName, strings.Join(columns, ","))
	return rt
}

// GenerateAlterModifyDDL used for modify existed column type
func GenerateAlterModifyDDL(table Table) string {
	columns := []string{}
	for _, column := range table.columns {
		columns = append(columns, fmt.Sprintf("ALTER TABLE %s MODIFY %s %s(%d);", table.tableName, column.columnName, column.columnType, column.columnLength))
	}
	rt := strings.Join(columns, "")
	return rt
}
