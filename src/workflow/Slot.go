package workflow



type BaseSlot struct {
	Data map[string]interface{}
}

func (base *BaseSlot)SetData(data map[string]interface{}){
	base.Data = data
}
func (base *BaseSlot)GetData()map[string]interface{}{
	return base.Data
}
func NewSlot()*BaseSlot{
	res:=&BaseSlot{Data:map[string]interface{}{}}
	return res
	}

type InputSlot struct {
	*BaseSlot
}

func NewInputSlot(data map[string]interface{})*InputSlot{
	temp:=NewSlot()
	temp.SetData(data)
	res:=&InputSlot{BaseSlot:temp}
	return res
}

type OutSlot struct {
	*BaseSlot
}

func NewOutSlot(data map[string]interface{})*OutSlot{
	temp:=NewSlot()
	temp.SetData(data)
	res:=&OutSlot{BaseSlot:temp}
	return res
}

func (out *OutSlot)ToInputSlot()*InputSlot{

	res := &InputSlot{BaseSlot:NewSlot()}
	res.SetData(out.GetData())
	return res
}