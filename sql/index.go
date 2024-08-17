package sql

type IndexSchema struct {
	Column *ColumnSchema
	Unique bool
}

func (this IndexSchema) Generate() string {
	var sb stringBuiler

	sb.Writef("CREATE")

	if this.Unique || this.Column.UniqueIndexed {
		sb.Writef(" UNIQUE")
	}

	sb.Writef(" INDEX idx_%s ON %s (%s);",
		this.Column.Table.Name,
		this.Column.Table.Name,
		this.Column.Name,
	)

	return sb.String()
}
