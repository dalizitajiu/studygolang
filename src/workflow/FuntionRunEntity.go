package workflow

import "log"

type FunctionRunEntity struct {
	myfunc EdgeFuntion
	Name string
}
func NewFunctionRunEntity(fun EdgeFuntion)*FunctionRunEntity{
	res:=&FunctionRunEntity{}
	res.myfunc = fun
	return res
}
func (fun *FunctionRunEntity)Run(slot Slot)*OutSlot{
	log.Println("[function is running]")
	return NewOutSlot(fun.myfunc(slot.GetData()))
}

func (fun *FunctionRunEntity)GetName()string{
	return fun.Name
}
func (fun *FunctionRunEntity)SetName(name string){
	fun.Name = name
}