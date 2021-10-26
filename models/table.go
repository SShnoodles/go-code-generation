package models

type Schema struct {
	Name    string
	Comment string
	Tables  []Table
}

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

type Column struct {
	Name          string
	Type          string
	Comment       string
	PrimaryKey    bool
	Nullable      bool
	Length        int
	DecimalLength int
}
