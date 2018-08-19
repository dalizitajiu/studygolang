package matrix

import (
	"log"
	"math"
)

type ScalaFunc func(float64)float64
type VecFunc func(float64,float64)float64
type OpeVecFunc func(Vec,Vec)*RawVec


func checkType(vec Vec,vec2 Vec){
	if vec.GetVecType()!=vec2.GetVecType(){
		log.Fatalln("向量的类型不同,无法进行操作")
	}
}
func checkSize(vec Vec,vec2 Vec){
	if vec.GetSize()!=vec2.GetSize(){
		log.Fatalln("wonng args for add",vec.GetData(),vec2.GetData())
	}
}


func Add(vec Vec,vec2 Vec)*RawVec {
	checkType(vec,vec2)
	checkSize(vec,vec2)
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,vec.Get(k)+vec2.Get(k))
	}
	res.SetVecType(vec.GetVecType())
	return res
}

func Mul(vec Vec,factor float64)*RawVec {
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,vec.Get(k)*factor)
	}
	return res
}
func MulVec(vec Vec,vec2 Vec)*RawVec {
	checkType(vec,vec2)
	checkSize(vec,vec2)
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,vec.Get(k)*vec2.Get(k))
	}
	res.SetVecType(vec.GetVecType())
	return res
}
func DivVec(vec Vec,vec2 Vec)*RawVec {
	checkType(vec,vec2)
	checkSize(vec,vec2)
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,vec.Get(k)/vec2.Get(k))
	}
	res.SetVecType(vec.GetVecType())
	return res
}
func ToHVec(vec Vec)*HVec{
	return NewHVecFromVec(vec)
}
func ToVVec(vec Vec)*VVec{
	return NewVVecFromVec(vec)
}
func Div(vec Vec,factor float64)*RawVec {
	if math.Abs(factor)<1e-7{
		log.Fatalln("unable to div zero!!")
	}
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,vec.Get(k)/factor)
	}
	return res
}
//协方差
func Cov(vec Vec,vec2 Vec)float64{
	checkSize(vec,vec2)
	checkType(vec,vec2)
	avga:=vec.Avg()
	avgb:=vec2.Avg()
	var sum float64 = 0
	for i:=0;i<vec.GetSize();i++{
		sum+=(vec.Get(i)-avga)*(vec2.Get(i)-avgb)
	}
	return sum
}
//皮尔逊相关系数
func PCCs(vec Vec,vec2 Vec)float64{
	return Cov(vec,vec2)/(vec.StdNoDiv()*vec2.StdNoDiv())
}


// 横向量乘以纵向量
func VVecDotHVet(vec *VVec,vec2 *HVec)float64{
	checkSize(vec,vec2)
	var res float64 = 0
	for k,_:=range vec.GetData(){
		res+=vec.Get(k)*vec2.Get(k)
	}
	return res
}
func Pow(vec Vec,factor float64)*RawVec {
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,math.Pow(vec.Get(k),factor))
	}
	return res
}
func Del(vec Vec,vec2 Vec)*RawVec {
	checkType(vec,vec2)
	checkSize(vec,vec2)
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,vec.Get(k)-vec2.Get(k))
	}
	return res
}

func ScalaOpe(vec Vec,factor float64,scalaFunc ScalaFunc)*RawVec {
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,scalaFunc(vec.Get(k)))
	}
	return res
}
func VecOpe(vec Vec,vec2 Vec,vecFunc VecFunc)*RawVec {
	checkType(vec,vec2)
	checkSize(vec,vec2)
	res:=NewEmptyVec(vec.GetSize())
	for k,_:=range vec.GetData(){
		res.Set(k,vecFunc(vec.Get(k),vec2.Get(k)))
	}
	return res
}
