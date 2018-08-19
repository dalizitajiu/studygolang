package matrix

import (
	"log"
	"reflect"
	"container/list"
	"fmt"
	"math"
)

type Vec interface {
	GetSize()int
	SetSize(int)
	Set(int,float64)
	Get(int)float64
	Print()
	GetVecType()int
	SetVecType(int)
	GetData()[]float64
	SetData([]float64)
	Avg()float64
	Sum()float64
	Std()float64
	Normal()
	Max()float64
	Min()float64
	StdNoDiv()float64
}

type RawVec struct {
	Data    []float64
	Size    int
	VecType int //向量类型
}
func (vec *RawVec)GetSize()int{
	return vec.Size;
}
func (vec *RawVec)Avg()float64{
	return vec.Sum()/float64(vec.GetSize());
}
func (vec *RawVec)Max()float64{
	res:=vec.Get(0)
	for _,v:=range vec.Data{
		if v>res{
			res = v
		}
	}
	return res
}
func (vec *RawVec)Min()float64{
	res:=vec.Get(0)
	for _,v:=range vec.Data{
		if v<res{
			res = v
		}
	}
	return res
}
//归一化处理Xi=(Xi-min)/(max-min)
func (vec *RawVec)Normal(){
	max:=vec.Max()
	min:=vec.Min()
	delta:=max-min
	if delta<1e-6{
		return
	}
	for k,v:=range vec.GetData(){
		vec.Set(k,(v-min)/delta)
	}
}
//标准差sqrt(sum(Xi-avg(X))**2/(n-1))
func (vec *RawVec)Std()float64{
	var avg = vec.Avg()
	var res float64 = 0

	for _,v:=range vec.GetData(){
		res+=math.Pow(v-avg,2)
	}
	return math.Sqrt(res/float64(vec.GetSize()-1))
}

func (vec *RawVec)StdNoDiv()float64{
	var avg = vec.Avg()
	var res float64 = 0

	for _,v:=range vec.GetData(){
		res+=math.Pow(v-avg,2)
	}
	return math.Sqrt(res)
}

func(vec *RawVec)Sum()float64{
	var res float64 = 0
	for _,v:=range vec.GetData(){
		res+=v
	}
	return res
}
func (vec *RawVec)SetSize(size int){
	vec.Size=size
}
func (vec *RawVec)GetVecType()int{
	return vec.VecType;
}
func (vec *RawVec)SetVecType(vectype int){
	vec.VecType=vectype
}
func (vec *RawVec)GetData()[]float64{
	return vec.Data
}
func (vec *RawVec)SetData(data []float64) {
	vec.Data=data
}

func (vec *RawVec)Set(index int,value float64){
	if index<0 || index>(vec.Size-1){
		log.Fatalln("index out of range",index)
	}
	vec.Data[index]=value
}
func (vec *RawVec)Get(index int)float64{
	if index<0 || index>(vec.Size-1){
		log.Fatalln("index out of range",index)
	}
	return vec.Data[index]
}

func Copy(vec Vec)Vec {
	res:=NewEmptyVec(vec.GetSize())
	for i:=0;i<res.GetSize();i++{
		res.Set(i,vec.Get(i))
	}
	return res
}


func (vec *RawVec)Print(){
	fmt.Println(vec.Data)
}
func ToFloat64Slice(arr interface{}) ([]float64, int) {
	v := reflect.ValueOf(arr)
	var l int
	var ret []float64
	if v.Kind() == reflect.Slice {
		l = v.Len()
		ret = make([]float64, l)
		for i := 0; i < l; i++ {
			ret[i] = v.Index(i).Interface().(float64)
		}
		return ret, l
	}

	res, ok := arr.(*list.List)
	if ok {
		ret = make([]float64, res.Len())
		count := 0
		for e := res.Front(); e != nil; e = e.Next() {
			ret[count] = e.Value.(float64)
			count++
		}
		return ret, count
	}

	return nil, 0
}

func NewVec(rraw interface{})*RawVec {
	raw,_:= ToFloat64Slice(rraw)
	res:=&RawVec{Data:raw, Size:len(raw)}
	return res
}

func NewVecWithDefaultValue(size int,defaultvalue float64)*RawVec{
	res:=NewEmptyVec(size)
	for i:=0;i<res.GetSize();i++{
		res.Set(i,defaultvalue)
	}
	return res
}

func NewEmptyVec(size int)*RawVec {
	data:=make([]float64,size)
	return NewVec(data)
}

