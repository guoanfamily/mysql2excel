package service

type TableModel struct {
	TableName     string `db:"TableName"`
	TableComment  string `db:"TableComment"`
	ColumnName    string `db:"ColumnName"`
	ColumnComment string `db:"ColumnComment"`
	ColumnType    string `db:"ColumnType"`
	IsPrimary     string `db:"IsPrimary"`
	CanNull       string `db:"CanNull"`
}
