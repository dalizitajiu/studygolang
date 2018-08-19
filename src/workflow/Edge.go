package workflow

type EdgeFuntion func(map[string]interface{})map[string]interface{}

const(
	FunctionType = 1
	OperatorType = 2

)
var RunTypeMap = map[int]string{
	FunctionType:"函数形",
	OperatorType:"算子形",
}
const(
	Runing = 1
	Stoping = 0
	Finished =2

)

type Edge struct {
	StartVertex *Vertex
	EndVertex *Vertex
	EdgeType int
	IsDone bool
	Runable RunInterface
	Status int
}

func (edge *Edge)GetEdgeType()string{
	return RunTypeMap[edge.EdgeType]
}

func NewFunctionEdge(fun EdgeFuntion)*Edge{
	res:=&Edge{}
	res.EdgeType = FunctionType
	res.IsDone = false
	res.Runable = NewFunctionRunEntity(fun)
	res.Status = Stoping
	return res
}

func NewOperatorEdge(fun string)*Edge{
	res:=&Edge{}
	res.EdgeType = OperatorType
	res.IsDone = false
	res.Runable = NewOperatorRunEntity(fun)
	res.Status = Stoping
	return res
}

/*
成功跑完后
 */

//func (edge *Edge)DagNext(){
//	edge.EndVertex.AddInputSlot(edge.StartVertex.GetOutputSlot().ToInputSlot())
//	edge.EndVertex.Supply()
//	edge.IsDone = true
//}