package matrix

type Rule1 struct {
	rule OpeVecFunc
	scope1 int
	scope2 int
}

type ExtendRule struct {
	Rules []*Rule1
}

func NewExtendRules()*ExtendRule{
	res:=&ExtendRule{}
	res.Rules = make([]*Rule1,0)
	return res
}

func (extendrules *ExtendRule)AddRule(fun OpeVecFunc,i,j int)*ExtendRule{
	extendrules.Rules = append(extendrules.Rules,&Rule1{fun,i,j})
	return extendrules
}
func (extendRules *ExtendRule)Apply(matrix *Matrix)*Matrix{
	copy:=matrix.SubMatrix(0,matrix.HSize)
	for _,v:=range extendRules.Rules{
		copy.Extends(ToHVec(v.rule(matrix.GetHVec(v.scope1),matrix.GetHVec(v.scope2))))
	}
	return copy
}
