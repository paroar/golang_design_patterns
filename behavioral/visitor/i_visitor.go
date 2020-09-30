package visitor

type IVisitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}
