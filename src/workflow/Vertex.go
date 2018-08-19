package workflow

import "log"

/*
Slot只包含
*/
const(
	Failed  = 0
	Success = 1
)

type Vertex struct {
	InputSlots []*InputSlot
	OutputSlot *OutSlot
	Status int
	RequiredInputSlotNum int
}

func NewVertex()*Vertex{
	res:=&Vertex{}
	res.InputSlots = make([]*InputSlot,0)
	res.Status = Failed
	return res
	}

func (vertex *Vertex)SetOutputSlot(slot *OutSlot){
	vertex.OutputSlot = slot
}

/*
获取合并后的input
 */
func (vertex *Vertex)GetMergedInput()map[string]interface{}{
	res:=map[string]interface{}{}
	for _,v:=range vertex.InputSlots{
		for k2,v2:= range v.GetData(){
			res[k2] = v2
		}
	}
	return res
}

func (vertex *Vertex)GetOutputSlot()*OutSlot{
	return vertex.OutputSlot
}

func (vertex *Vertex)AddInputSlot(slot *InputSlot){
	vertex.InputSlots = append(vertex.InputSlots,slot)
}

/*
end vertext调用,所需要的slotnum减一，可用的inputslot++
 */
func (vertex *Vertex)Supply(input *InputSlot){
	log.Println(input.GetData())
	vertex.RequiredInputSlotNum--
	vertex.AddInputSlot(input)
}





