package workflow

import (
	"log"
	"time"
)

type Dag struct {
	Edges   []*Edge
	TaskNum int
}

func NewDag()*Dag{
	res:=&Dag{}
	res.Edges = make([]*Edge,0)
	res.TaskNum = 0

	return res
}

func (dag *Dag)Connect(start *Vertex,end *Vertex,edge *Edge){
	edge.StartVertex = start
	edge.EndVertex = end
	dag.TaskNum++
	end.RequiredInputSlotNum++
	dag.Edges = append(dag.Edges,edge)
}

func (dag *Dag)Run(){
	log.Println("[dag.tasknum]",dag.TaskNum)
	log.Println("[num of edges]",len(dag.Edges))
	for{
		if dag.TaskNum <1{
			break
		}
		for k,v:=range dag.Edges{
			tk:=k
			tv:=v
			log.Println("[$$$]",tk,"[required num]",tv.StartVertex.RequiredInputSlotNum,tv.Status,tv.EndVertex.RequiredInputSlotNum)

			//如果开始的条件满足,满足的节点合并inputslot,然后开始跑
			//log.Println("[边index]",k,"[inputslots]",v.StartVertex.GetMergedInput(),v.StartVertex.RequiredInputSlotNum)
			if tv.StartVertex.RequiredInputSlotNum == 0 && tv.Status == Stoping{
				tv.Status = Runing
				log.Println("---",tk)
				go func(current int) {
					mergedmap:=tv.StartVertex.GetMergedInput()//获取开始节点的合并map
					res:=tv.Runable.Run(NewInputSlot(mergedmap))
					tv.EndVertex.Supply(res.ToInputSlot())
					log.Println("[endslot required slot num]",tv.EndVertex.RequiredInputSlotNum)
					dag.TaskNum--
					tv.Status = Finished
					log.Println("[dag.TaskNum]",dag.TaskNum)
				}(tk)
			}
		}
		time.Sleep(time.Second)
	}

}
