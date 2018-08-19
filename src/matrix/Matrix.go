package matrix

import (
	"fmt"
	"log"
)

type Matrix struct {
	Data []*HVec
	VSize int
	HSize int
}

func (matrix *Matrix)Set(x,y int,value float64){
	matrix.Data[y].Set(x,value)
}
func (matrix *Matrix)Get(x,y int)float64{
	return matrix.Data[y].Get(x)
}
//vvec 行向量
//hvec 列向量
func NewMatrixFromVVec(vecs []*VVec)*Matrix{
	res:=&Matrix{}
	res.VSize = len(vecs)
	res.HSize = vecs[0].GetSize()
	res.Data=make([]*HVec,res.HSize)

	for i:=0;i<res.HSize;i++{
		res.Data[i] = NewEmptyHVec(res.VSize)
	}
	for i,v:=range vecs{

		for j,v2:=range v.Data{
			res.Set(i,j,v2)
		}
	}
	res.HSize = vecs[0].GetSize()
	res.VSize = len(vecs)
	return res
}
func NewMatrixFromHvec(vecs []*HVec)*Matrix{
	res:=&Matrix{}
	res.VSize = vecs[0].GetSize()
	res.HSize = len(vecs)
	res.Data = vecs
	return res
}
func (matrix *Matrix)GetHVec(index int)*HVec{
	if index<0{
		return matrix.Data[matrix.HSize+index]
	}
	return matrix.Data[index]
}
func (matrix *Matrix)GetVVec(index int)*VVec{
	actualindex:=index
	if index<0{
		actualindex=matrix.VSize+index
	}
	res:=NewEmptyVVec(matrix.HSize)
	for k,v:=range matrix.Data{
		res.Set(k,v.Get(actualindex))
	}
	return res
}
func (matrix *Matrix)Print(){
	for i:=0;i<matrix.VSize;i++{
		for j:=0;j<matrix.HSize;j++{
			fmt.Print(matrix.Get(i,j),"\t")
		}
		fmt.Print("\n")
	}
	fmt.Println("[HSize]",matrix.HSize,"[VSize]",matrix.VSize)
}
func (matrix *Matrix)SubMatrix(start,end int)*Matrix{
	temp:=make([]*HVec,0)
	for i:=start;i<end;i++{

		temp=append(temp,ToHVec(Copy(matrix.GetHVec(i))))
	}
	return NewMatrixFromHvec(temp)
}
func (matrix *Matrix)MulHVec(vec *HVec)*HVec{
	if matrix.HSize!=vec.GetSize(){
		vec.Print()
		log.Fatalln("参数尺寸不匹配",matrix.HSize,vec.GetSize())
	}
	res:=NewEmptyHVec(matrix.VSize)
	for i:=0;i<matrix.VSize;i++{
		temp:=matrix.GetVVec(i)
		res.Set(i,VVecDotHVet(temp,vec))
	}
	return res
}
func (matrix *Matrix)Normal(){
	for k,_:=range matrix.Data{
		matrix.Data[k].Normal()
	}
}
func (matrix *Matrix)NormalIndex(indexs...int){
	for k,_:=range indexs{
		matrix.Data[k].Normal()
	}
}
func (matrix *Matrix)GetHVecAvg()*HVec{
	res:=make([]float64,0)
	for _,v:=range matrix.Data{
		res=append(res,v.Avg())
	}
	return NewHVec(res)
}
func (matrix *Matrix)GetHVecMax()*HVec{
	res:=make([]float64,0)
	for _,v:=range matrix.Data{
		res=append(res,v.Max())
	}
	return NewHVec(res)
}
func (matrix *Matrix)GetHVecMin()*HVec{
	res:=make([]float64,0)
	for _,v:=range matrix.Data{
		res=append(res,v.Min())
	}
	return NewHVec(res)
}
func (matrix *Matrix)GetHVecStd()*HVec{
	res:=make([]float64,0)
	for _,v:=range matrix.Data{
		res=append(res,v.Std())
	}
	return NewHVec(res)
}
func (matrix *Matrix)Extends(vecs...*HVec)*Matrix{
	for _,v:=range vecs{
		if !checkMatrixVecSize(matrix,v){
			log.Fatalln("参数长度不匹配",v)
		}
		matrix.Data = append(matrix.Data,v)
		matrix.HSize++
	}
	return matrix
}

func (matrix *Matrix)Remove(index int)bool{
	realindex := index
	if index>=matrix.HSize{
		log.Fatalln("[Remove] 超出了元素范围")
		return false
	}
	if index<0{
		realindex = matrix.HSize+index
	}
	matrix.Data=append(matrix.Data[:realindex],matrix.Data[realindex+1:]...)
	matrix.HSize--
	return true
}

func checkMatrixVecSize(matrix *Matrix,vec Vec)bool{
	if matrix.VSize!=vec.GetSize(){
		return false
	}
	return true
}

