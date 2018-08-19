package matrix
type HVec struct {
	*RawVec
}
func NewHVec(rraw interface{})*HVec{
	vec:=NewVec(rraw)
	vec.SetVecType(0)
	return ToHVec(vec)
}
func NewHVecFromVec(vec Vec)*HVec{
	res:=&HVec{}
	res.RawVec = NewEmptyVec(vec.GetSize())
	res.SetData(vec.GetData())
	res.SetSize(vec.GetSize())
	res.SetVecType(0)
	return res
}
func NewEmptyHVec(size int)*HVec{
	data:=make([]float64,size)
	return NewHVec(data)
}
func NewHVecWithDefaultValue(size int,defaultvalue float64)*HVec{
	res:=NewVecWithDefaultValue(size,defaultvalue)
	return ToHVec(res)
}
func(vvec *HVec)T()*VVec{
	res:=&VVec{}
	res.Data = vvec.GetData()
	res.Size = vvec.GetSize()
	res.VecType = 0
	return res
}