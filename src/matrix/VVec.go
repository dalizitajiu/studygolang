package matrix

type VVec struct {
	*RawVec
}
func NewVVec(rraw interface{})*VVec{
	vec:=NewVec(rraw)
	return NewVVecFromVec(vec)
}

func NewEmptyVVec(size int)*VVec{
	data:=make([]float64,size)
	return NewVVec(data)
}
func NewVVecFromVec(vec Vec)*VVec{
	res:=&VVec{}
	res.RawVec = NewEmptyVec(vec.GetSize())
	res.SetData(vec.GetData())
	res.SetSize(vec.GetSize())
	res.SetVecType(1)
	return res
}
func(vvec *VVec)T()*HVec{
	res:=&HVec{}
	res.Data = vvec.GetData()
	res.Size = vvec.GetSize()
	res.VecType = 0
	return res
}

