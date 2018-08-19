package workflow

import (
	"os/exec"
	"fmt"
	"encoding/json"
	"log"
)

type OperatorRunEntity struct {
	RunPath string
	Name string
}
func NewOperatorRunEntity(runpath string)*OperatorRunEntity{
	res:=&OperatorRunEntity{RunPath:runpath}
	return res
}

func (this *OperatorRunEntity)GetName()string{
	return this.Name
}
func (this *OperatorRunEntity)SetName(name string){
	this.Name = name
}

func map2arglist(argmap map[string]interface{})[]string{
	arglist := make([]string,0)
	for k,v:=range argmap{
		arglist = append(arglist,"--"+k)
		arglist = append(arglist,fmt.Sprintf("%v",v))
	}
	return arglist
}

func (this *OperatorRunEntity)Run(slot Slot)*OutSlot{
	log.Println("[operator is running]")
	argmap := slot.GetData()
	arglist := map2arglist(argmap)
	log.Println("[arglist]",arglist)
	cmd:=exec.Cmd{Path:this.RunPath,Args:arglist}
	out,err:=cmd.Output()
	log.Println(string(out))
	checkError(err)
	var outmap = map[string]interface{}{}
	json.Unmarshal(out,&outmap)
	log.Println("[realout]",outmap,outmap["line"])
	return NewOutSlot(outmap)
}
