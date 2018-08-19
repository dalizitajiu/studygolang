package main

import (
	"math"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"../matrix"
)

func sigmoid(raw float64)float64{
	return 1/(1+math.Pow(math.E,raw))
}

func genArray(datas...float64)[]float64{
	return datas
}
func toFloat64(str []string)[]float64{
	res:=make([]float64,len(str))
	for k,v:=range str{
		temp,err:= strconv.ParseFloat(strings.TrimSpace(v),64)

		if err!=nil{
			log.Fatalln("parse float64 failed",v)
		}
		res[k]=temp
	}
	return res
}

func getData(filepath string)*matrix.Matrix{
	//log.Fatalln("here")
	file,err:=os.Open(filepath)
	if err!=nil{
		log.Fatalln("open file failed",filepath)
	}
	reader:=bufio.NewReader(file)
	res:=make([]*matrix.VVec,0)
	for {
		temp,err:=reader.ReadString('\n')
		if err==nil{
			temp2:=strings.Split(string(temp),",")
			res=append(res,matrix.NewVVec(toFloat64(temp2)))
		}else {
			//处理最后一行的异常
			log.Print("err",err,temp)
			temp2:=strings.Split(string(temp),",")
			temp3:=toFloat64(temp2)
			res=append(res,matrix.NewVVec(temp3))
			break
		}

	}
	return matrix.NewMatrixFromVVec(res)
}
func getInitArgMat(num int)*matrix.Matrix{
	res:=make([]*matrix.VVec,0)
	for i:=0;i<num;i++{
		res=append(res,matrix.NewVVec(genArray(1,1)))
	}
	return matrix.NewMatrixFromVVec(res)
}
//计算fx
func fx(argvec *matrix.HVec,rawdata *matrix.Matrix)*matrix.HVec{

	res := rawdata.MulHVec(argvec)
	//log.Println("[fx]",res.Data)
	return res
}

func predict(argvec *matrix.HVec,avgvec *matrix.HVec,stdvec *matrix.HVec,target *matrix.HVec)float64{
	log.Println("[target]",target.GetData())
	log.Println("[avgvec in predict]",avgvec.Data)
	for i:=0;i<target.GetSize()-1;i++{
		min:=avgvec.Get(i)
		max:=stdvec.Get(i)
		log.Printf("min=%f\t max=%f",min,max)
		target.Set(i,(target.Get(i)-min)/(max-min))
	}
	log.Println("[target]",target.GetData())
	log.Println("[avgvec in predict]",avgvec.Data)
	return matrix.MulVec(argvec,target).Sum()
}
// 更新参数
func getUpdatedArg(argvec *matrix.HVec,learnRate float64,rawdata *matrix.Matrix,y *matrix.HVec,n int)*matrix.HVec{
	res:=matrix.NewEmptyHVec(argvec.GetSize())

	for i:=0;i<argvec.GetSize();i++{
		res.Set(i,argvec.Get(i)-2*learnRate*matrix.MulVec(matrix.Del(fx(argvec,rawdata),y),rawdata.GetHVec(0)).Sum()/float64(n))
	}
	return res
	}



func calcLoss(fx *matrix.HVec,y *matrix.HVec)*matrix.HVec{
	return matrix.ToHVec(matrix.Div(matrix.Pow(matrix.Del(fx,y),2),float64(fx.GetSize())))
}
func floss(argvect *matrix.HVec,rawdata *matrix.Matrix,y *matrix.HVec)float64{
	return matrix.Pow(matrix.Del(rawdata.MulHVec(argvect),y),2).Sum()/float64(y.GetSize())
}

func genPermutation(num int)[][]int{
	res:=make([][]int,0)
	for i:=0;i<num;i++{
		for j:=0;j<i;j++{
			res = append(res,[]int{j,i})
		}
	}
	return res
}
//自动扩展特征
func autoExtendFeature(ma *matrix.Matrix)(*matrix.Matrix,*matrix.ExtendRule){
	num:=ma.HSize
	log.Println("[原始特征数量]",num)
	rules:=matrix.NewExtendRules()
	perm:=genPermutation(num)
	for _,v:=range perm{
		rules.AddRule(matrix.Add,v[0],v[1])
		rules.AddRule(matrix.Del,v[0],v[1])
		rules.AddRule(matrix.MulVec,v[0],v[1])
		//不进行除法的特征
		//rules.AddRule(matrix.DivVec,v[0],v[1])
	}
	return rules.Apply(ma),rules
}

//多元线性回归
func main(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	mat:=getData("C:\\\\Users\\Administrator\\Desktop\\gobigdata\\src\\testdata\\data.csv")
	argvec :=matrix.NewHVec(genArray(3,1,1,1,1,1))
	rawdata := mat.SubMatrix(0,2)

	//自动构建特征
	finalma,_:=autoExtendFeature(rawdata)

	//扩展常数列
	finalma.Extends(matrix.NewHVecWithDefaultValue(finalma.VSize,1))

	minvec:=finalma.GetHVecMin()
	maxvec:=finalma.GetHVecMax()
	finalma.NormalIndex(0,1,2,3,4)
	log.Println("===============")
	finalma.Print()

	learnRate:=0.05
	y:=mat.GetHVec(-1)
	n:=y.GetSize()

	for i:=0;i<40;i++{
		argvec =getUpdatedArg(argvec,learnRate,finalma,y,n)
		log.Println(calcLoss(fx(argvec,finalma),y).Sum())
		}
	target:=matrix.NewHVec(genArray(105,455,560,-350,47775,1))
	log.Println("[预估值]",predict(argvec,minvec,maxvec,target))

	//v1:=rawdata.GetHVec(0)
	//v2:=rawdata.GetHVec(1)
	//res:=matrix.PCCs(v1,v2)//计算
	//println(res)
	}
