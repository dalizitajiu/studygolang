package main
import  (
	"../workflow"
	"log"
)

func Double(input map[string]interface{})map[string]interface{}{
	raw:=input["line"].([]float64)
	temp:=make([]float64,len(raw))
	for k,v:=range raw{
		temp[k]=2*v
	}
	res:=map[string]interface{}{}
	res["line"] = temp
	return res
}

func main(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	dag:=workflow.NewDag()

	vert1:= workflow.NewVertex()

	vert1.RequiredInputSlotNum=1
	data:=map[string]interface{}{}
	data["line"] =[]float64{4.5,6.7}
	//log.Println(vert1.OutputSlot.GetData())
	vert2:= workflow.NewVertex()

	edge1:=workflow.NewFunctionEdge(Double)
	dag.Connect(vert1,vert2,edge1)
	log.Println(edge1.Status)

	vert3:=workflow.NewVertex()
	edge2:=workflow.NewOperatorEdge("C:\\Users\\Administrator\\Desktop\\gobigdata\\src\\TestOperator\\Funny.exe")
	dag.Connect(vert2,vert3,edge2)

	vert1.Supply(workflow.NewInputSlot(data))
	log.Println("开始前的mergedmap",vert1.GetMergedInput())
	for _,v:=range dag.Edges{
		log.Println(v.GetEdgeType())
	}
	dag.Run()



	//log.Println(vert3.GetMergedInput())
	}
