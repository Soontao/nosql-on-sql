package nosql

import (
	"testing"
)

func TestGenerateDDLFromDefination(t *testing.T) {
	type args struct {
		table Table
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple ddl generate test",
			args{Table{tableName: "User", columns: map[string]Column{"ObjectID": Column{columnName: "ObjectID", columnType: "VARCHAR", columnLength: 32}}}},
			"CREATE TABLE IF NOT EXISTS User (ObjectID VARCHAR(32));",
		},
		{
			"complex ddl generate test",
			args{
				Table{
					"User4",
					map[string]Column{
						"ObjectID": Column{columnName: "ObjectID", columnType: "VARCHAR", columnLength: 32},
						"UserName": Column{columnName: "UserName", columnType: "VARCHAR", columnLength: 32},
						"Address":  Column{columnName: "Address", columnType: "VARCHAR", columnLength: 128},
						"Age":      Column{columnName: "Age", columnType: "INTEGER", columnLength: 8},
					},
				},
			},
			"CREATE TABLE IF NOT EXISTS User4 (ObjectID VARCHAR(32),UserName VARCHAR(32),Address VARCHAR(128),Age INTEGER(8));",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateDDLFromDefination(tt.args.table); got != tt.want {
				t.Errorf("GenerateDDLFromDefination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateAlterModifyDDL(t *testing.T) {
	type args struct {
		table Table
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple test",
			args{Table{"User3", map[string]Column{
				"UserName": Column{"UserName", "VARCHAR", 32},
				"Age":      Column{"Age", "INTEGER", 2},
			}}},
			"ALTER TABLE User3 MODIFY UserName VARCHAR(32);ALTER TABLE User3 MODIFY Age INTEGER(2);",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateAlterModifyDDL(tt.args.table); got != tt.want {
				t.Errorf("GenerateAlterModifyDDL() = %v, want %v", got, tt.want)
			}
		})
	}
}
