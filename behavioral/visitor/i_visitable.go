package visitor

type IVistable interface {
	Accept(IVisitor)
}
