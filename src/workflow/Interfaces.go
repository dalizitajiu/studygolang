package workflow
type RunInterface interface {
	Run(Slot)*OutSlot
	GetName()string
}
type Slot interface {
	SetData(map[string]interface{})
	GetData()map[string]interface{}
}