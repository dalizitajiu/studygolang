package Entity

type SchemaTerm struct {
	Name string
	Type string
}
func NewSchemaTerm(name string,schematype string)SchemaTerm{
	return SchemaTerm{name,schematype}
}

type Table struct {
	Id int64
	Prn string
	Schemas []SchemaTerm
	Url string
}

